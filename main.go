package main

import (
	"fmt"
	"os"

	"git.vsu.ru/ai/ups-agent/cardconfig"
	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
	"git.vsu.ru/ai/ups-agent/monitor"
	"git.vsu.ru/ai/ups-agent/status"
)

func main() {

	conf, err := config.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if conf.IsDump() {
		_ = config.Dump(conf)
		return
	}
	if conf.IsTest() {
		fmt.Println("config: Syntax OK")
		return
	}

	if conf.IsStatus() {
		status.Print(conf)
		return
	}
	if conf.IsCardConfig() {
		cardconfig.Run(conf)
		return
	}

	logger := log.New(conf)
	defer logger.Close()

	trapc := make(chan string, 16)

	go monitor.Run(conf, logger, trapc)

	logger.Info("UPS Agent is started")

	waitSignals(logger)

	logger.Info("UPS Agent is stopped")
}
