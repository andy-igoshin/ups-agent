// Для работы пакета необходим Capability CAP_SETUID

package poweroff

import (
	"strings"
	"sync"
	"time"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

type PowerOff struct {
	once        sync.Once
	logger      *log.Logger
	cmdName     string
	cmdArgs     []string
	host        bool // выполнять ли выключение хоста
	ups         bool // выполнять ли выключение ИБП
	upspoweroff func(string, string) error
	addr        string
	community   string
}

func New(conf *config.Config, logger *log.Logger, upspoweroff func(string, string) error) *PowerOff {
	p := PowerOff{}
	p.logger = logger

	fields := strings.Fields(conf.PowerCtl.PoweroffCmd)
	if len(fields) > 0 {
		p.cmdName = fields[0]
	}
	if len(fields) > 1 {
		p.cmdArgs = fields[1:]
	}

	p.host = conf.PowerCtl.Poweroff
	p.ups = conf.PowerCtl.UpsPoweroff
	p.upspoweroff = upspoweroff
	p.addr = conf.Snmp.Addr
	p.community = conf.Snmp.Community
	return &p
}

func (p *PowerOff) Poweroff(msg string) {
	p.once.Do(func() {
		if p.host {
			if p.ups {
				p.logger.Notice("Sending 'upsSmartControlUpsOff' to UPS")
				err := p.upspoweroff(p.addr, p.community)
				if err != nil {
					p.logger.Error("%s", err)
				}

				// Wait for SNMP packet to be delivered before host shutdown
				const hostPoweroffDelay = 5 * time.Second
				time.Sleep(hostPoweroffDelay)
			}
			p.hostpoweroff(msg)
		}
	})
}
