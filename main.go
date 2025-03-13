package main

import (
	"sync"

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
	var wg sync.WaitGroup

	go monitor.Run(conf, logger, trapc, &wg)

	logger.Info("UPS Agent is started")

	waitSignals(logger)

	wg.Wait()

	logger.Info("UPS Agent is stopped")
}
