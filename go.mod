module gitlab.com/vocdoni/go-dvote

go 1.14

// For testing purposes while dvote-protobuf becomes stable
// replace github.com/vocdoni/dvote-protobuf  => ../dvote-protobuf

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/benbjohnson/clock v1.1.0 // indirect
	github.com/cosmos/iavl v0.15.0-rc5
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/deroproject/graviton v0.0.0-20200906044921-89e9e09f9601
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/ethereum/go-ethereum v1.9.24
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/gobuffalo/packr/v2 v2.8.1
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/go-cmp v0.5.2
	github.com/google/gopacket v1.1.19 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/iden3/go-iden3-core v0.0.8-0.20200325104031-1ed04a261b78
	github.com/iden3/go-iden3-crypto v0.0.4
	github.com/ipfs/go-bitswap v0.3.3 // indirect
	github.com/ipfs/go-blockservice v0.1.4 // indirect
	github.com/ipfs/go-filestore v1.0.0 // indirect
	github.com/ipfs/go-graphsync v0.5.1 // indirect
	github.com/ipfs/go-ipfs v0.7.0
	github.com/ipfs/go-ipfs-blockstore v1.0.3 // indirect
	github.com/ipfs/go-ipfs-config v0.10.0
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipld-cbor v0.0.5 // indirect
	github.com/ipfs/go-log v1.0.4
	github.com/ipfs/interface-go-ipfs-core v0.4.0
	github.com/klauspost/compress v1.11.1
	github.com/koron/go-ssdp v0.0.2 // indirect
	github.com/libp2p/go-libp2p v0.12.0
	github.com/libp2p/go-libp2p-asn-util v0.0.0-20201026210036-4f868c957324 // indirect
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.7.0
	github.com/libp2p/go-libp2p-discovery v0.5.0
	github.com/libp2p/go-libp2p-gostream v0.3.0 // indirect
	github.com/libp2p/go-libp2p-kad-dht v0.11.0
	github.com/libp2p/go-libp2p-noise v0.1.2 // indirect
	github.com/libp2p/go-libp2p-pubsub-router v0.4.0 // indirect
	github.com/libp2p/go-libp2p-quic-transport v0.9.2 // indirect
	github.com/libp2p/go-libp2p-yamux v0.4.1 // indirect
	github.com/libp2p/go-netroute v0.1.4 // indirect
	github.com/libp2p/go-reuseport v0.0.2
	github.com/libp2p/go-sockaddr v0.1.0 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/prometheus/client_golang v1.8.0
	github.com/recws-org/recws v1.2.2
	github.com/rogpeppe/rjson v0.0.0-20151026200957-77220b71d327
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.0
	github.com/tendermint/tm-db v0.6.3
	github.com/vocdoni/dvote-protobuf v0.1.1
	github.com/vocdoni/eth-storage-proof v0.1.4-0.20201128112323-de7513ce5e25
	github.com/whyrusleeping/cbor-gen v0.0.0-20200826160007-0b9f6c5fb163 // indirect
	gitlab.com/vocdoni/go-external-ip v0.0.0-20190919225616-59cf485d00da
	go.opencensus.io v0.22.5 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/sys v0.0.0-20201130171929-760e229fe7c5 // indirect
	golang.org/x/text v0.3.4
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
	nhooyr.io/websocket v1.8.6
)

// Duktape is very slow to build, and can't be built with multiple cores since
// it includes a lot of C in a single file. Until
// https://github.com/ethereum/go-ethereum/issues/20590 is fixed, stub it out
// with a replace directive. The stub was hacked together with vim.
replace gopkg.in/olebedev/go-duktape.v3 => ./duktape-stub
