package scrutinizer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"go.vocdoni.io/proto/build/go/models"

	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
)

// ProcessInfo returns the available information regarding an election process id
func (s *Scrutinizer) ProcessInfo(pid []byte) (*Process, error) {
	proc := &Process{}
	return proc, s.db.One("ID", pid, proc)
}

// ProcessList returns a list of process identifiers (PIDs) registered in the Vochain.
// EntityID, namespace and status are optional filters, if declared as zero-values
// will be ignored. Status is one of READY, CANCELED, ENDED, PAUSED, RESULTS
func (s *Scrutinizer) ProcessList(entityID []byte, namespace uint32,
	status string, withResults bool, from, max int) ([][]byte, error) {
	// For filtering on Status we use a badgerhold match function.
	// If status is not defined, then the match function will return always true.
	/*	statusnum := int32(-1)
		statusfound := false
		if status != "" {
			if statusnum, statusfound = models.ProcessStatus_value[status]; !statusfound {
				return nil, fmt.Errorf("processList: status %s is unknown", status)
			}
		}
		statusMatchFunc := func(r *badgerhold.RecordAccess) (bool, error) {
			if statusnum == -1 {
				return true, nil
			}
			if r.Field().(int32) == statusnum {
				return true, nil
			}
			return false, nil
		}

		// For filtering on withResults we use also a match function
		wResultsMatchFunc := func(r *badgerhold.RecordAccess) (bool, error) {
			if !withResults {
				return true, nil
			}
			return r.Field().(bool), nil
		}

		// For EntityID and Namespace we use different queries, since they are indexes and the
		// performance improvement is quite relevant.
		var err error
		var procs [][]byte
		/*	switch {
			case namespace == 0 && len(entityID) > 0:
				err = s.db.ForEach(
					badgerhold.Where("EntityID").Eq(entityID).
						And("Status").MatchFunc(statusMatchFunc).
						And("HaveResults").MatchFunc(wResultsMatchFunc).
						Index("EntityID").
						SortBy("CreationTime").
						Skip(from).
						Limit(max),
					func(p *Process) error {
						procs = append(procs, p.ID)
						return nil
					})
			case namespace > 0 && len(entityID) == 0:
				err = s.db.ForEach(
					badgerhold.Where("Namespace").Eq(namespace).
						And("Status").MatchFunc(statusMatchFunc).
						And("HaveResults").MatchFunc(wResultsMatchFunc).
						Index("Namespace").
						SortBy("CreationTime").
						Skip(from).
						Limit(max),
					func(p *Process) error {
						procs = append(procs, p.ID)
						return nil
					})
			case namespace == 0 && len(entityID) == 0:
				err = s.db.ForEach(
					badgerhold.Where("ID").Ne([]byte{}).
						And("Status").MatchFunc(statusMatchFunc).
						And("HaveResults").MatchFunc(wResultsMatchFunc).
						SortBy("CreationTime").
						Skip(from).
						Limit(max),
					func(p *Process) error {
						procs = append(procs, p.ID)
						return nil
					})
			default:
				err = s.db.ForEach(
					badgerhold.Where("EntityID").Eq(entityID).
						And("Namespace").Eq(namespace).
						And("Status").MatchFunc(statusMatchFunc).
						And("HaveResults").MatchFunc(wResultsMatchFunc).
						Index("EntityID").
						SortBy("CreationTime").
						Skip(from).
						Limit(max),
					func(p *Process) error {
						procs = append(procs, p.ID)
						return nil
					})
			}
	*/
	//return procs, err
	return nil, nil
}

// ProcessCount return the number of processes indexed
func (s *Scrutinizer) ProcessCount() int64 {
	c, err := s.db.Count(&Process{})
	if err != nil {
		log.Warnf("cannot count processes: %v", err)
	}
	return int64(c)
}

// EntityList returns the list of entities indexed by the scrutinizer
func (s *Scrutinizer) EntityList(max, from int) []string {
	entities := []string{}
	if err := s.db.Select().OrderBy("CreationTime").Skip(from).Limit(max).Each(&Entity{},
		func(record interface{}) error {
			e := record.(*Entity)
			entities = append(entities, fmt.Sprintf("%x", e.ID))
			return nil
		}); err != nil {
		log.Warnf("error listing entities: %v", err)
	}
	return entities
}

// EntityCount return the number of entities indexed by the scrutinizer
func (s *Scrutinizer) EntityCount() int64 {
	c, err := s.db.Count(&Entity{})
	if err != nil {
		log.Warnf("cannot count entities: %v", err)
	}
	return int64(c)
}

// Return whether a process must have live results or not
func (s *Scrutinizer) isOpenProcess(processID []byte) (bool, error) {
	p, err := s.App.State.Process(processID, false)
	if err != nil {
		return false, err
	}
	if p == nil || p.EnvelopeType == nil {
		return false, fmt.Errorf("cannot fetch process %x or envelope type not defined", processID)
	}
	return !p.EnvelopeType.EncryptedVotes, nil
}

// compute results if the current heigh has scheduled ending processes
func (s *Scrutinizer) computePendingProcesses(height uint32) {
	procs := []Process{}
	if err := s.db.Find("Rheight", height, &procs); err != nil {
		if err.Error() == storm.ErrNotFound.Error() {
			return
		}
		log.Error(err)
		return
	}
	for _, p := range procs {
		initT := time.Now()
		if err := s.ComputeResult(p.ID); err != nil {
			log.Warnf("cannot compute results for %x: (%v)", p.ID, err)
			continue
		}
		log.Infof("results compute on %x took %s", p.ID, time.Since(initT).String())
		p.FinalResults = true
		p.HaveResults = true
		if err := s.db.Update(&p); err != nil {
			log.Errorf("cannot update results for process %x: %v", p.ID, err)
		}
	}
}

