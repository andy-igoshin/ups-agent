package card

import g "github.com/gosnmp/gosnmp"

func newSnmpClient(addr, community string) *g.GoSNMP {
	return &g.GoSNMP{
		Target:    addr,
		Port:      g.Default.Port,
		Community: community,
		Version:   g.Default.Version,
		Timeout:   g.Default.Timeout,
	}
}
