package card

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

var status = []string{
	SysDescr,
	UpsOutputCurrent,
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

func Status(addr, community string) (map[string]any, error) {

	snmp := &g.GoSNMP{
		Target:    addr,
		Port:      g.Default.Port,
		Community: community,
		Version:   g.Default.Version,
		Timeout:   g.Default.Timeout,
	}

	err := snmp.Connect()
	if err != nil {
		return nil, fmt.Errorf("snmp connect(): %s", err)
	}
	defer snmp.Conn.Close()

	res, err := snmp.Get(status)
	if err != nil {
		return nil, fmt.Errorf("snmp get(): %s", err)
	}

	st := make(map[string]any, len(res.Variables))

	for _, v := range res.Variables {
		switch v.Type {
		case g.OctetString:
			st[v.Name] = string(v.Value.([]byte))
		default:
			st[v.Name] = v.Value.(int)
		}
	}

	return st, nil
}
