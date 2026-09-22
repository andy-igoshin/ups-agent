package poweroff

import (
	"errors"
	"os"
	"testing"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

func TestPowerOff_OnceAndCallbacks(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-poweroff-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "poweroff.log"
	conf.Logger.Level = "debug"
	conf.PowerCtl.Poweroff = true
	conf.PowerCtl.UpsPoweroff = true
	conf.PowerCtl.PoweroffCmd = "/bin/true" // harmless command
	conf.Snmp.Addr = "127.0.0.1"
	conf.Snmp.Community = "public"

	logger, _ := log.New(conf)
	defer logger.Close()

	upsPoweroffCalled := 0
	var gotAddr, gotCommunity string
	mockUpsPoweroff := func(addr, community string) error {
		upsPoweroffCalled++
		gotAddr = addr
		gotCommunity = community
		return nil
	}

	p := New(conf, logger, mockUpsPoweroff)

	// Call Poweroff first time
	p.Poweroff("test poweroff message")

	if upsPoweroffCalled != 1 {
		t.Errorf("expected upsPoweroff to be called 1 time, got %d", upsPoweroffCalled)
	}
	if gotAddr != "127.0.0.1" {
		t.Errorf("expected addr '127.0.0.1', got %q", gotAddr)
	}
	if gotCommunity != "public" {
		t.Errorf("expected community 'public', got %q", gotCommunity)
	}

	// Call Poweroff second time - sync.Once should prevent it from running again
	p.Poweroff("another message")

	if upsPoweroffCalled != 1 {
		t.Errorf("expected upsPoweroff to still be called 1 time (sync.Once), got %d", upsPoweroffCalled)
	}
}

func TestPowerOff_NoUps(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-poweroff-test-noups")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "poweroff.log"
	conf.Logger.Level = "debug"
	conf.PowerCtl.Poweroff = true
	conf.PowerCtl.UpsPoweroff = false
	conf.PowerCtl.PoweroffCmd = "/bin/true"

	logger, _ := log.New(conf)
	defer logger.Close()

	upsPoweroffCalled := 0
	mockUpsPoweroff := func(addr, community string) error {
		upsPoweroffCalled++
		return nil
	}

	p := New(conf, logger, mockUpsPoweroff)
	p.Poweroff("test poweroff message")

	if upsPoweroffCalled != 0 {
		t.Errorf("expected upsPoweroff NOT to be called, got %d", upsPoweroffCalled)
	}
}

func TestPowerOff_NoHost(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-poweroff-test-nohost")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "poweroff.log"
	conf.Logger.Level = "debug"
	conf.PowerCtl.Poweroff = false
	conf.PowerCtl.UpsPoweroff = true

	logger, _ := log.New(conf)
	defer logger.Close()

	upsPoweroffCalled := 0
	mockUpsPoweroff := func(addr, community string) error {
		upsPoweroffCalled++
		return nil
	}

	p := New(conf, logger, mockUpsPoweroff)
	p.Poweroff("test poweroff message")

	if upsPoweroffCalled != 0 {
		t.Errorf("expected upsPoweroff NOT to be called when host poweroff is disabled, got %d", upsPoweroffCalled)
	}
}

func TestPowerOff_UpsPoweroffError(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-poweroff-test-upserr")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "poweroff.log"
	conf.Logger.Level = "debug"
	conf.PowerCtl.Poweroff = true
	conf.PowerCtl.UpsPoweroff = true
	conf.PowerCtl.PoweroffCmd = "/bin/true"

	logger, _ := log.New(conf)
	defer logger.Close()

	mockUpsPoweroff := func(addr, community string) error {
		return errors.New("snmp error")
	}

	p := New(conf, logger, mockUpsPoweroff)
	// Should not panic or crash, just log the error
	p.Poweroff("test poweroff message")
}
