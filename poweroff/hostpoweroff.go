//go:build !windows
// Для работы пакета необходим Capability CAP_SETUID

package poweroff

import (
	"syscall"
)

func (p *PowerOff) hostpoweroff(msg string) {

	p.logger.Emerg(msg)

	syscall.Sync()
	syscall.Sync()
	syscall.Sync()

	err := syscall.Seteuid(0)
	if err != nil {
		p.logger.Error("Poweroff Seteuid(0): %s", err)
		return
	}

	err = p.cmd.Run()
	if err != nil {
		p.logger.Error("Poweroff Run(): %s", err)
	}
}
