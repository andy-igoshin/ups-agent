package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"git.vsu.ru/ai/ups-agent/config"
)

func waitSignals(conf *config.Config, logger *log.Logger) {
	var s os.Signal
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s = <-c
		if s == syscall.SIGHUP {
			logger.Printf("Got signal: %v", s)
			logReopen(conf, logger)
		} else {
			break
		}
	}
	logger.Printf("Got signal: %v, exiting", s)
}
