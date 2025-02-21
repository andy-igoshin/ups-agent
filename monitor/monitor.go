package monitor

import (
	"fmt"
	"log"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
)

func Run(conf *config.Config, logger *log.Logger, trapr <-chan string) {

	onBattery := false

	for {

		select {
		case trap := <-trapr:
			fmt.Println("has trap: ", trap)
			logger.Print("has trap: ", trap)
		case <-time.After(conf.Snmp.PollTime):
			fmt.Println("next iteration")
			logger.Print("next iteration")
		}

		md, err := card.Monitor(conf.Snmp.Addr, conf.Snmp.Community)
		if err != nil {
			log.Print(err)
			continue
		}

fmt.Println(md)

		switch md.UpsBaseOutputStatus {
		case card.UpsBaseOutputStatusOnline:
			if onBattery {
				onBattery = false
				logger.Print("The UPS output is online")
			}
			continue
		case card.UpsBaseOutputStatusOnBattery:
			onBattery = true
			// Send Notify to email and system terminal broadcast if Alert timer
			logger.Print("The UPS switched to battery")
		default:
			// Send Notify
			continue
		}

		if md.UpsBaseBatteryStatus == card.UpsBaseBatteryStatusLow {
			// Shutdown
		}

	}
}
