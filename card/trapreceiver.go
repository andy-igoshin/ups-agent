package card

import (
	"fmt"
	"net"
	"sync"

	g "github.com/gosnmp/gosnmp"

	"git.vsu.ru/ai/ups-agent/log"
)

func TrapReceiver(listenAddr string, srcAddr string, trapCommunity string, logger *log.Logger,
	trapc chan string, wg *sync.WaitGroup) {

	trapAcl := net.ParseIP(srcAddr)

	tl := g.NewTrapListener()
	tl.Params = g.Default
	tl.Params.Community = trapCommunity
	tl.OnNewTrap = func(packet *g.SnmpPacket, addr *net.UDPAddr) {
		trapHandler(packet, addr, trapAcl, logger, trapc, wg)
	}

	err := tl.Listen(listenAddr)
	if err != nil {
		logger.Fatalf("Listen failed: %s", err)
	}
}

func getTrapInfo(packet *g.SnmpPacket, logger *log.Logger) (oid, desc string) {

	for _, v := range packet.Variables {
		if v.Type == g.OctetString {
			oid = v.Name
			desc = string(v.Value.([]byte))
			break
		}
	}

	if len(oid) == 0 {
	ex:
		for i, v := range packet.Variables {
			var val string
			switch v.Type {
			case g.ObjectIdentifier:
				val = fmt.Sprintf("%s", v.Value)
				if ts, ok := traps[val]; ok {
					oid = val
					desc = ts.desc
					break ex
				}

			case g.TimeTicks:
				a := g.ToBigInt(v.Value)
				val = fmt.Sprintf("%d", (*a).Int64())
			case g.Null:
				val = ""
			default:
				a := g.ToBigInt(v.Value)
				val = fmt.Sprintf("%d", (*a).Int64())
			}
			logger.Debug("- oid[%d]: %s (%s) = %v \n", i, v.Name, v.Type, val)
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

	oid, desc := getTrapInfo(packet, logger)

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
