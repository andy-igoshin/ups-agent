package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/configor"
)

type Config struct {
	Hostname string
	UpsName  string
	Model    string

	StatusCli bool

	Logger struct {
		LogDir   string `default:"/var/log/ups-agent"`
		FileName string `default:"ups-agent.log"`
		Level    string `default:"info"`
	}

	PowerCtl struct {
		Poweroff     bool
		PoweroffCmd  string        `default:"/sbin/poweroff"`
		BatteryLevel int           `default:"10"`
		Minutes      time.Duration `default:"10m"`
		Timeout      time.Duration
	}

	Snmp struct {
		Addr          string
		Community     string        `default:"public"`
		PollTime      time.Duration `default:"60s"`
		TrapCommunity string        `default:"public"`
		TrapListener  string
	}
}

func New() *Config {

	config := new(Config)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var configfile string
	var configdump bool
	var configtest bool
	var status bool

	flag.StringVar(&configfile, "config", filepath.Join(dir, "ups-agent.yml"), "Set configuration file")
	flag.BoolVar(&configdump, "configdump", false, "Perform configuration file dump")
	flag.BoolVar(&configtest, "configtest", false, "Perform configuration file test")
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

	if status {
		config.StatusCli = status
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
