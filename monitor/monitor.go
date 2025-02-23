package monitor

import (
	"sync"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

func Run(conf *config.Config, logger *log.Logger, trapc chan string, wg *sync.WaitGroup) {

	go card.TrapReceiver(conf.Snmp.TrapListener, conf.Snmp.Addr, conf.Snmp.TrapCommunity, logger, trapc, wg)

	for {

		select {
		case trap := <-trapc:
			if card.NeedPoweroff(trap) {
				// shutdown
			}

		case <-time.After(conf.Snmp.PollTime):
			// next polling
		}

		md, err := card.Monitor(conf.Snmp.Addr, conf.Snmp.Community)
		if err != nil {
			logger.Error("%s", err)
			continue
		}

		switch md.UpsBaseOutputStatus {
		case card.UpsBaseOutputStatusOnline:
			logger.Debug("UPS output run Online")
			continue

		case card.UpsBaseOutputStatusOnBattery:
			logger.Alert("UPS run on battery power, TimeOnBattery(%s), BatteryCapacity(%d%%), RunTimeRemaining(%s)",
				md.UpsBaseBatteryTimeOnBattery,
				md.UpsSmartBatteryCapacity,
				md.UpsSmartBatteryRunTimeRemaining)

		case card.UpsBaseOutputStatusBypass:
			logger.Alert("UPS run in Bypass mode")
			continue

		case card.UpsBaseOutputStatusSleep:
			logger.Alert("UPS enter Sleep mode")
			continue

		default:
			logger.Debug("UPS output status is %d", md.UpsBaseOutputStatus)
			continue
		}

		if md.UpsBaseBatteryStatus == card.UpsBaseBatteryStatusLow {
			logger.Emerg("Initiated system shutdown process, low battery level, BatteryCapacity(%d%%)",
				md.UpsSmartBatteryCapacity)
			// Shutdown
		}

		if md.UpsSmartBatteryCapacity < conf.BatteryLevel {
			logger.Emerg("Initiated system shutdown process, BatteryCapacity(%d%%) < BatteryLevel(%d%%)",
				md.UpsSmartBatteryCapacity, conf.BatteryLevel)
			// Shutdown
		}

		if md.UpsSmartBatteryRunTimeRemaining < conf.Minutes {
			logger.Emerg("Initiated system shutdown process, RunTimeRemaining(%s) < Minutes(%s)",
				md.UpsSmartBatteryRunTimeRemaining, conf.Minutes)
			// Shutdown
		}

		if conf.Timeout > 0 && md.UpsBaseBatteryTimeOnBattery > conf.Timeout {
			logger.Emerg("Initiated system shutdown process, TimeOnBattery(%s) > Timeout(%s)",
				md.UpsBaseBatteryTimeOnBattery, conf.Timeout)
			// Shutdown
		}
	}
}
