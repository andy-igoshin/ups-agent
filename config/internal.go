package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func dumpConfig(config *Config) {
	yml, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("config:", err)
	} else {
		fmt.Println(string(yml))
	}
}
