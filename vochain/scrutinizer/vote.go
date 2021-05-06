package scrutinizer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/timshannon/badgerhold/v3"
	"go.vocdoni.io/proto/build/go/models"
	"google.golang.org/protobuf/proto"

	"go.vocdoni.io/dvote/crypto/nacl"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
)

// ErrNoResultsYet is an error returned to indicate the process exist but
// it does not have yet reuslts
var ErrNoResultsYet = fmt.Errorf("no results yet")

// ErrNotFoundIndatabase is raised if a database query returns no results
var ErrNotFoundInDatabase = badgerhold.ErrNotFound

// The string to search on the KV database error to identify a transaction conflict.
// If the KV (currently badger) returns this error, it is considered non fatal and the
// transaction will be retried until it works.
// This check is made comparing string in order to avoid importing a specific KV
// implementation.
const kvErrorStringForRetry = "Transaction Conflict"

// GetVoteReference gets the reference for an AddVote transaction.
// This reference can then be used to fetch the vote transaction directly from the BlockStore.
func (s *Scrutinizer) GetEnvelopeReference(nullifier []byte) (*VoteReference, error) {
	txRef := &VoteReference{}
	return txRef, s.db.FindOne(txRef, badgerhold.Where(badgerhold.Key).Eq(nullifier))
}

// GetEnvelope retreives an Envelope from the Blockchain block store identified by its nullifier.
// Returns the envelope and the signature (if any).
func (s *Scrutinizer) GetEnvelope(nullifier []byte) (*models.EnvelopePackage, []byte, error) {
	txRef, err := s.GetEnvelopeReference(nullifier)
	if err != nil {
		return nil, nil, err
	}
	stx, _, err := s.App.GetTx(txRef.Height, txRef.TxIndex)
	if err != nil {
		return nil, nil, err
	}
	tx := &models.Tx{}
	if err := proto.Unmarshal(stx.Tx, tx); err != nil {
		return nil, nil, err
	}
	envelope := tx.GetVote()
	if envelope == nil {
		return nil, nil, fmt.Errorf("transaction is not an Envelope")
	}
	return &models.EnvelopePackage{
		Envelope: envelope,
		Weight:   txRef.Weight.Bytes(),
		TxIndex:  txRef.TxIndex,
		Height:   txRef.Height,
	}, stx.Signature, nil
}

// WalkEnvelopes executes callback for each envelopes of the ProcessId.
// The callback function is executed async (in a goroutine) if async=true.
// The method will return once all goroutines have finished the work.
func (s *Scrutinizer) WalkEnvelopes(processId []byte, async bool,
	callback func(*models.VoteEnvelope, *big.Int)) error {
	wg := sync.WaitGroup{}
	err := s.db.ForEach(
		badgerhold.Where("ProcessID").Eq(processId),
		func(txRef *VoteReference) error {
			wg.Add(1)
			processVote := func() {
				defer wg.Done()
				stx, _, err := s.App.GetTx(txRef.Height, txRef.TxIndex)
				if err != nil {
					log.Errorf("could not get tx: %v", err)
					return
				}
				tx := &models.Tx{}
				if err := proto.Unmarshal(stx.Tx, tx); err != nil {
					log.Warnf("could not unmarshal tx: %v", err)
					return
				}
				envelope := tx.GetVote()
				if envelope == nil {
					log.Errorf("transaction is not an Envelope")
					return
				}
				callback(envelope, txRef.Weight)
			}
			if async {
				go processVote()
			} else {
				processVote()
			}
			return nil
		})
	wg.Wait()
	return err
}

