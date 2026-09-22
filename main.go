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
		if err := status.Print(conf); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	if conf.IsCardConfig() {
		if err := cardconfig.Run(conf); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	logger, err := log.New(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer logger.Close()

	trapc := make(chan string, 16)

	go monitor.Run(conf, logger, trapc)

	logger.Info("UPS Agent is started")

	waitSignals(logger)

	logger.Info("UPS Agent is stopped")
}
