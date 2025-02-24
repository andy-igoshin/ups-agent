// Для работы пакета необходим Capability CAP_SETUID

package poweroff

import (
	"os/exec"
	"syscall"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

type PowerOff struct {
	cmd *exec.Cmd
	log *log.Logger
}

func New(conf *config.Config, logger *log.Logger) *PowerOff {
	p := PowerOff{}
	p.cmd = exec.Command(conf.PoweroffCmd)
	p.log = logger
	return &p
}

func (p *PowerOff) Poweroff() {

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
