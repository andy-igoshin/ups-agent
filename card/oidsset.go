package card

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func SetOids(addr, community string, data map[string]int) error {

	pdus := make([]g.SnmpPDU, 0, len(data))

	for k, v := range data {
		pdu := g.SnmpPDU{Name: k, Type: g.Integer, Value: v}
		pdus = append(pdus, pdu)
	}

	snmp := newSnmpClient(addr, community)

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