// GetEnvelopes retreives all Envelopes of a ProcessId from the Blockchain block store
func (s *Scrutinizer) GetEnvelopes(processId []byte) (*models.EnvelopePackageList, error) {
	envelopes := new(models.EnvelopePackageList)
	err := s.db.ForEach(
		badgerhold.Where("ProcessID").Eq(processId).Index("ProcessID"),
		func(txRef *VoteReference) error {
			stx, _, err := s.App.GetTx(txRef.Height, txRef.TxIndex)
			if err != nil {
				return err
			}
			tx := &models.Tx{}
			if err := proto.Unmarshal(stx.Tx, tx); err != nil {
				return err
			}
			envelope := tx.GetVote()
			if envelope == nil {
				return fmt.Errorf("transaction is not an Envelope")
			}
			envelopes.Envelopes = append(envelopes.GetEnvelopes(), &models.EnvelopePackage{
				Envelope: envelope,
				Weight:   txRef.Weight.Bytes(),
				TxIndex:  txRef.TxIndex,
				Height:   txRef.Height,
			})
			return nil
		})
	return envelopes, err
}

// GetEnvelopeHeight returns the number of envelopes for a processId.
// If processId is empty, returns the total number of envelopes.
func (s *Scrutinizer) GetEnvelopeHeight(processId []byte) (uint64, error) {
	if len(processId) > 0 {
		cc, ok := s.envelopeHeightCache.Get(string(processId))
		if ok {
			return cc.(uint64), nil
		}
		// TODO: Warning, int can overflow
		c, err := s.db.Count(&VoteReference{},
			badgerhold.Where("ProcessID").Eq(processId).Index("ProcessID"))
		if err != nil {
			return 0, err
		}
		c64 := uint64(c)
		s.envelopeHeightCache.Add(string(processId), c64)
		return c64, nil
	}

	// If no processId is provided, count all envelopes
	return atomic.LoadUint64(s.countTotalEnvelopes), nil
}

// ComputeResult process a finished voting, compute the results and saves it in the Storage.
// Once this function is called, any future live vote event for the processId will be discarted.
func (s *Scrutinizer) ComputeResult(processID []byte) error {
	log.Debugf("computing results for %x", processID)

	// Get process from database
	p, err := s.ProcessInfo(processID)
	if err != nil {
		return fmt.Errorf("computeResult: cannot load processID %x from database: %w", processID, err)
	}

	// Compute the results
	var results *Results
	if results, err = s.computeFinalResults(p); err != nil {
		return err
	}
	p.HaveResults = true
	p.FinalResults = true
	if err := s.db.Update(processID, p); err != nil {
		return fmt.Errorf("computeResults: cannot update processID %x: %w, ", processID, err)
	}
	s.addVoteLock.Lock()
	defer s.addVoteLock.Unlock()
	if err := s.db.Upsert(processID, results); err != nil {
		return err
	}

	// Execute callbacks
	for _, l := range s.eventListeners {
		go l.OnComputeResults(results)
	}
	return nil
}

// GetResults returns the current result for a processId aggregated in a two dimension int slice
func (s *Scrutinizer) GetResults(processID []byte) (*Results, error) {
	if n, err := s.db.Count(&Process{},
		badgerhold.Where("ID").Eq(processID).
			And("HaveResults").Eq(true).
			And("Status").Ne(int32(models.ProcessStatus_CANCELED))); err != nil {
		return nil, err
	} else if n == 0 {
		return nil, ErrNoResultsYet
	}

	s.addVoteLock.RLock()
	defer s.addVoteLock.RUnlock()
	results := &Results{}
	if err := s.db.FindOne(results, badgerhold.Where("ProcessID").Eq(processID)); err != nil {
		if err == badgerhold.ErrNotFound {
			return nil, ErrNoResultsYet
		}
		return nil, err
	}
	return results, nil
}

// GetResultsWeight returns the current weight of cast votes for a processId.
func (s *Scrutinizer) GetResultsWeight(processID []byte) (*big.Int, error) {
	s.addVoteLock.RLock()
	defer s.addVoteLock.RUnlock()
	results := &Results{}
	if err := s.db.FindOne(results, badgerhold.Where("ProcessID").Eq(processID)); err != nil {
		return nil, err
	}
	return results.Weight, nil
}

