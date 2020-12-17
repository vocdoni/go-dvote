package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go.vocdoni.io/dvote/chain"
	"go.vocdoni.io/dvote/config"
	"go.vocdoni.io/dvote/log"
)

type ethereumStandaloneCfgWrapper struct {
	w3Config  *config.W3Cfg  // node web3 config
	ethConfig *config.EthCfg // node config
	dataDir   string
	logLevel  string
	logOutput string
}

func newConfig() (*ethereumStandaloneCfgWrapper, config.Error) {
	var err error
	var cfgError config.Error
	// create base config
	ethereumCfgWrapper := &ethereumStandaloneCfgWrapper{
		w3Config:  new(config.W3Cfg),
		ethConfig: new(config.EthCfg),
	}
	// get current user home dir
	home, err := os.UserHomeDir()
	if err != nil {
		cfgError = config.Error{
			Critical: true,
			Message:  fmt.Sprintf("cannot get user home directory with error: %s", err),
		}
		return nil, cfgError
	}

	// CLI flags will be used if something fails from this point
	// CLI flags have preference over the config file

	// general
	flag.StringVar(&ethereumCfgWrapper.dataDir, "dataDir", home+"/.dvote", "directory where data is stored")
	ethereumCfgWrapper.logLevel = *flag.String("logLevel", "info", "Log level (debug, info, warn, error, fatal)")
	ethereumCfgWrapper.logOutput = *flag.String("logOutput", "stdout", "Log output (stdout, stderr or filepath)")
	// ethereum node
	ethereumCfgWrapper.ethConfig.SigningKey = *flag.String("ethSigningKey", "", "signing private Key (if not specified the Ethereum keystore will be used)")
	ethereumCfgWrapper.ethConfig.ChainType = *flag.String("ethChain", "goerli", fmt.Sprintf("Ethereum blockchain to use: %s", chain.AvailableChains))
	ethereumCfgWrapper.ethConfig.LightMode = *flag.Bool("ethChainLightMode", true, "synchronize Ethereum blockchain in light mode")
	ethereumCfgWrapper.ethConfig.NodePort = *flag.Int("ethNodePort", 30303, "Ethereum p2p node port to use")
	ethereumCfgWrapper.ethConfig.DataDir = ethereumCfgWrapper.dataDir + "/ethereum"
	// ethereum web3
	ethereumCfgWrapper.w3Config.RPCPort = *flag.Int("w3RPCPort", 9091, "web3 RPC port")
	ethereumCfgWrapper.w3Config.RPCHost = *flag.String("w3RPCHost", "127.0.0.1", "web3 RPC host")

	// parse flags
	flag.Parse()
	// setting up viper
	viper := viper.New()
	viper.AddConfigPath(ethereumCfgWrapper.dataDir)
	viper.SetConfigName("web3-standalone")
	viper.SetConfigType("yml")
	// binding flags to viper
	// general
	viper.BindPFlag("dataDir", flag.Lookup("dataDir"))
	viper.BindPFlag("logLevel", flag.Lookup("logLevel"))
	viper.BindPFlag("logOutput", flag.Lookup("logOutput"))
	// ethereum node
	viper.BindPFlag("ethConfig.signingKey", flag.Lookup("ethSigningKey"))
	viper.BindPFlag("ethConfig.chainType", flag.Lookup("ethChain"))
	viper.BindPFlag("ethConfig.lightMode", flag.Lookup("ethChainLightMode"))
	viper.BindPFlag("ethConfig.nodePort", flag.Lookup("ethNodePort"))
	// ethereum web3
	viper.BindPFlag("w3Config.RPCPort", flag.Lookup("w3RPCPort"))
	viper.BindPFlag("w3Config.RPCHost", flag.Lookup("w3RPCHost"))

	// check if config file exists
	_, err = os.Stat(ethereumCfgWrapper.dataDir + "/web3-standalone.yml")
	if os.IsNotExist(err) {
		cfgError = config.Error{
			Message: fmt.Sprintf("creating new config file in %s", ethereumCfgWrapper.dataDir),
		}
		// creting config folder if not exists
		err = os.MkdirAll(ethereumCfgWrapper.dataDir, os.ModePerm)
		if err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot create data directory (%s)", err),
			}
		}
		// create config file if not exists
		if err := viper.SafeWriteConfig(); err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot write config file into config dir (%s)", err),
			}
		}
	} else {
		// read config file
		err = viper.ReadInConfig()
		if err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot read loaded config file in %s (%s)", err, ethereumCfgWrapper.dataDir),
			}
		}
	}
	err = viper.Unmarshal(&ethereumCfgWrapper)
	if err != nil {
		cfgError = config.Error{
			Message: fmt.Sprintf("cannot unmarshal loaded config file (%s)", err),
		}
	}

	return ethereumCfgWrapper, cfgError
}

/*
Example code for using web3 implementation

Testing the RPC can be performed with curl and/or websocat
 curl -X POST -H "Content-Type:application/json" --data '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":74}' localhost:9091
 echo '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":74}' | websocat ws://0.0.0.0:9092
*/
func main() {
	// setup config
	globalCfg, cfgErr := newConfig()
	if globalCfg == nil {
		panic("cannot read configuration")
	}
	fmt.Println(globalCfg.logLevel)
	log.Init(globalCfg.logLevel, globalCfg.logOutput)

	log.Debugf("initializing gateway config %+v", globalCfg)

	// check if errors during config creation and determine if Critical
	if cfgErr.Critical && cfgErr.Message != "" {
		log.Fatalf("Critical error loading config: %s", cfgErr.Message)
	} else if !cfgErr.Critical && cfgErr.Message != "" {
		log.Warnf("non Critical error loading config: %s", cfgErr.Message)
	} else if !cfgErr.Critical && cfgErr.Message == "" {
		log.Infof("config file loaded successfully, remember CLI flags have preference")
	}

	log.Info("starting gateway")

	cfg, err := chain.NewConfig(globalCfg.ethConfig, globalCfg.w3Config)
	if err != nil {
		log.Fatal(err)
	}

	node, err := chain.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	node.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)

	for {
		time.Sleep(1 * time.Second)
	}
}
