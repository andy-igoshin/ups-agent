package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func cliCardConfigType(o string) (int, error) {
	switch o {
	case "get":
		return CliCardConfigGet, nil
	case "set":
		return CliCardConfigSet, nil
	case "default":
		return CliCardConfigDefault, nil
	default:
		return CliUnknown, fmt.Errorf("cardconfig: unknown operation: %s", o)
	}
}

func Dump(config *Config) error {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	fmt.Println(string(yml))
	return nil
}

func poweroffCmd(c string) (string, error) {
	cmd := strings.TrimSpace(c)
	if len(cmd) == 0 {
		return "", errors.New("config: poweroffcmd() cannot be empty")
	}
	_, err := os.Stat(cmd)
	if err == nil {
		return cmd, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("config: poweroffcmd(%s) file not exists", cmd)
	}
	return "", fmt.Errorf("config: poweroffcmd(%s) error: %w", cmd, err)
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
