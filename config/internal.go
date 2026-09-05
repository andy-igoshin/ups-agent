package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func cliCardConfigType(o string) (t int) {
	switch o {
	case "get":
		t = CliCardConfigGet
	case "set":
		t = CliCardConfigSet
	case "default":
		t = CliCardConfigDefault
	default:
		fmt.Println("cardconfig: unknown operation:", o)
		os.Exit(1)
	}
	return
}

func dumpConfig(config *Config) {
	yml, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("config:", err)
	} else {
		fmt.Println(string(yml))
	}
}

func poweroffCmd(c string) string {
	cmd := strings.TrimSpace(c)
	if len(cmd) == 0 {
		fmt.Println("config: poweroffcmd() cannot be empty")
		os.Exit(1)
	}
	_, err := os.Stat(cmd)
	if err == nil {
		return cmd
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("config: poweroffcmd(%s) file not exists\n", cmd)
	} else {
		fmt.Printf("config: poweroffcmd(%s) error: %v\n", cmd, err)
	}
	os.Exit(1)
	return ""
}

func validate(c *Config) error {
	if c.Snmp.Addr == "" {
		return errors.New("snmp.addr is empty")
	}

	if net.ParseIP(c.Snmp.Addr) == nil {
		return fmt.Errorf("snmp.addr is not a valid IP address: %s", c.Snmp.Addr)
	}

	if c.Snmp.PollTime <= 0 {
		return fmt.Errorf("snmp.polltime must be positive, got %s", c.Snmp.PollTime)
	}

	if c.PowerCtl.BatteryLevel < 0 || c.PowerCtl.BatteryLevel > 100 {
		return fmt.Errorf("powerctl.batterylevel must be between 0 and 100, got %d",
			c.PowerCtl.BatteryLevel)
	}

	if c.PowerCtl.Minutes < 0 {
		return fmt.Errorf("powerctl.minutes must be >= 0, got %s", c.PowerCtl.Minutes)
	}

	if c.PowerCtl.Timeout < 0 {
		return fmt.Errorf("powerctl.timeout must be >= 0, got %s", c.PowerCtl.Timeout)
	}

	if c.Snmp.TrapListener != "" {
		if _, _, err := net.SplitHostPort(c.Snmp.TrapListener); err != nil {
			return fmt.Errorf("snmp.traplistener is invalid: %w", err)
		}
	}

	return nil
}
