package card

import (
	"fmt"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type MonitorData struct {
	UpsBaseBatteryStatus            int
	UpsBaseBatteryTimeOnBattery     time.Duration
	UpsSmartBatteryCapacity         int
	UpsSmartBatteryRunTimeRemaining time.Duration
	UpsSmartInputLineFailCause      int
	UpsBaseOutputStatus             int
}

var monitorOids = []string{
	UpsBaseBatteryStatus,
	UpsBaseBatteryTimeOnBattery,
	UpsSmartBatteryCapacity,
	UpsSmartBatteryRunTimeRemaining,
	UpsSmartInputLineFailCause,
	UpsBaseOutputStatus,
}

func Monitor(addr, community string) (*MonitorData, error) {

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

	res, err := snmp.Get(monitorOids)
	if err != nil {
		return nil, fmt.Errorf("snmp get(): %s", err)
	}

	md := MonitorData{}

	for _, v := range res.Variables {

		if v.Value == nil {
			return nil, fmt.Errorf("OID %s is nil", v.Name)
		}

		val, ok := v.Value.(int)
		if !ok {
			return nil, fmt.Errorf("OID %s: expected int, got %T", v.Name, v.Value)
		}

		switch v.Name {
		case UpsBaseBatteryStatus:
			md.UpsBaseBatteryStatus = val

		case UpsBaseBatteryTimeOnBattery:
			md.UpsBaseBatteryTimeOnBattery = time.Duration(val) * time.Second

		case UpsSmartBatteryCapacity:
			md.UpsSmartBatteryCapacity = val

		case UpsSmartBatteryRunTimeRemaining:
			md.UpsSmartBatteryRunTimeRemaining = time.Duration(val) * time.Second

		case UpsSmartInputLineFailCause:
			md.UpsSmartInputLineFailCause = val

		case UpsBaseOutputStatus:
			md.UpsBaseOutputStatus = val
		}
	}

	return &md, nil
}
