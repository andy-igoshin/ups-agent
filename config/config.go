package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/configor"
)

const (
	cliUnknown = iota
	cliStatus
	cliUpsConfigGet
	cliUpsConfigSet
	cliUpsConfigDefault
)

type Config struct {
	Hostname string
	UpsName  string
	Model    string

	Logger struct {
		LogDir   string `default:"logs"`
		FileName string `default:"ups-agent.log"`
		Level    string `default:"info"`
	}

	PowerCtl struct {
		Poweroff     bool          `default:"true"`
		PoweroffCmd  string        `default:"/sbin/poweroff"`
		BatteryLevel int           `default:"10"`
		Minutes      time.Duration `default:"10m"`
		Timeout      time.Duration `default:"0s"`
		UpsPoweroff  bool
	}

	Snmp struct {
		Addr          string
		Community     string        `default:"public"`
		PollTime      time.Duration `default:"60s"`
		TrapCommunity string        `default:"public"`
		TrapListener  string
	}

	UpsConfig struct {
		RatedOutputVoltage int
		HighTransferVolt   int
		LowTransferVolt    int
		Alarm              int
		AlarmTimer         time.Duration `default:"-1s"`
		MinReturnCapacity  int           `default:"-1"`
		Sensitivity        int
		LowBatteryRunTime  time.Duration
		ReturnDelay        time.Duration `default:"-1s"`
		ShutoffDelay       time.Duration
		UpsSleepTime       time.Duration `default:"-1s"`
	}

	cliType int
}

func New() *Config {

	config := new(Config)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var configfile string
	var configdump bool
	var configtest bool
	var status bool
	var upsconfig string

	flag.StringVar(&configfile, "config", filepath.Join(dir, "ups-agent.yml"), "Set configuration file")
	flag.BoolVar(&configdump, "configdump", false, "Perform configuration file dump")
	flag.BoolVar(&configtest, "configtest", false, "Perform configuration file test")
	flag.BoolVar(&status, "status", false, "Print UPS status")
	flag.StringVar(&upsconfig, "upsconfig", "", "`get|set|default` UPS configuration")
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
		config.cliType = cliStatus
	}

	if len(upsconfig) > 0 {
		config.cliType = cliUpsConfigType(upsconfig)
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

func (c *Config) IsStatus() (t bool) {
	if c.cliType == cliStatus {
		t = true
	}
	return
}

func (c *Config) IsUpsConfig() (t bool) {
	if c.cliType >= cliUpsConfigGet {
		t = true
	}
	return
}

func cliUpsConfigType(o string) (t int) {
	switch o {
	case "get":
		t = cliUpsConfigGet
	case "set":
		t = cliUpsConfigSet
	case "default":
		t = cliUpsConfigDefault
	default:
		fmt.Println("upsconfig: unknown operation:", o)
		os.Exit(1)
	}
	return
}
