package card

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

var configOids = []string{
	UpsSmartConfigRatedOutputVoltage,
	UpsSmartConfigHighTransferVolt,
	UpsSmartConfigLowTransferVolt,
	UpsSmartConfigAlarm,
	UpsSmartConfigAlarmTimer,
	UpsSmartConfigMinReturnCapacity,
	UpsSmartConfigSensitivity,
	UpsSmartConfigLowBatteryRunTime,
	UpsSmartConfigReturnDelay,
	UpsSmartConfigShutoffDelay,
	UpsSmartConfigUpsSleepTime,
}

func GetConfig(addr, community string) (map[string]any, error) {

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

	res, err := snmp.Get(configOids)
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

func SetOids(addr, community string, data map[string]int) error {

	pdus := make([]g.SnmpPDU, 0, len(data))

	for k, v := range data {
		pdu := g.SnmpPDU{Name: k, Type: g.Integer, Value: v}
		pdus = append(pdus, pdu)
	}

	snmp := &g.GoSNMP{
		Target:    addr,
		Port:      g.Default.Port,
		Community: community,
		Version:   g.Default.Version,
		Timeout:   g.Default.Timeout,
	}

	err := snmp.Connect()
	if err != nil {
		return fmt.Errorf("snmp connect(): %s", err)
	}
	defer snmp.Conn.Close()

	_, err = snmp.Set(pdus)
	if err != nil {
		return fmt.Errorf("snmp set(): %s", err)
	}

	return nil
}
