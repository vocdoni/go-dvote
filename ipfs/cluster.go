package ipfs

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	//crdt "github.com/ipfs/go-ds-crdt"
	clusterconfig "github.com/ipfs/ipfs-cluster/config"

	ipfscluster "github.com/ipfs/ipfs-cluster"
	"github.com/ipfs/ipfs-cluster/allocator/ascendalloc"
	"github.com/ipfs/ipfs-cluster/allocator/descendalloc"
	"github.com/ipfs/ipfs-cluster/api/ipfsproxy"
	"github.com/ipfs/ipfs-cluster/api/rest"

	//	"github.com/ipfs/ipfs-cluster/consensus/crdt"

	"github.com/ipfs/ipfs-cluster/consensus/raft"
	"github.com/ipfs/ipfs-cluster/informer/disk"
	"github.com/ipfs/ipfs-cluster/informer/numpin"
	"github.com/ipfs/ipfs-cluster/ipfsconn/ipfshttp"
	"github.com/ipfs/ipfs-cluster/monitor/pubsubmon"
	"github.com/ipfs/ipfs-cluster/observations"
	"github.com/ipfs/ipfs-cluster/pintracker/maptracker"
	"github.com/ipfs/ipfs-cluster/pintracker/stateless"

	"go.opencensus.io/tag"

	ds "github.com/ipfs/go-datastore"
	host "github.com/libp2p/go-libp2p-host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	peer "github.com/libp2p/go-libp2p-peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"

	config "gitlab.com/vocdoni/go-dvote/config"
	"gitlab.com/vocdoni/go-dvote/log"

	errors "github.com/pkg/errors"
)

//const programName = `dvote-ipfs-cluster-service`
var ProgramName string

var (
	// DefaultFolder is the name of the cluster folder
	DefaultFolder = ".ipfs-cluster"
	// DefaultPath is set on init() to $HOME/DefaultFolder
	// and holds all the ipfs-cluster data
	DefaultPath string
	// The name of the configuration file inside DefaultPath
	DefaultConfigFile = "service.json"
	// The name of the identity file inside DefaultPath
	DefaultIdentityFile = "identity.json"
)

var (
	configPath   string
	identityPath string
)

func checkErr(doing string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			doing = fmt.Sprintf(doing, args...)
		}
		log.Error(doing, err)
		err = locker.tryUnlock()
		if err != nil {
			log.Errorf("error releasing execution lock: %s\n", err)
		}
		os.Exit(1)
	}
}

func parseBootstraps(flagVal []string) (bootstraps []ma.Multiaddr) {
	for _, a := range flagVal {
		bAddr, err := ma.NewMultiaddr(a)
		checkErr("error parsing bootstrap multiaddress (%s)", err, a)
		bootstraps = append(bootstraps, bAddr)
	}
	return
}

func InitCluster(path, configFile, idFile string, c config.ClusterCfg) error {
	DefaultFolder = path
	DefaultConfigFile = path + "/" + configFile
	DefaultIdentityFile = path + "/" + idFile
	configPath = DefaultConfigFile
	identityPath = DefaultIdentityFile

	makeConfigFolder()
	//TODO:: check if id exists first!
	id, err := clusterconfig.NewIdentity()
	if err != nil {
		return err
	}
	id.ID = c.PeerID
	id.PrivateKey = c.Private
	err = id.SaveJSON(identityPath)
	if err != nil {
		return err
	}

	// //TODO:: check if config exists first!
	mgr, _ := makeConfigs()
	mgr.Default()
	saveConfig(mgr)
	return nil
}

