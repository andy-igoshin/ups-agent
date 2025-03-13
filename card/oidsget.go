package card

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func getOids(addr, community string, oids []string) (map[string]any, error) {

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

	res, err := snmp.Get(oids)
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
