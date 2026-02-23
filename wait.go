package main

import (
	"os"
	"os/signal"
	"syscall"

	"git.vsu.ru/ai/ups-agent/log"
)

func waitSignals(logger *log.Logger) {
	var s os.Signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s = <-c
		if s == syscall.SIGHUP {
			logger.Info("Got signal: %v", s)
			logger.Reopen()
		} else {
			break
		}
	}
	logger.Info("Got signal: %v, exiting", s)
}