// Runs the cluster
func RunCluster(c config.ClusterCfg, ch chan *ipfscluster.Cluster) (error) {
	/*
		currUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		DefaultPath = currUser.HomeDir + "/" + DefaultFolder
	*/
	locker = &lock{path: DefaultPath}

	log.Infof("cluster initializing. Please wait...")

	ctx, cancel := context.WithCancel(context.Background())

	bootstraps := parseBootstraps(c.Bootstraps)

	// Execution lock
	locker.lock()
	defer locker.tryUnlock()
	log.Infof("execution lock established.")

	// Load all the configurations and identity
	cfgMgr, ident, cfgs := makeAndLoadConfigs(c)
	log.Infof("config loaded")

	defer cfgMgr.Shutdown()

	if c.Stats {
		cfgs.metricsCfg.EnableStats = true
	}

	cfgs = propagateTracingConfig(ident, cfgs, c.Tracing)

	// Cleanup state if bootstrapping
	raftStaging := false
	if len(bootstraps) > 0 && c.Consensus == "raft" {
		raft.CleanupRaft(cfgs.raftCfg)
		raftStaging = true
	}

	if c.Leave {
		cfgs.clusterCfg.LeaveOnShutdown = true
	}

	log.Infof("creating host")
	host, pubsub, dht, err := ipfscluster.NewClusterHost(ctx, ident, cfgs.clusterCfg)
	checkErr("creating libp2p host", err)

	log.Infof("Creating cluster with key: %s", hex.EncodeToString(cfgs.clusterCfg.Secret))
	cluster, err := createCluster(ctx, &c, host, pubsub, dht, ident, cfgs, raftStaging)
	ch <- cluster
	checkErr("starting cluster", err)
	log.Debug("created cluster")

	// noop if no bootstraps
	// if bootstrapping fails, consensus will never be ready
	// and timeout. So this can happen in background and we
	// avoid worrying about error handling here (since Cluster
	// will realize).

	go bootstrap(ctx, cluster, bootstraps)

	return handleSignals(ctx, cancel, cluster, host, dht)
}

// createCluster creates all the necessary things to produce the cluster
// object and returns it along the datastore so the lifecycle can be handled
// (the datastore needs to be Closed after shutting down the Cluster).
func createCluster(
	ctx context.Context,
	c *config.ClusterCfg,
	host host.Host,
	pubsub *pubsub.PubSub,
	dht *dht.IpfsDHT,
	ident *clusterconfig.Identity,
	cfgs *cfgs,
	raftStaging bool,
) (*ipfscluster.Cluster, error) {
	ctx, err := tag.New(ctx, tag.Upsert(observations.HostKey, host.ID().Pretty()))
	checkErr("tag context with host id", err)

	api, err := rest.NewAPIWithHost(ctx, cfgs.apiCfg, host)
	checkErr("creating REST API component", err)

	proxy, err := ipfsproxy.New(cfgs.ipfsproxyCfg)
	checkErr("creating IPFS Proxy component", err)

	apis := []ipfscluster.API{api, proxy}

	connector, err := ipfshttp.NewConnector(cfgs.ipfshttpCfg)
	checkErr("creating IPFS Connector component", err)

	tracker := setupPinTracker(
		c.PinTracker,
		host,
		cfgs.maptrackerCfg,
		cfgs.statelessTrackerCfg,
		cfgs.clusterCfg.Peername,
	)

	informer, alloc := setupAllocation(
		c.Alloc,
		cfgs.diskInfCfg,
		cfgs.numpinInfCfg,
	)

	ipfscluster.ReadyTimeout = cfgs.raftCfg.WaitForLeaderTimeout + 5*time.Second

	err = observations.SetupMetrics(cfgs.metricsCfg)
	checkErr("setting up Metrics", err)

	tracer, err := observations.SetupTracing(cfgs.tracingCfg)
	checkErr("setting up Tracing", err)

	store := setupDatastore(c.Consensus, ident, cfgs)

	cons, err := setupConsensus(
		c.Consensus,
		host,
		dht,
		pubsub,
		cfgs,
		store,
		raftStaging,
	)
	if err != nil {
		store.Close()
		checkErr("setting up Consensus", err)
	}

	var peersF func(context.Context) ([]peer.ID, error)
	if c.Consensus == "raft" {
		peersF = cons.Peers
	}

	mon, err := pubsubmon.New(ctx, cfgs.pubsubmonCfg, pubsub, peersF)
	if err != nil {
		store.Close()
		checkErr("setting up PeerMonitor", err)
	}

	return ipfscluster.NewCluster(
		ctx,
		host,
		dht,
		cfgs.clusterCfg,
		store,
		cons,
		apis,
		connector,
		tracker,
		mon,
		alloc,
		informer,
		tracer,
	)
}

