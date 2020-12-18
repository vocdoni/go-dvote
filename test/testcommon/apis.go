package testcommon

import (
	"fmt"
	"math/rand"
	"testing"

	"go.vocdoni.io/dvote/census"
	"go.vocdoni.io/dvote/censustree/gravitontree"
	"go.vocdoni.io/dvote/config"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/data"
	dnet "go.vocdoni.io/dvote/net"
	"go.vocdoni.io/dvote/router"

	"go.vocdoni.io/dvote/types"
)

// DvoteAPIServer contains all the required pieces for running a go-dvote api server
type DvoteAPIServer struct {
	Signer         *ethereum.SignKeys
	VochainCfg     *config.VochainCfg
	CensusDir      string
	IpfsDir        string
	ScrutinizerDir string
	PxyAddr        string
	Storage        data.Storage
	IpfsPort       int
}

/*
Start starts a basic dvote server
1. Create signing key
2. Starts the Proxy
3. Starts the IPFS storage
4. Starts the Census Manager
5. Starts the Vochain miner if vote api enabled
6. Starts the Dvote API router if enabled
7. Starts the scrutinizer service and API if enabled
*/
func (d *DvoteAPIServer) Start(tb testing.TB, apis ...string) {
	// create signer
	d.Signer = ethereum.NewSignKeys()
	d.Signer.Generate()

	// create the proxy to handle HTTP queries
	pxy := NewMockProxy(tb)
	d.PxyAddr = fmt.Sprintf("ws://%s/dvote", pxy.Addr)

	// Create WebSocket endpoint
	ws := new(dnet.WebsocketHandle)
	ws.Init(new(types.Connection))
	ws.SetProxy(pxy)

	// Create the listener for routing messages
	listenerOutput := make(chan types.Message)
	ws.Listen(listenerOutput)

	// Create the API router
	var err error
	d.IpfsDir = tb.TempDir()
	ipfsStore := data.IPFSNewConfig(d.IpfsDir)
	ipfs := data.IPFSHandle{}
	d.IpfsPort = 14000 + rand.Intn(2048)
	if err = ipfs.SetMultiAddress(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", d.IpfsPort)); err != nil {
		tb.Fatal(err)
	}
	if err = ipfs.Init(ipfsStore); err != nil {
		tb.Fatal(err)
	}
	d.Storage = &ipfs
	tb.Cleanup(func() {
		if err := d.Storage.Stop(); err != nil {
			tb.Error(err)
		}
	})

	routerAPI := router.InitRouter(listenerOutput, d.Storage, d.Signer, nil, true)

	// Create the Census Manager and enable it trough the router
	var cm census.Manager
	d.CensusDir = tb.TempDir()

	if err := cm.Init(d.CensusDir, "", gravitontree.NewTree); err != nil {
		tb.Fatal(err)
	}

	for _, api := range apis {
		switch api {
		case "file":
			routerAPI.EnableFileAPI()
		case "census":
			routerAPI.EnableCensusAPI(&cm)
		case "vote":
			vnode := NewMockVochainNode(tb, d)
			sc := NewMockScrutinizer(tb, d, vnode)
			routerAPI.Scrutinizer = sc
			routerAPI.EnableVoteAPI(vnode, nil)
		default:
			tb.Fatalf("unknown api: %q", api)
		}
	}

	go routerAPI.Route()
	ws.AddProxyHandler("/dvote")
}

// NewMockProxy creates a new testing proxy with predefined valudes
func NewMockProxy(tb testing.TB) *dnet.Proxy {
	pxy := dnet.NewProxy()
	pxy.C.Address = "127.0.0.1"
	pxy.C.Port = 0
	err := pxy.Init()
	if err != nil {
		tb.Fatal(err)
	}
	return pxy
}
