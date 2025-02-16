package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/configor"
)

type SnmpConfig struct {
	Addr      string
	Community string
	PollTime  time.Duration
}

type Config struct {
	Debug        bool
	Status       bool
	Hostname     string
	UpsName      string
	Model        string
	Snmp         SnmpConfig
	BatteryLevel int           `default:"10"`
	Minutes      time.Duration `default:"10m"`
}

func New() *Config {

	config := new(Config)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var configfile string
	var configdump bool
	var configtest bool
	var debug bool
	var status bool

	flag.StringVar(&configfile, "config", dir+"/ups-agent.yml", "Set configuration file")
	flag.BoolVar(&configdump, "configdump", false, "Perform configuration file dump")
	flag.BoolVar(&configtest, "configtest", false, "Perform configuration file test")
	flag.BoolVar(&debug, "debug", false, "Run in debug mode")
	flag.BoolVar(&status, "status", false, "Print UPS status")
	flag.Parse()

	err := configor.Load(config, configfile)
	if err != nil {
		fmt.Println("config:", err)
		os.Exit(1)
	}

	if len(config.Hostname) == 0 {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("config:", err)
			os.Exit(1)
		}
		config.Hostname = hostname
	}

	if debug {
		config.Debug = debug
	}

	if status {
		config.Status = status
	}

	if configdump {
		dumpConfig(config)
		os.Exit(0)
	}

	if configtest {
		fmt.Println("config: Syntax OK")
		os.Exit(0)
	}

	return config
}
