package config

import (
	"errors"
	"fmt"
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
