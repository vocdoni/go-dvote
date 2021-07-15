package apioracle

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/storage-proofs-eth-go/token/mapbased"
	"go.vocdoni.io/dvote/api"
	"go.vocdoni.io/dvote/crypto/ethereum"
	chain "go.vocdoni.io/dvote/ethereum"
	ethereumhandler "go.vocdoni.io/dvote/ethereum/handler"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/oracle"
	"go.vocdoni.io/dvote/router"
	"go.vocdoni.io/dvote/util"
	"go.vocdoni.io/proto/build/go/models"
)

const (
	ethQueryTimeOut      = 20 * time.Second
	aclPurgePeriod       = 10 * time.Minute
	aclTimeWindow        = 5 * 24 * time.Hour // 5 days
	maxProposalPerWindow = 3
)

var srcNetworkIds = map[string]models.SourceNetworkId{
	"default":   models.SourceNetworkId_UNKNOWN,
	"mainnet":   models.SourceNetworkId_ETH_MAINNET_SIGNALING,
	"homestead": models.SourceNetworkId_ETH_MAINNET_SIGNALING,
	"rinkeby":   models.SourceNetworkId_ETH_RINKEBY_SIGNALING,
}

type APIoracle struct {
	Namespace        uint32
	oracle           *oracle.Oracle
	router           *router.Router
	eh               *ethereumhandler.EthereumHandler
	chainNames       map[string]bool
	erc20proposalACL *proposalACL
}

func NewAPIoracle(o *oracle.Oracle, r *router.Router) (*APIoracle, error) {
	a := &APIoracle{router: r, oracle: o, chainNames: make(map[string]bool)}
	return a, nil
}

func (a *APIoracle) EnableERC20(chainName string, web3Endpoints []string) error {
	if chainName == "" || len(web3Endpoints) == 0 {
		return fmt.Errorf("no web3 endpoint or chain name provided")
	}
	specs, err := chain.SpecsFor(chainName)
	if err != nil {
		return err
	}
	a.erc20proposalACL = NewProposalACL(aclTimeWindow, maxProposalPerWindow)
	srcNetId, ok := srcNetworkIds[chainName]
	if !ok {
		srcNetId = srcNetworkIds["default"]
	}
	// TODO: use multiple web3
	a.eh, err = ethereumhandler.NewEthereumHandler(specs.Contracts, srcNetId, web3Endpoints[0])
	if err != nil {
		return err
	}
	a.eh.WaitSync()
	for k, v := range srcNetworkIds {
		if v.String() == srcNetId.String() {
			a.chainNames[k] = true
		}
	}
	a.router.RegisterPublic("newERC20process", a.handleNewEthProcess)
	a.router.APIs = append(a.router.APIs, "oracle")
	return nil
}

func (a *APIoracle) handleNewEthProcess(req router.RouterRequest) {
	var response api.MetaResponse
	if req.NewProcess == nil {
		a.router.SendError(req, "newProcess is empty")
		return
	}
	if _, ok := a.chainNames[req.NewProcess.NetworkId]; !ok {
		a.router.SendError(req, fmt.Sprintf("provided chainId does not match ours (%s)",
			req.NewProcess.NetworkId))
		return
	}
	if req.NewProcess.EthIndexSlot == nil {
		a.router.SendError(req, "index slot not provided")
		return
	}
	if req.NewProcess.SourceHeight == nil {
		a.router.SendError(req, "no source height provided")
		return
	}

	pidseed := fmt.Sprintf("%d%d%x%x",
		a.Namespace,
		req.NewProcess.StartBlock,
		req.NewProcess.EntityID,
		util.RandomBytes(32),
	)

	p := &models.Process{
		EntityId:          req.NewProcess.EntityID,
		StartBlock:        req.NewProcess.StartBlock,
		BlockCount:        req.NewProcess.BlockCount,
		CensusRoot:        req.NewProcess.CensusRoot,
		EnvelopeType:      req.NewProcess.EnvelopeType,
		VoteOptions:       req.NewProcess.VoteOptions,
		EthIndexSlot:      req.NewProcess.EthIndexSlot,
		SourceBlockHeight: req.NewProcess.SourceHeight,
		Metadata:          &req.NewProcess.Metadata,
		ProcessId:         ethereum.HashRaw([]byte(pidseed)),
		Status:            models.ProcessStatus_READY,
		Namespace:         a.Namespace,
		CensusOrigin:      models.CensusOrigin_ERC20,
		Mode:              &models.ProcessMode{AutoStart: true},
		SourceNetworkId:   a.eh.SrcNetworkId,
		Owner:             req.GetAddress().Bytes(),
	}

	// Check the ACL
	if err := a.erc20proposalACL.add(p.Owner, p.EntityId); err != nil {
		a.router.SendError(req, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ethQueryTimeOut)
	defer cancel()
	index, err := a.getIndexSlot(ctx, p.EntityId, p.GetSourceBlockHeight(), p.CensusRoot)
	if err != nil {
		a.router.SendError(req, err.Error())
		return
	}
	log.Infof("fetched index slot %d for contract %x", index, p.EntityId)
	if index != p.GetEthIndexSlot() {
		a.router.SendError(req, "index slot does not match")
		return
	}

	if req.EthProof == nil {
		a.router.SendError(req, "storage proof is nil")
		return
	}
	err = mapbased.VerifyProof(*req.GetAddress(), common.BytesToHash(p.CensusRoot),
		*req.EthProof, int(index), new(big.Int).SetBytes(req.EthProof.Value), nil)

	if err != nil {
		a.router.SendError(req, fmt.Sprintf("proof is not valid: %v", err))
		return
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
	return a.eh.GetStorageRoot(ctx, contractAddr, big.NewInt(int64(blockNum)))
}
