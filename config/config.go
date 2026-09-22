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
	CliUnknown = iota
	CliCardConfigGet
	CliCardConfigSet
	CliCardConfigDefault
	CliStatus
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

	CardConfig struct {
		RatedOutputVoltage int           `default:"-1"`
		HighTransferVolt   int           `default:"-1"`
		LowTransferVolt    int           `default:"-1"`
		Alarm              int           `default:"-1"`
		AlarmTimer         time.Duration `default:"-1s"`
		MinReturnCapacity  int           `default:"-1"`
		Sensitivity        int           `default:"-1"`
		LowBatteryRunTime  time.Duration `default:"-1s"`
		ReturnDelay        time.Duration `default:"-1s"`
		ShutoffDelay       time.Duration `default:"-1s"`
		UpsSleepTime       time.Duration `default:"-1m"`
	}

	cliType int
	cliDump bool
	cliTest bool
}

func New() (*Config, error) {

	config := new(Config)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var configfile string
	var configdump bool
	var configtest bool
	var status bool
	var cardconfig string

	flag.StringVar(&cardconfig, "cardconfig", "", "Perform `get|set|default` UPS configuration")
	flag.StringVar(&configfile, "config", filepath.Join(dir, "ups-agent.yml"), "Set configuration file `path`")
	flag.BoolVar(&configdump, "configdump", false, "Perform configuration file dump")
	flag.BoolVar(&configtest, "configtest", false, "Perform configuration file test")
	flag.BoolVar(&status, "status", false, "Print UPS status")
	flag.Parse()

	if err := configor.Load(config, configfile); err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}

	if err := validate(config); err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}

	if len(config.Hostname) == 0 {
		hostname, err := os.Hostname()
		if err != nil {
			return nil, fmt.Errorf("config: %w", err)
		}
		config.Hostname = hostname
	}

	if cmd, err := poweroffCmd(config.PowerCtl.PoweroffCmd); err != nil {
		return nil, err
	} else {
		config.PowerCtl.PoweroffCmd = cmd
	}

	if status {
		config.cliType = CliStatus
	}

	if len(cardconfig) > 0 {
		t, err := cliCardConfigType(cardconfig)
		if err != nil {
			return nil, err
		}
		config.cliType = t
	}

	config.cliDump = configdump
	config.cliTest = configtest

	return config, nil
}

func (c *Config) CardOp() (op int) {
	switch c.cliType {
	case CliCardConfigGet, CliCardConfigSet, CliCardConfigDefault:
		op = c.cliType
	}
	return
}

func (c *Config) IsCardConfig() (t bool) {
	switch c.cliType {
	case CliCardConfigGet, CliCardConfigSet, CliCardConfigDefault:
		t = true
	}
	return
}

func (c *Config) IsStatus() (t bool) {
	if c.cliType == CliStatus {
		t = true
	}
	return
}

func (c *Config) IsDump() bool {
	return c.cliDump
}

func (c *Config) IsTest() bool {
	return c.cliTest
}
