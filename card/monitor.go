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
	UpsSmartConfigAlarmTimer        time.Duration
}

var monitor = []string{
	UpsBaseBatteryStatus,
	UpsBaseBatteryTimeOnBattery,
	UpsSmartBatteryCapacity,
	UpsSmartBatteryRunTimeRemaining,
	UpsSmartInputLineFailCause,
	UpsBaseOutputStatus,
	UpsSmartConfigAlarmTimer,
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

	res, err := snmp.Get(monitor)
	if err != nil {
		return nil, fmt.Errorf("snmp get(): %s", err)
	}

	md := MonitorData{}

	for _, v := range res.Variables {

		switch v.Name {
		case UpsBaseBatteryStatus:
			md.UpsBaseBatteryStatus = v.Value.(int)

		case UpsBaseBatteryTimeOnBattery:
			md.UpsBaseBatteryTimeOnBattery = time.Duration(v.Value.(int)) * time.Second

		case UpsSmartBatteryCapacity:
			md.UpsSmartBatteryCapacity = v.Value.(int)

		case UpsSmartBatteryRunTimeRemaining:
			md.UpsSmartBatteryRunTimeRemaining = time.Duration(v.Value.(int)) * time.Second

		case UpsSmartInputLineFailCause:
			md.UpsSmartInputLineFailCause = v.Value.(int)

		case UpsBaseOutputStatus:
			md.UpsBaseOutputStatus = v.Value.(int)

		case UpsSmartConfigAlarmTimer:
			md.UpsSmartConfigAlarmTimer = time.Duration(v.Value.(int)) * time.Second

		}
	}

	return &md, nil
}
