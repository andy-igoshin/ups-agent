package monitor

import (
	"fmt"
	"sync"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
	"git.vsu.ru/ai/ups-agent/poweroff"
)

func Run(conf *config.Config, logger *log.Logger, trapc chan string, wg *sync.WaitGroup) {

	go card.TrapReceiver(conf.Snmp.TrapListener, conf.Snmp.Addr, conf.Snmp.TrapCommunity, logger, trapc, wg)

	p := poweroff.New(conf, logger)

	for {
		select {
		case trap := <-trapc:
			if card.TrapClaimPoweroff(trap) {
				p.Poweroff(fmt.Sprintf("Initiated system shutdown process, trap(%s)", trap))
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
			logger.Debug("UPS output run Online, BatteryCapacity(%d%%)", md.UpsSmartBatteryCapacity)
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
			p.Poweroff(fmt.Sprintf("Initiated system shutdown process, low battery level, BatteryCapacity(%d%%)",
				md.UpsSmartBatteryCapacity))
		}

		if md.UpsSmartBatteryCapacity <= conf.PowerCtl.BatteryLevel {
			p.Poweroff(fmt.Sprintf("Initiated system shutdown process, BatteryCapacity(%d%%) <= BatteryLevel(%d%%)",
				md.UpsSmartBatteryCapacity, conf.PowerCtl.BatteryLevel))
		}

		if md.UpsSmartBatteryRunTimeRemaining <= conf.PowerCtl.Minutes {
			p.Poweroff(fmt.Sprintf("Initiated system shutdown process, RunTimeRemaining(%s) <= Minutes(%s)",
				md.UpsSmartBatteryRunTimeRemaining, conf.PowerCtl.Minutes))
		}

		if conf.PowerCtl.Timeout > 0 && md.UpsBaseBatteryTimeOnBattery >= conf.PowerCtl.Timeout {
			p.Poweroff(fmt.Sprintf("Initiated system shutdown process, TimeOnBattery(%s) >= Timeout(%s)",
				md.UpsBaseBatteryTimeOnBattery, conf.PowerCtl.Timeout))
		}
	}
}