// unmarshalVote decodes the base64 payload to a VotePackage struct type.
// If the votePackage is encrypted the list of keys to decrypt it should be provided.
// The order of the Keys must be as it was encrypted.
// The function will reverse the order and use the decryption keys starting from the
// last one provided.
func unmarshalVote(votePackage []byte, keys []string) (*types.VotePackage, error) {
	var vote types.VotePackage
	rawVote := make([]byte, len(votePackage))
	copy(rawVote, votePackage)
	// if encryption keys, decrypt the vote
	if len(keys) > 0 {
		for i := len(keys) - 1; i >= 0; i-- {
			priv, err := nacl.DecodePrivate(keys[i])
			if err != nil {
				return nil, fmt.Errorf("cannot create private key cipher: (%s)", err)
			}
			if rawVote, err = priv.Decrypt(rawVote); err != nil {
				return nil, fmt.Errorf("cannot decrypt vote with index key %d: %w", i, err)
			}
		}
	}
	if err := json.Unmarshal(rawVote, &vote); err != nil {
		return nil, fmt.Errorf("cannot unmarshal vote: %w", err)
	}
	return &vote, nil
}

// addLiveVote adds the envelope vote to the results. It does not commit to the database.
// This method is triggered by OnVote callback for each vote added to the blockchain.
// If encrypted vote, only weight will be updated.
func (s *Scrutinizer) addLiveVote(pid []byte, votePackage []byte, weight *big.Int,
	results *Results) error {
	// If live process, add vote to temporary results
	var vote *types.VotePackage
	if open, err := s.isOpenProcess(pid); open && err == nil {
		vote, err = unmarshalVote(votePackage, []string{})
		if err != nil {
			log.Warnf("cannot unmarshal vote: %v", err)
			vote = nil
		}
	} else if err != nil {
		return fmt.Errorf("cannot check if process is open: %v", err)
	}

	// Add the vote only if the election is unencrypted
	if vote != nil {
		if err := results.AddVote(vote.Votes, weight, nil); err != nil {
			return err
		}
	} else {
		// If encrypted, just add the weight
		results.Weight.Add(results.Weight, weight)
	}
	return nil
}

// addVoteIndex adds the nullifier reference to the kv for fetching vote Txs from BlockStore.
// This method is triggered by Commit callback for each vote added to the blockchain.
// If txn is provided the vote will be added on the transaction (without performing a commit).
func (s *Scrutinizer) addVoteIndex(nullifier, pid []byte, blockHeight uint32,
	weight []byte, txIndex int32, txn *badger.Txn) error {
	if txn != nil {
		return s.db.TxInsert(txn, nullifier, &VoteReference{
			Nullifier:    nullifier,
			ProcessID:    pid,
			Height:       blockHeight,
			Weight:       new(big.Int).SetBytes(weight),
			TxIndex:      txIndex,
			CreationTime: time.Now(),
		})
	}
	return s.db.Insert(nullifier, &VoteReference{
		Nullifier:    nullifier,
		ProcessID:    pid,
		Height:       blockHeight,
		Weight:       new(big.Int).SetBytes(weight),
		TxIndex:      txIndex,
		CreationTime: time.Now(),
	})
}

func (s *Scrutinizer) addProcessToLiveResults(pid []byte) {
	s.liveResultsProcs.Store(string(pid), true)
}

func (s *Scrutinizer) delProcessFromLiveResults(pid []byte) {
	s.liveResultsProcs.Delete(string(pid))
}

func (s *Scrutinizer) isProcessLiveResults(pid []byte) bool {
	_, ok := s.liveResultsProcs.Load(string(pid))
	return ok
}

