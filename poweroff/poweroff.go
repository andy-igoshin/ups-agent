// Для работы пакета необходим Capability CAP_SETUID

package poweroff

import (
	"os/exec"
	"sync"
	"syscall"
	"time"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

type PowerOff struct {
	once sync.Once
	log  *log.Logger
	cmd  *exec.Cmd
	host bool // выполнять ли выключение хоста
	ups  bool // выполнять ли выключение ИБП
}

func New(conf *config.Config, logger *log.Logger) *PowerOff {
	p := PowerOff{}
	p.cmd = exec.Command(conf.PowerCtl.PoweroffCmd)
	p.log = logger
	p.host = conf.PowerCtl.Poweroff
	p.ups = conf.PowerCtl.UpsPoweroff
	return &p
}

func (p *PowerOff) Poweroff(msg string) {
	p.once.Do(func() {
		if p.host {
			if p.ups {
				time.Sleep(2 * time.Second)
			}
			p.poweroff(msg)
		}
	})
}

func (p *PowerOff) poweroff(msg string) {

	p.log.Emerg(msg)

	syscall.Sync()
	syscall.Sync()
	syscall.Sync()

	err := syscall.Seteuid(0)
	if err != nil {
		p.log.Error("Poweroff Seteuid(0): %s", err)
		return
	}

	err = p.cmd.Run()
	if err != nil {
		p.log.Error("Poweroff Run(): %s", err)
	}
}
