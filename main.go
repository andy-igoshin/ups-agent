package main

import (
	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/monitor"
	"git.vsu.ru/ai/ups-agent/status"
)

func main() {

	conf := config.New()

	if conf.Status {
		status.Print(conf)
		return
	}

	logger := logNew(conf)
	defer logClose()

	trapc := make(chan string, 5)

	go monitor.Run(conf, logger, trapc)

	logger.Print("UPS Agent is started")

	waitSignals(conf, logger)

	logger.Print("UPS Agent is stopped")
}
