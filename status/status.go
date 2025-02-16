package status

import (
	"fmt"
	"time"

	"git.vsu.ru/ai/ups-agent/card"
)

func Print(hostname, upsname, model, addr, community string, batterylevel int, minutes time.Duration) {

	st, err := card.GetStatus(addr, community)
	if err != nil {
		fmt.Println(err)
		return
	}

	form := "%-20s: %s\n"

	fmt.Printf(form, "Hostname", hostname)
	fmt.Printf(form, "UPS Name", getUpsName(st, upsname))
	fmt.Printf(form, "Model", getUpsModel(st, model))
	fmt.Printf(form, "Status", getStatus(st))
	fmt.Printf(form, "Line Voltage", getLineVoltage(st))
	fmt.Printf(form, "Line Frequency", getLineFreq(st))
	fmt.Printf(form, "Line Fail Cause", getLineFailCause(st))
	fmt.Printf(form, "Output Voltage", getOutputVoltage(st))
	fmt.Printf(form, "Output Load", getOutputLoad(st))
	fmt.Printf(form, "Battery Capacity", getCapacity(st))
	fmt.Printf(form, "RunTime Remaining", getRunTimeRemaining(st))
	fmt.Printf(form, "Min Battery Level", getBatteryLevel(batterylevel))
	fmt.Printf(form, "Min Time Left", minutes.String())
	fmt.Printf(form, "Return Delay", getReturnDelay(st))
	fmt.Printf(form, "Shutoff Delay", getShutoffDelay(st))
	fmt.Printf(form, "Low Battery RunTime", getLowBatteryRunTime(st))
	fmt.Printf(form, "Low Transfer Volt", getLowTransferVolt(st))
	fmt.Printf(form, "High Transfer Volt", getHighTransferVolt(st))
	fmt.Printf(form, "Min Return Capacity", getMinReturnCapacity(st))
	fmt.Printf(form, "Battery Temperature", getBatteryTemperature(st))
	fmt.Printf(form, "Battery Voltage", getBatteryVoltage(st))
	fmt.Printf(form, "Alarm Timer", getAlarmTimer(st))
	fmt.Printf(form, "Sensitivity", getSensitivity(st))

	fmt.Printf(form, "Firmware UPS", getFw(st))
	fmt.Printf(form, "Firmware SNMP Card", getFwCard(st))


}
