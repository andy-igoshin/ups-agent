package card

import (
	"fmt"
	"net"
	"time"

	g "github.com/gosnmp/gosnmp"

	"git.vsu.ru/ai/ups-agent/log"
)

func TrapReceiver(listenAddr string, srcAddr string, trapCommunity string, logger *log.Logger, trapc chan string) {

	trapAcl := net.ParseIP(srcAddr)

	for {
		tl := g.NewTrapListener()
		tl.Params = g.Default
		tl.Params.Community = trapCommunity
		tl.OnNewTrap = func(packet *g.SnmpPacket, addr *net.UDPAddr) {
			trapHandler(packet, addr, trapAcl, logger, trapc)
		}

		logger.Info("Starting SNMP Trap Listener on %s", listenAddr)
		err := tl.Listen(listenAddr)
		if err != nil {
			logger.Error("Trap Listener failed: %s. Retrying in 10 seconds...", err)
		} else {
			logger.Warn("Trap Listener stopped. Retrying in 10 seconds...")
		}

		time.Sleep(10 * time.Second)
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
				continue
			case g.Null:
				val = ""
			default:
				if z := g.ToBigInt(v.Value); z != nil {
					val = fmt.Sprintf("%d", (*z).Int64())
				}
			}
			logger.Debug("- oid[%d]: %s (%s) = %v", i, v.Name, v.Type, val)
		}
	}

	return
}

func trapHandler(packet *g.SnmpPacket, addr *net.UDPAddr, trapAcl net.IP, logger *log.Logger, trapw chan<- string) {

	if !trapAcl.Equal(addr.IP) {
		return
	}

	oid, desc := getTrapInfo(packet, logger)

	if ts, ok := traps[oid]; ok {
		logger.Notice("Got trap '%s', %s", ts.name, desc)

		if ts.send {
			select {
			case trapw <- oid:
			default:
				logger.Warn("Trap queue is full, dropped trap '%s' (%s)", ts.name, oid)
			}
		}

	} else {
		logger.Error("Got unknown trap: %s", oid)
	}
}
