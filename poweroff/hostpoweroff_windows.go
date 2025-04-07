//go:build windows

package poweroff

func (p *PowerOff) hostpoweroff(msg string) {

	p.logger.Emerg(msg)

	err := p.cmd.Run()
	if err != nil {
		p.logger.Error("Poweroff Run(): %s", err)
	}
}
