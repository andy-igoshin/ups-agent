package cardconfig

import (
//        "fmt"
//        "strconv"
//        "strings"
//        "time"

        _ "git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
)

func Run(conf *config.Config) {

	switch conf.CardOp() {
	case config.CliCardConfigGet:
		printConfig()

	case config.CliCardConfigSet:
		writeConfig(conf)

	case config.CliCardConfigDefault:
		setDefaultConfig()
	}
}

func printConfig() {
}

func setDefaultConfig() {
}

func writeConfig(conf *config.Config) {
}
