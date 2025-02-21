package main

import (
	"log"
	"os"
	"path"

	"git.vsu.ru/ai/ups-agent/config"
)

var logFh *os.File

func logClose() {
	logFh.Close()
}

func logNew(conf *config.Config) *log.Logger {

	logFh, err := os.OpenFile(path.Join(conf.Logger.LogDir, conf.Logger.FileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Open log error: ", err)
	}

	return log.New(logFh, "", log.LstdFlags|log.Lmicroseconds)
}

func logReopen(conf *config.Config, logH *log.Logger) {

	logFh.Close()
	file, err := os.OpenFile(path.Join(conf.Logger.LogDir, conf.Logger.FileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Reopen log error: %s", err)
	}

	logFh = file
	logH.SetOutput(file)
}
