package card

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func getOids(addr, community string, oids []string) (map[string]any, error) {

	snmp := newSnmpClient(addr, community)

	err := snmp.Connect()
	if err != nil {
		return nil, fmt.Errorf("snmp connect(): %s", err)
	}
	defer snmp.Conn.Close()

	res, err := snmp.Get(oids)
	if err != nil {
		return nil, fmt.Errorf("snmp get(): %s", err)
	}

	st := make(map[string]any, len(res.Variables))

	for _, v := range res.Variables {

		if v.Value == nil {
			continue
		}

		switch v.Type {
		case g.OctetString:
			st[v.Name] = string(v.Value.([]byte))
		case g.Integer, g.Counter32, g.Gauge32, g.TimeTicks, g.Counter64:
			if z := g.ToBigInt(v.Value); z != nil {
				st[v.Name] = int(z.Int64())
			}
		}
	}

	return st, nil
}
