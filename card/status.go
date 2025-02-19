package card

import (
	"fmt"

	snmp "github.com/gosnmp/gosnmp"
)

var oidsStatus = []string{
	SysDescr,
	UpsBaseIdentModel,
	UpsBaseIdentUpsName,
	UpsSmartIdentFirmwareRevision,
	UpsSmartIdentAgentFirmwareRevision,
	UpsBaseBatteryStatus,
	UpsBaseBatteryLastReplaceDate,
	UpsSmartBatteryCapacity,
	UpsSmartBatteryVoltage,
	UpsSmartBatteryTemperature,
	UpsSmartBatteryRunTimeRemaining,
	UpsSmartBatteryReplaceIndicator,
	UpsSmartInputLineVoltage,
	UpsSmartInputMaxLineVoltage,
	UpsSmartInputMinLineVoltage,
	UpsSmartInputFrequency,
	UpsSmartInputLineFailCause,
	UpsBaseOutputStatus,
	UpsSmartOutputVoltage,
	UpsSmartOutputFrequency,
	UpsSmartOutputLoad,
	UpsSmartConfigHighTransferVolt,
	UpsSmartConfigLowTransferVolt,
	UpsSmartConfigAlarmTimer,
	UpsSmartConfigMinReturnCapacity,
	UpsSmartConfigSensitivity,
	UpsSmartConfigLowBatteryRunTime,
	UpsSmartConfigReturnDelay,
	UpsSmartConfigShutoffDelay,
	UpsSmartTestDiagnosticSchedule,
	UpsSmartTestDiagnosticsResults,
	UpsSmartTestLastDiagnosticsDate,
	UpsEnvTemperature,
	UpsEnvHumidity,
	UpsEnvOverTemperature,
	UpsEnvUnderTemperature,
	UpsEnvOverHumidity,
	UpsEnvUnderHumidity,
}

func GetStatus(addr, community string) (map[string]any, error) {

	snmp.Default.Community = community
	snmp.Default.Target = addr

	err := snmp.Default.Connect()
	if err != nil {
		return nil, fmt.Errorf("snmp connect(): %s", err)
	}
	defer snmp.Default.Conn.Close()

	res, err := snmp.Default.Get(oidsStatus)
	if err != nil {
		return nil, fmt.Errorf("snmp get(): %s", err)
	}

	st := make(map[string]any, len(res.Variables))

	for _, v := range res.Variables {
		switch v.Type {
		case snmp.OctetString:
			st[v.Name] = string(v.Value.([]byte))
		default:
			st[v.Name] = v.Value.(int)
		}
	}

	return st, nil
}