// commitVotes adds the votes and weight from results to the local database.
// Important: it does not overwrite the already stored results but update them
// by adding the new content to the existing one.
func (s *Scrutinizer) commitVotes(pid []byte, partialResults *Results, height uint32) error {
	// If the recovery bootstrap is running, wait.
	s.recoveryBootLock.RLock()
	defer s.recoveryBootLock.RUnlock()
	// The next lock avoid Transaction Conflicts
	s.addVoteLock.Lock()
	defer s.addVoteLock.Unlock()
	update := func(record interface{}) error {
		stored, ok := record.(*Results)
		if !ok {
			return fmt.Errorf("record isn't the correct type! Wanted Result, got %T", record)
		}
		return stored.Add(partialResults)
	}

	var err error
	// If badgerhold returns a transaction conflict error message, we try again
	// until it works (while maxTries < 1000)
	maxTries := 1000
	for maxTries > 0 {
		err = s.db.UpdateMatching(&Results{},
			badgerhold.Where("ProcessID").Eq(pid).And("Final").Eq(false), update)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), kvErrorStringForRetry) {
			maxTries--
			time.Sleep(5 * time.Millisecond)
			continue
		}
		return err
	}
	if maxTries == 0 {
		err = fmt.Errorf("got too much transaction conflicts")
	}
	if err != nil {
		log.Debugf("saved %d votes with total weight of %s on process %x", len(partialResults.Votes),
			partialResults.Weight, pid)
	}
	return err
}

func (s *Scrutinizer) computeFinalResults(p *Process) (*Results, error) {
	if p == nil {
		return nil, fmt.Errorf("process is nil")
	}
	if p.VoteOpts.MaxCount == 0 || p.VoteOpts.MaxValue == 0 {
		return nil, fmt.Errorf("computeNonLiveResults: maxCount or maxValue are zero")
	}
	if p.VoteOpts.MaxCount > MaxQuestions || p.VoteOpts.MaxValue > MaxOptions {
		return nil, fmt.Errorf("maxCount or maxValue overflows hardcoded maximums")
	}
	results := &Results{
		Votes:        newEmptyVotes(int(p.VoteOpts.MaxCount), int(p.VoteOpts.MaxValue)+1),
		ProcessID:    p.ID,
		Weight:       new(big.Int).SetUint64(0),
		Final:        true,
		VoteOpts:     p.VoteOpts,
		EnvelopeType: p.Envelope,
		Height:       s.App.Height(),
	}

	var nvotes uint64
	var err error
	lock := sync.Mutex{}

	if err = s.WalkEnvelopes(p.ID, true, func(vote *models.VoteEnvelope,
		weight *big.Int) {
		var vp *types.VotePackage
		var err error
		if p.Envelope.GetEncryptedVotes() {
			if len(p.PrivateKeys) < len(vote.GetEncryptionKeyIndexes()) {
				log.Errorf("encryptionKeyIndexes has too many fields")
				return
			} else {
				keys := []string{}
				for _, k := range vote.GetEncryptionKeyIndexes() {
					if k >= types.KeyKeeperMaxKeyIndex {
						log.Warn("key index overflow")
						return
					}
					keys = append(keys, p.PrivateKeys[k])
				}
				if len(keys) == 0 || err != nil {
					log.Warn("no keys provided or wrong index")
					return
				} else {
					vp, err = unmarshalVote(vote.GetVotePackage(), keys)
				}
			}
		} else {
			vp, err = unmarshalVote(vote.GetVotePackage(), []string{})
		}
		if err != nil {
			log.Debugf("vote invalid: %v", err)
			return
		}

		if err = results.AddVote(vp.Votes, weight, &lock); err != nil {
			log.Warnf("addVote failed: %v", err)
			return
		}
		atomic.AddUint64(&nvotes, 1)
	}); err == nil {
		log.Infof("computed results for process %x with %d votes", p.ID, nvotes)
		log.Debugf("results: %s", results)
	}
	return results, err
}

// BuildProcessResult takes the indexer Results type and builds the protobuf type ProcessResult.
// EntityId should be provided as addition field to include in ProcessResult.
func BuildProcessResult(results *Results, entityID []byte) *models.ProcessResult {
	// build the protobuf type for Results
	qr := []*models.QuestionResult{}
	for i := range results.Votes {
		qr = append(qr, &models.QuestionResult{})
		for j := range results.Votes[i] {
			qr[i].Question = append(qr[i].Question, results.Votes[i][j].Bytes())
		}
	}
	return &models.ProcessResult{
		ProcessId: results.ProcessID,
		EntityId:  entityID,
		Votes:     qr,
	}
}
