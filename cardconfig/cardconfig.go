package cardconfig

import (
	"fmt"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
)

func Run(conf *config.Config) {

	switch conf.CardOp() {
	case config.CliCardConfigGet:
		printConfig(conf)

	case config.CliCardConfigSet:
		writeConfig(conf)

	case config.CliCardConfigDefault:
		setDefaultConfig(conf)
	}
}

func printConfig(conf *config.Config) {

	st, err := card.GetConfig(conf.Snmp.Addr, conf.Snmp.Community)
	if err != nil {
		fmt.Println(err)
		return
	}

	form := "%-19s: %s\n"

	fmt.Printf(form, "RatedOutputVoltage", getRatedOutputVoltage(st))
	fmt.Printf(form, "HighTransferVolt", getHighTransferVolt(st))
	fmt.Printf(form, "LowTransferVolt", getLowTransferVolt(st))
	fmt.Printf(form, "Alarm", getAlarm(st))
	fmt.Printf(form, "AlarmTimer", getAlarmTimer(st))
	fmt.Printf(form, "MinReturnCapacity", getMinReturnCapacity(st))
	fmt.Printf(form, "Sensitivity", getSensitivity(st))
	fmt.Printf(form, "LowBatteryRunTime", getLowBatteryRunTime(st))
	fmt.Printf(form, "ReturnDelay", getReturnDelay(st))
	fmt.Printf(form, "ShutoffDelay", getShutoffDelay(st))
	fmt.Printf(form, "UpsSleepTime", getUpsSleepTime(st))
}

func setDefaultConfig(conf *config.Config) {

	data := make(map[string]int, 8)

	data[card.UpsSmartConfigRatedOutputVoltage] = 2300 // 230 В
	data[card.UpsSmartConfigAlarm] = 3                 // Никогда (3)
	data[card.UpsSmartConfigMinReturnCapacity] = 0     // 0%
	data[card.UpsSmartConfigSensitivity] = 1           // Авто (1)
	data[card.UpsSmartConfigLowBatteryRunTime] = 120   // секунд
	data[card.UpsSmartConfigReturnDelay] = 60          // секунд
	data[card.UpsSmartConfigShutoffDelay] = 120        // секунд
	data[card.UpsSmartConfigUpsSleepTime] = 2          // минут

	err := card.SetOids(conf.Snmp.Addr, conf.Snmp.Community, data)
	if err != nil {
		fmt.Println(err)
	}
}

func writeConfig(conf *config.Config) {

	c := conf.CardConfig
	data := make(map[string]int, 11)

	// UpsSmartConfigRatedOutputVoltage
	switch c.RatedOutputVoltage {
	case 220, 225, 230, 240:
		data[card.UpsSmartConfigRatedOutputVoltage] = c.RatedOutputVoltage * 10
	case -1:
	default:
		fmt.Println("cardconfig: invalid RatedOutputVoltage:", c.RatedOutputVoltage)
		return
	}

	// UpsSmartConfigHighTransferVolt
	switch c.HighTransferVolt {
	case 253, 264, 271, 280:
		data[card.UpsSmartConfigHighTransferVolt] = c.HighTransferVolt * 10
	case -1:
	default:
		fmt.Println("cardconfig: invalid HighTransferVolt:", c.HighTransferVolt)
		return
	}

	// UpsSmartConfigLowTransferVolt
	switch c.LowTransferVolt {
	case 188, 196, 204, 208:
		data[card.UpsSmartConfigLowTransferVolt] = c.LowTransferVolt * 10
	case -1:
	default:
		fmt.Println("cardconfig: invalid LowTransferVolt:", c.LowTransferVolt)
		return
	}

	// UpsSmartConfigAlarm
	switch c.Alarm {
	case 1, 2, 3:
		data[card.UpsSmartConfigAlarm] = c.Alarm
	case -1:
	default:
		fmt.Println("cardconfig: invalid Alarm:", c.Alarm)
		return
	}

	// UpsSmartConfigAlarmTimer
	switch int(c.AlarmTimer / time.Second) {
	case 0, 30:
		data[card.UpsSmartConfigAlarmTimer] = int(c.AlarmTimer / time.Second)
	case -1:
	default:
		fmt.Println("cardconfig: invalid AlarmTimer:", c.AlarmTimer)
		return
	}

	// UpsSmartConfigMinReturnCapacity
	switch c.MinReturnCapacity {
	case 0, 10, 25, 90:
		data[card.UpsSmartConfigMinReturnCapacity] = c.MinReturnCapacity
	case -1:
	default:
		fmt.Println("cardconfig: invalid MinReturnCapacity:", c.MinReturnCapacity)
		return
	}

	// UpsSmartConfigSensitivity
	switch c.Sensitivity {
	case 1, 2, 3, 4:
		data[card.UpsSmartConfigSensitivity] = c.Sensitivity
	case -1:
	default:
		fmt.Println("cardconfig: invalid Sensitivity:", c.Sensitivity)
		return
	}

	// UpsSmartConfigLowBatteryRunTime
	switch int(c.LowBatteryRunTime / time.Second) {
	case 120, 300, 420, 600:
		data[card.UpsSmartConfigLowBatteryRunTime] = int(c.LowBatteryRunTime / time.Second)
	case -1:
	default:
		fmt.Println("cardconfig: invalid LowBatteryRunTime:", c.LowBatteryRunTime)
		return
	}

	// UpsSmartConfigReturnDelay
	switch int(c.ReturnDelay / time.Second) {
	case 0, 60, 180, 300:
		data[card.UpsSmartConfigReturnDelay] = int(c.ReturnDelay / time.Second)
	case -1:
	default:
		fmt.Println("cardconfig: invalid ReturnDelay:", c.ReturnDelay)
		return
	}

	// UpsSmartConfigShutoffDelay
	switch int(c.ShutoffDelay / time.Second) {
	case 12, 18, 24, 30, 36, 42, 48, 54, 60, 120, 180, 240, 300, 360, 420, 480, 540, 600:
		data[card.UpsSmartConfigShutoffDelay] = int(c.ShutoffDelay / time.Second)
	case -1:
	default:
		fmt.Println("cardconfig: invalid ShutoffDelay:", c.ShutoffDelay)
		return
	}

	// UpsSmartConfigUpsSleepTime
	switch int(c.UpsSleepTime / time.Minute) {
	case -1:
	default:
		data[card.UpsSmartConfigUpsSleepTime] = int(c.UpsSleepTime / time.Minute)
	}

	err := card.SetOids(conf.Snmp.Addr, conf.Snmp.Community, data)
	if err != nil {
		fmt.Println(err)
	}
}
