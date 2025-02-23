package card

import (
	"net"
	"sync"

	g "github.com/gosnmp/gosnmp"

	"git.vsu.ru/ai/ups-agent/log"
)

func TrapReceiver(listenAddr string, srcAddr string, logger *log.Logger, trapc chan string, wg *sync.WaitGroup) {

	trapAcl := net.ParseIP(srcAddr)

	tl := g.NewTrapListener()
	tl.Params = g.Default
	tl.OnNewTrap = func(packet *g.SnmpPacket, addr *net.UDPAddr) {
		trapHandler(packet, addr, trapAcl, logger, trapc, wg)
	}

	err := tl.Listen(listenAddr)
	if err != nil {
		logger.Fatalf("Listen failed: %s", err)
	}
}

func getTrapInfo(packet *g.SnmpPacket) (oid, desc string) {
	for _, v := range packet.Variables {
		if v.Type == g.OctetString {
			oid = v.Name
			desc = string(v.Value.([]byte))
			break
		}
	}
	return
}

func trapHandler(packet *g.SnmpPacket, addr *net.UDPAddr, trapAcl net.IP, logger *log.Logger,
	trapw chan<- string, wg *sync.WaitGroup) {

	if !trapAcl.Equal(addr.IP) {
		return
	}

	wg.Add(1)

	oid, desc := getTrapInfo(packet)

	if ts, ok := traps[oid]; ok {
		logger.Notice("Got trap '%s', %s", ts.name, desc)

		if ts.send {
			trapw <- oid
		}

	} else {
		logger.Error("Got unknown trap: %s", oid)
	}

	wg.Done()
}