// newEmptyProcess creates a new empty process and stores it into the database.
// The process must exist on the Vochain state, else an error is returned.
func (s *Scrutinizer) newEmptyProcess(pid []byte) error {
	p, err := s.App.State.Process(pid, false)
	if err != nil {
		return fmt.Errorf("cannot create new empty process: %w", err)
	}
	options := p.GetVoteOptions()
	if options == nil {
		return fmt.Errorf("newEmptyProcess: vote options is nil")
	}
	if options.MaxCount == 0 || options.MaxValue == 0 {
		return fmt.Errorf("newEmptyProcess: maxCount or maxValue are zero")
	}

	// Check for overflows
	if options.MaxCount > MaxQuestions || options.MaxValue > MaxOptions {
		return fmt.Errorf("maxCount or maxValue overflows hardcoded maximums")
	}

	// Create results in the indexer database
	s.addVoteLock.Lock()
	if err := s.db.Save(&Results{
		ProcessID: pid,
		// MaxValue requires +1 since 0 is also an option
		Votes:      newEmptyVotes(int(options.MaxCount), int(options.MaxValue)+1),
		Weight:     new(big.Int).SetUint64(0),
		Signatures: []types.HexBytes{},
	}); err != nil {
		s.addVoteLock.Unlock()
		return fmt.Errorf("newEmptyProcess: cannot save Results: %w", err)
	}
	s.addVoteLock.Unlock()

	// Get the block time from the Header
	currentBlockTime := time.Unix(s.App.State.Header(false).Timestamp, 0)

	// Add the entity to the indexer database
	eid := p.GetEntityId()
	if err := s.db.Save(&Entity{ID: eid, CreationTime: currentBlockTime}); err != nil {
		if err == storm.ErrAlreadyExists {
			if err := s.db.Update(&Entity{ID: eid, CreationTime: currentBlockTime}); err != nil {
				return fmt.Errorf("newEmptyProcess: cannot update entity: %w", err)
			}
		} else {
			return fmt.Errorf("newEmptyProcess: cannot update entity: %w", err)
		}
	}

	compResultsHeight := uint32(0)
	if live, err := s.isOpenProcess(pid); err != nil {
		return fmt.Errorf("newEmptyProcess: cannot check if process live: %w", err)
	} else {
		if live {
			compResultsHeight = p.GetBlockCount() + p.GetStartBlock()
		}
	}

	// Create and store process in the indexer database
	proc := &Process{
		ID:           pid,
		EntityID:     eid,
		StartBlock:   p.GetStartBlock(),
		EndBlock:     p.GetBlockCount() + p.GetStartBlock(),
		Rheight:      compResultsHeight,
		HaveResults:  compResultsHeight > 0,
		CensusRoot:   p.GetCensusRoot(),
		CensusURI:    p.GetCensusURI(),
		CensusOrigin: int32(p.GetCensusOrigin()),
		Status:       int32(p.GetStatus()),
		Namespace:    p.GetNamespace(),
		PrivateKeys:  p.EncryptionPrivateKeys,
		PublicKeys:   p.EncryptionPublicKeys,
		Envelope:     p.GetEnvelopeType(),
		Mode:         p.GetMode(),
		VoteOpts:     p.GetVoteOptions(),
		CreationTime: currentBlockTime,
	}
	log.Debugf("new indexer process %s", proc.String())
	return s.db.Save(proc)
}

// updateProcess synchronize those fields that can be updated on a existing process
// with the information obtained from the Vochain state
func (s *Scrutinizer) updateProcess(pid []byte) error {
	p, err := s.App.State.Process(pid, false)
	if err != nil {
		return fmt.Errorf("updateProcess: cannot fetch process %x: %w", pid, err)
	}
	update := Process{}
	if err := s.db.One("ID", pid, &update); err != nil {
		return err
	}
	update.EndBlock = p.GetBlockCount() + p.GetStartBlock()
	update.CensusRoot = p.GetCensusRoot()
	update.CensusURI = p.GetCensusURI()
	update.PrivateKeys = p.EncryptionPrivateKeys
	update.PublicKeys = p.EncryptionPublicKeys
	// If the process is transacting to CANCELED, ensure results are not computed and remove
	// them from the KV database.
	if update.Status != int32(models.ProcessStatus_CANCELED) &&
		p.GetStatus() == models.ProcessStatus_CANCELED {
		update.HaveResults = false
		update.FinalResults = true
		update.Rheight = 0
		if err := s.db.DeleteStruct(&Results{ProcessID: pid}); err != nil {
			log.Warnf("cannot remove CANCELED results: %v", err)
		}
	}
	update.Status = int32(p.GetStatus())
	return s.db.Update(&update)
}

func (s *Scrutinizer) setResultsHeight(pid []byte, height uint32) error {
	return s.db.Select(q.Eq("ID", pid)).Each(&Process{},
		func(record interface{}) error {
			update, ok := record.(*Process)
			if !ok {
				return fmt.Errorf("record isn't the correct type! Wanted Result, got %T", record)
			}
			update.Rheight = height
			return nil
		})
}

func newEmptyVotes(questions, options int) [][]*big.Int {
	if questions == 0 || options == 0 {
		return nil
	}
	results := [][]*big.Int{}
	for i := 0; i < questions; i++ {
		question := []*big.Int{}
		for j := 0; j < options; j++ {
			question = append(question, big.NewInt(0))
		}
		results = append(results, question)
	}
	return results
}