// bootstrap will bootstrap this peer to one of the bootstrap addresses
// if there are any.
func bootstrap(ctx context.Context, cluster *ipfscluster.Cluster, bootstraps []ma.Multiaddr) {
	for _, bstrap := range bootstraps {
		log.Infof("Bootstrapping to %s", bstrap)
		err := cluster.Join(ctx, bstrap)
		if err != nil {
			log.Errorf("bootstrap to %s failed: %s", bstrap, err)
		}
	}
}

func handleSignals(
	ctx context.Context,
	cancel context.CancelFunc,
	cluster *ipfscluster.Cluster,
	host host.Host,
	dht *dht.IpfsDHT,
) error {
	signalChan := make(chan os.Signal, 20)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	var ctrlcCount int
	for {
		select {
		case <-signalChan:
			ctrlcCount++
			handleCtrlC(ctx, cluster, ctrlcCount)
		case <-cluster.Done():
			cancel()
			dht.Close()
			host.Close()
			return nil
		}
	}
}

func handleCtrlC(ctx context.Context, cluster *ipfscluster.Cluster, ctrlcCount int) {
	switch ctrlcCount {
	case 1:
		go func() {
			err := cluster.Shutdown(ctx)
			checkErr("shutting down cluster", err)
		}()
	case 2:
		log.Infof(`


!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Shutdown is taking too long! Press Ctrl-c again to manually kill cluster.
Note that this may corrupt the local cluster state.
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!


`)
	case 3:
		log.Infof("exiting cluster NOW")
		locker.tryUnlock()
		os.Exit(-1)
	}
}

func setupAllocation(
	name string,
	diskInfCfg *disk.Config,
	numpinInfCfg *numpin.Config,
) (ipfscluster.Informer, ipfscluster.PinAllocator) {
	switch name {
	case "disk", "disk-freespace":
		informer, err := disk.NewInformer(diskInfCfg)
		checkErr("creating informer", err)
		return informer, descendalloc.NewAllocator()
	case "disk-reposize":
		informer, err := disk.NewInformer(diskInfCfg)
		checkErr("creating informer", err)
		return informer, ascendalloc.NewAllocator()
	case "numpin", "pincount":
		informer, err := numpin.NewInformer(numpinInfCfg)
		checkErr("creating informer", err)
		return informer, ascendalloc.NewAllocator()
	default:
		err := errors.New("unknown allocation strategy")
		checkErr("", err)
		return nil, nil
	}
}

func setupPinTracker(
	name string,
	h host.Host,
	mapCfg *maptracker.Config,
	statelessCfg *stateless.Config,
	peerName string,
) ipfscluster.PinTracker {
	switch name {
	case "map":
		ptrk := maptracker.NewMapPinTracker(mapCfg, h.ID(), peerName)
		log.Debugf("map pintracker loaded")
		return ptrk
	case "stateless":
		ptrk := stateless.New(statelessCfg, h.ID(), peerName)
		log.Debugf("stateless pintracker loaded")
		return ptrk
	default:
		err := errors.New("unknown pintracker type")
		checkErr("", err)
		return nil
	}
}

func setupDatastore(
	consensus string,
	ident *clusterconfig.Identity,
	cfgs *cfgs,
) ds.Datastore {
	stmgr := newStateManager(consensus, ident, cfgs)
	store, err := stmgr.GetStore()
	checkErr("creating datastore", err)
	return store
}

func setupConsensus(
	name string,
	h host.Host,
	dht *dht.IpfsDHT,
	pubsub *pubsub.PubSub,
	cfgs *cfgs,
	store ds.Datastore,
	raftStaging bool,
) (ipfscluster.Consensus, error) {
	switch name {
	case "raft":
		rft, err := raft.NewConsensus(
			h,
			cfgs.raftCfg,
			store,
			raftStaging,
		)
		if err != nil {
			return nil, errors.Wrap(err, "creating Raft component")
		}
		return rft, nil
		/*
			case "crdt":
				convrdt, err := crdt.New(
					h,
					dht,
					pubsub,
					cfgs.crdtCfg,
					store,
				)
				if err != nil {
					return nil, errors.Wrap(err, "creating CRDT component")
				}
				return convrdt, nil
		*/
	default:
		return nil, errors.New("unknown consensus component")
	}
}
