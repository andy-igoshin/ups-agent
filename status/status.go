package status

import (
	"fmt"

	"git.vsu.ru/ai/ups-agent/card"
	"git.vsu.ru/ai/ups-agent/config"
)

func Print(conf *config.Config) {

	st, err := card.Status(conf.Snmp.Addr, conf.Snmp.Community)
	if err != nil {
		fmt.Println(err)
		return
	}

	form := "%-22s: %s\n"

	fmt.Printf(form, "Hostname", conf.Hostname)
	fmt.Printf(form, "UPS Name", getUpsName(st, conf.UpsName))
	fmt.Printf(form, "Model", getUpsModel(st, conf.Model))
	fmt.Printf(form, "Status", getOutputStatus(st))
	fmt.Printf(form, "Line Voltage", getLineVoltage(st))
	fmt.Printf(form, "Line Frequency", getLineFreq(st))
	fmt.Printf(form, "Line Fail Cause", getLineFailCause(st))
	fmt.Printf(form, "Output Voltage", getOutputVoltage(st))
	fmt.Printf(form, "Output Frequency", getOutputFrequency(st))
	fmt.Printf(form, "Output Load", getOutputLoad(st))
	fmt.Printf(form, "Battery Status", getBatteryStatus(st))
	fmt.Printf(form, "Battery Capacity", getBatteryCapacity(st))
	fmt.Printf(form, "RunTime Remaining", getRunTimeRemaining(st))
	fmt.Printf(form, "Min Battery Level", getBatteryLevel(conf.PowerCtl.BatteryLevel))
	fmt.Printf(form, "Min Time Left", conf.PowerCtl.Minutes.String())
	fmt.Printf(form, "Return Delay", getReturnDelay(st))
	fmt.Printf(form, "Shutoff Delay", getShutoffDelay(st))
	fmt.Printf(form, "Low Battery RunTime", getLowBatteryRunTime(st))

	if tv := getTransferVoltage(st); len(tv) > 0 {
		fmt.Printf(form, "Transfer Voltage", tv)
	}

	fmt.Printf(form, "Min Return Capacity", getMinReturnCapacity(st))
	fmt.Printf(form, "Battery Temperature", getBatteryTemperature(st))
	fmt.Printf(form, "Battery Voltage", getBatteryVoltage(st))
	fmt.Printf(form, "Sensitivity", getSensitivity(st))
	fmt.Printf(form, "Diagnostics Results", getDiagnosticsResults(st))
	fmt.Printf(form, "Diagnostic Schedule", getDiagnosticSchedule(st))
	fmt.Printf(form, "Last Diagnostics Date", getLastDiagnosticsDate(st))
	fmt.Printf(form, "Battery Last Replace", getBatteryLastReplaceDate(st))

	if bri := getBatteryReplaceIndicator(st); len(bri) > 0 {
		fmt.Printf(form, "Battery Replace Ind.", bri)
	}

	if et := getEnvTemperature(st); len(et) > 0 {
		fmt.Printf(form, "Env Temperature", et)
	}

	if eh := getEnvHumidity(st); len(eh) > 0 {
		fmt.Printf(form, "Env Humidity", eh)
	}

	fmt.Printf(form, "Firmware UPS", getFw(st))
	fmt.Printf(form, "Firmware SNMP Card", getFwCard(st))
}
