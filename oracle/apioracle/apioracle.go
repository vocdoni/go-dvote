package apioracle

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/storage-proofs-eth-go/token"
	"go.vocdoni.io/dvote/api"
	"go.vocdoni.io/dvote/chain"
	"go.vocdoni.io/dvote/chain/ethereumhandler"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/oracle"
	"go.vocdoni.io/dvote/router"
	"go.vocdoni.io/dvote/util"
	"go.vocdoni.io/proto/build/go/models"
)

type APIoracle struct {
	Namespace uint32
	oracle    *oracle.Oracle
	router    *router.Router
	eh        *ethereumhandler.EthereumHandler
}

func NewAPIoracle(o *oracle.Oracle, r *router.Router, chainName,
	web3Endpoint string) (*APIoracle, error) {
	a := &APIoracle{router: r, oracle: o}
	specs, err := chain.SpecsFor(chainName)
	if err != nil {
		return nil, err
	}
	a.eh, err = ethereumhandler.NewEthereumHandler(specs.Contracts, web3Endpoint)
	if err != nil {
		return nil, err
	}
	r.RegisterPublic("newERC20process", a.handleNewEthProcess)
	return a, nil
}

func (a *APIoracle) handleNewEthProcess(req router.RouterRequest) {
	var response api.MetaResponse
	if req.NewProcess == nil {
		a.router.SendError(req, "newProcess is empty")
		return
	}

	pid := fmt.Sprintf("%d%d%x%x",
		a.Namespace,
		req.NewProcess.StartBlock,
		req.NewProcess.EntityID,
		util.RandomBytes(32),
	)

	p := &models.Process{
		EntityId:     req.NewProcess.EntityID,
		StartBlock:   req.NewProcess.StartBlock,
		BlockCount:   req.NewProcess.BlockCount,
		CensusRoot:   req.NewProcess.CensusRoot,
		EnvelopeType: req.NewProcess.EnvelopeType,
		VoteOptions:  req.NewProcess.VoteOptions,
		EthIndexSlot: &req.NewProcess.EthIndexSlot,

		ProcessId:    ethereum.HashRaw([]byte(pid)),
		Status:       models.ProcessStatus_READY,
		Namespace:    a.Namespace,
		CensusOrigin: models.CensusOrigin_ERC20,
		Mode:         &models.ProcessMode{AutoStart: true},
	}

	if err := a.oracle.NewProcess(p); err != nil {
		a.router.SendError(req, err.Error())
		return
	}

	response.ProcessID = p.ProcessId
	if err := req.Send(a.router.BuildReply(req, &response)); err != nil {
		log.Warn(err)
	}
}

func (a *APIoracle) getIndexSlot(ctx context.Context, contractAddr []byte,
	evmBlockHeight uint64, rootHash []byte) (uint32, error) {
	// check valid storage root provided
	if len(contractAddr) != common.AddressLength {
		return 0, fmt.Errorf("contractAddress length is not correct")
	}
	addr := common.Address{}
	copy(addr[:], contractAddr[:])
	fetchedRoot, err := a.getStorageRoot(ctx, addr, evmBlockHeight)
	if err != nil {
		return 0, fmt.Errorf("cannot check EVM storage root: %w", err)
	}
	if bytes.Equal(fetchedRoot.Bytes(), common.Hash{}.Bytes()) {
		return 0, fmt.Errorf("invalid storage root obtained from Ethereum: %x", fetchedRoot)
	}
	if !bytes.Equal(fetchedRoot.Bytes(), rootHash) {
		return 0, fmt.Errorf("invalid storage root, got: %x expected: %x",
			fetchedRoot, rootHash)
	}
	// get index slot from the token storage proof contract
	islot, err := a.eh.GetTokenBalanceMappingPosition(ctx, addr)
	if err != nil {
		return 0, fmt.Errorf("cannot get balance mapping position from the contract: %w", err)
	}
	return uint32(islot.Uint64()), nil
}

// getStorageRoot returns the storage Root Hash of the Ethereum Patricia trie for
// a contract address and a block height
func (a *APIoracle) getStorageRoot(ctx context.Context, contractAddr common.Address,
	blockNum uint64) (hash common.Hash, err error) {
	// create token storage proof artifact
	ts := token.ERC20Token{
		RPCcli: a.eh.EthereumRPC,
		Ethcli: a.eh.EthereumClient,
	}
	ts.Init(ctx, "", contractAddr.String())

	// get block
	blk, err := ts.GetBlock(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return common.Hash{}, fmt.Errorf("cannot get block: %w", err)
	}

	// get proof, we use a random token holder and a dummy index slot
	holder := ethereum.NewSignKeys()
	if err := holder.Generate(); err != nil {
		return common.Hash{}, fmt.Errorf("cannot check storage root, cannot generate random Ethereum address: %w", err)
	}

	log.Debugf("get EVM storage root for address %s and block %d", holder.Address(), blockNum)
	sproof, err := ts.GetProofWithIndexSlot(ctx, holder.Address(), blk, 1)
	if err != nil {
		return common.Hash{}, fmt.Errorf("cannot get storage root: %w", err)
	}

	// return the storage root hash
	return sproof.StorageHash, nil
}
