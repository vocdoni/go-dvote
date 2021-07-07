package vochain

import (
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	ethtoken "github.com/vocdoni/storage-proofs-eth-go/token"
	"go.vocdoni.io/dvote/crypto/ethereum"
	models "go.vocdoni.io/proto/build/go/models"
)

// AddToRollingCensus adds a new key to an existing rolling census.
// If census does not exist yet it will be created.
func (s *State) AddToRollingCensus(pid []byte, key []byte, weight *big.Int) error {
	/*
		// In the state we only store the last census root (as value) using key as index
		s.Lock()
		err := s.Store.Tree(CensusTree).Add(p.ProcessId, root)
		s.Unlock()
		if err != nil {
			return err
		}
	*/
	return nil
}

// PurgeRollingCensus removes a rolling census from the permanent store
// If the census does not exist, it does nothing.
func (s *State) PurgeRollingCensus(pid []byte) error {
	return nil
}

// GetRollingCensusRoot returns the last rolling census root for a process id
func (s *State) GetRollingCensusRoot(pid []byte, isQuery bool) ([]byte, error) {
	/*	s.Lock()
		err := s.Store.Tree(CensusTree).Add(p.ProcessId, root)
		s.Unlock()
		if err != nil {
			return err
		}GetCensusRoot
	*/
	return nil, nil
}

// RegisterKeyTxCheck validates a registerKeyTx transaction against the state
func RegisterKeyTxCheck(vtx *models.Tx, txBytes, signature []byte, state *State) error {
	tx := vtx.GetRegisterKey()

	// Sanity checks
	if tx == nil {
		return fmt.Errorf("register key transaction is nil")
	}
	process, err := state.Process(tx.ProcessId, false)
	if err != nil {
		return fmt.Errorf("cannot fetch processId: %w", err)
	}
	if process == nil || process.EnvelopeType == nil || process.Mode == nil {
		return fmt.Errorf("process %x malformed", tx.ProcessId)
	}
	if state.Height() >= process.StartBlock {
		return fmt.Errorf("process %x already started", tx.ProcessId)
	}
	if process.Status != models.ProcessStatus_READY {
		return fmt.Errorf("process %x not in READY state", tx.ProcessId)
	}
	if tx.Proof == nil {
		return fmt.Errorf("proof missing on registerKeyTx")
	}
	if signature == nil {
		return fmt.Errorf("signature missing on voteTx")
	}
	if len(tx.NewKey) < 32 { // TODO: check the correctnes of the new public key
		return fmt.Errorf("newKey wrong size")
	}
	if process.GetEnvelopeType().Anonymous {
		// TODO: if we finally use Poseidon(babyjub.PubKey) then we need to perform a different check
		bj := babyjub.NewPoint()
		var point [32]byte
		copy(point[:], tx.NewKey)
		if _, err := bj.Decompress(point); err != nil {
			return fmt.Errorf("invalid newKey: %w", err)
		}
	}

	var pubKey []byte
	pubKey, err = ethereum.PubKeyFromSignature(txBytes, signature)
	if err != nil {
		return fmt.Errorf("cannot extract public key from signature: (%w)", err)
	}
	addr, err := ethereum.AddrFromPublicKey(pubKey)
	if err != nil {
		return fmt.Errorf("cannot extract address from public key: (%w)", err)
	}

	// check census origin and compute vote digest identifier
	var pubKeyDigested []byte
	switch process.CensusOrigin {
	case models.CensusOrigin_OFF_CHAIN_TREE,
		models.CensusOrigin_OFF_CHAIN_TREE_WEIGHTED:
		pubKeyDigested = pubKey
	case models.CensusOrigin_OFF_CHAIN_CA:
		pubKeyDigested = addr.Bytes()
	case models.CensusOrigin_ERC20:
		if process.EthIndexSlot == nil {
			return fmt.Errorf("index slot not found for process %x", process.ProcessId)
		}
		slot, err := ethtoken.GetSlot(addr.Hex(), int(*process.EthIndexSlot))
		if err != nil {
			return fmt.Errorf("cannot fetch slot: %w", err)
		}
		pubKeyDigested = slot[:]
	default:
		return fmt.Errorf("census origin not compatible")
	}

	// check the digested payload has a minimum length
	if len(pubKeyDigested) < 20 { // Minimum size is an Ethereum Address
		return fmt.Errorf("cannot digest public key: (%w)", err)
	}

	// check census proof
	var valid bool
	var weight *big.Int
	valid, weight, err = CheckProof(tx.Proof,
		process.CensusOrigin,
		process.CensusRoot,
		process.ProcessId,
		pubKeyDigested)
	if err != nil {
		return fmt.Errorf("proof not valid: (%w)", err)
	}
	if !valid {
		return fmt.Errorf("proof not valid")
	}
	tx.Weight = weight.Bytes() // TODO: support weight

	return nil
}
