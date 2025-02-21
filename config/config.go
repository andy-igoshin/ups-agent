package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/configor"
)

type Config struct {
	Debug    bool
	Status   bool
	Hostname string
	UpsName  string
	Model    string

	BatteryLevel int           `default:"10"`
	Minutes      time.Duration `default:"10m"`

	Logger struct {
		LogDir   string `default:"/var/log"`
		FileName string `default:"ups-agent.log"`
	}

	Snmp struct {
		Addr      string
		Community string        `default:"public"`
		PollTime  time.Duration `default:"60s"`
	}
}

func New() *Config {

	config := new(Config)

	var configfile string
	var configdump bool
	var configtest bool
	var debug bool
	var status bool

	flag.StringVar(&configfile, "config", "/etc/ups-agent.yml", "Set configuration file")
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
