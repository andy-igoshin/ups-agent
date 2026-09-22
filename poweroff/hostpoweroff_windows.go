//go:build windows

package poweroff

import (
	"os/exec"
)

func (p *PowerOff) hostpoweroff(msg string) {

	p.logger.Emerg("%s", msg)

	err := exec.Command(p.cmdName, p.cmdArgs...).Run()
	if err != nil {
		p.logger.Error("Poweroff Run(): %s", err)
	}
}
