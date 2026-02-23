package main

import (
	"git.vsu.ru/ai/ups-agent/cardconfig"
	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
	"git.vsu.ru/ai/ups-agent/monitor"
	"git.vsu.ru/ai/ups-agent/status"
)

func main() {

	conf := config.New()

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

	trapc := make(chan string)

	go monitor.Run(conf, logger, trapc)

	logger.Info("UPS Agent is started")

	waitSignals(logger)

	logger.Info("UPS Agent is stopped")
}
