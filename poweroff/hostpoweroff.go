//go:build !windows

// Для работы пакета необходим Capability CAP_SETUID

package poweroff

import (
	"os/exec"
	"syscall"
)

func (p *PowerOff) hostpoweroff(msg string) {

	p.logger.Emerg("%s", msg)

	syscall.Sync()
	syscall.Sync()
	syscall.Sync()

	err := syscall.Seteuid(0)
	if err != nil {
		p.logger.Error("Poweroff Seteuid(0): %s", err)
		return
	}

	err = exec.Command(p.cmdName, p.cmdArgs...).Run()
	if err != nil {
		p.logger.Error("Poweroff Run(): %s", err)
	}
}
