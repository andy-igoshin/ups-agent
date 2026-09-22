package card

import (
	"net"
	"os"
	"testing"

	g "github.com/gosnmp/gosnmp"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/log"
)

func TestNewSnmpClient(t *testing.T) {
	c := newSnmpClient("192.168.1.1", "public")

	if c.Target != "192.168.1.1" {
		t.Errorf("Target = %q, want %q", c.Target, "192.168.1.1")
	}
	if c.Community != "public" {
		t.Errorf("Community = %q, want %q", c.Community, "public")
	}
	if c.Port != g.Default.Port {
		t.Errorf("Port = %d, want %d", c.Port, g.Default.Port)
	}
	if c.Version != g.Default.Version {
		t.Errorf("Version = %v, want %v", c.Version, g.Default.Version)
	}
}

func TestTrapClaimPoweroff(t *testing.T) {
	tests := []struct {
		oid  string
		want bool
	}{
		{".1.3.6.1.4.1.935.0.12", true},   // upsTurnedOff
		{".1.3.6.1.4.1.935.0.52", true},   // upsScheduleShutdown
		{".1.3.6.1.4.1.935.0.54", true},   // upsShortCircuitShutdown
		{".1.3.6.1.4.1.935.0.55", true},   // upsInverterOutputFailShutdown
		{".1.3.6.1.4.1.935.0.57", true},   // upsHighDCShutdown
		{".1.3.6.1.4.1.935.0.58", true},   // upsEmergencyStop
		{".1.3.6.1.4.1.935.0.61", true},   // upsOverTemperatureShutdown
		{".1.3.6.1.4.1.935.0.62", true},   // upsOverLoadShutdown
		{".1.3.6.1.4.1.935.0.67", true},   // upsLowBatteryShutdown
		{".1.3.6.1.4.1.935.0.4", false},   // upsDischarged (send only)
		{".1.3.6.1.4.1.935.0.1", false},   // communicationLost
		{".1.3.6.1.6.3.1.1.5.1", false},   // coldStart
		{".1.3.6.1.4.1.935.0.999", false}, // unknown
	}

	for _, tt := range tests {
		if got := TrapClaimPoweroff(tt.oid); got != tt.want {
			t.Errorf("TrapClaimPoweroff(%q) = %v, want %v", tt.oid, got, tt.want)
		}
	}
}

func newTestLogger(t *testing.T) *log.Logger {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "ups-agent-card-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tmpDir) })

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "card.log"
	conf.Logger.Level = "debug"

	return log.New(conf)
}

func TestGetTrapInfo_OctetString(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	packet := &g.SnmpPacket{
		Variables: []g.SnmpPDU{
			{Name: ".1.3.6.1.4.1.935.0.4", Type: g.OctetString, Value: []byte("upsDischarged")},
		},
	}

	oid, desc := getTrapInfo(packet, logger)
	if oid != ".1.3.6.1.4.1.935.0.4" {
		t.Errorf("oid = %q, want %q", oid, ".1.3.6.1.4.1.935.0.4")
	}
	if desc != "upsDischarged" {
		t.Errorf("desc = %q, want %q", desc, "upsDischarged")
	}
}

func TestGetTrapInfo_ObjectIdentifier(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	packet := &g.SnmpPacket{
		Variables: []g.SnmpPDU{
			{Name: ".1.3.6.1.6.3.1.1.4.1.0", Type: g.ObjectIdentifier, Value: ".1.3.6.1.6.3.1.1.5.1"},
			{Name: ".1.3.6.1.6.3.1.1.4.3.0", Type: g.TimeTicks, Value: uint32(100)},
		},
	}

	oid, desc := getTrapInfo(packet, logger)
	if oid != ".1.3.6.1.6.3.1.1.5.1" {
		t.Errorf("oid = %q, want %q", oid, ".1.3.6.1.6.3.1.1.5.1")
	}
	if desc != "UPS Cold Start" {
		t.Errorf("desc = %q, want %q", desc, "UPS Cold Start")
	}
}

func TestGetTrapInfo_Empty(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	packet := &g.SnmpPacket{}
	oid, desc := getTrapInfo(packet, logger)
	if oid != "" {
		t.Errorf("oid = %q, want empty", oid)
	}
	if desc != "" {
		t.Errorf("desc = %q, want empty", desc)
	}
}

func TestTrapHandler_ACLMismatch(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	trapw := make(chan string, 1)
	trapAcl := net.ParseIP("192.168.1.1")
	addr := &net.UDPAddr{IP: net.ParseIP("192.168.1.2")}

	packet := &g.SnmpPacket{
		Variables: []g.SnmpPDU{
			{Name: ".1.3.6.1.4.1.935.0.4", Type: g.OctetString, Value: []byte("upsDischarged")},
		},
	}

	trapHandler(packet, addr, trapAcl, logger, trapw)

	select {
	case <-trapw:
		t.Error("expected no trap to be sent for mismatched ACL")
	default:
	}
}

func TestTrapHandler_Send(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	trapw := make(chan string, 1)
	trapAcl := net.ParseIP("192.168.1.1")
	addr := &net.UDPAddr{IP: net.ParseIP("192.168.1.1")}

	packet := &g.SnmpPacket{
		Variables: []g.SnmpPDU{
			{Name: ".1.3.6.1.4.1.935.0.4", Type: g.OctetString, Value: []byte("upsDischarged")},
		},
	}

	trapHandler(packet, addr, trapAcl, logger, trapw)

	select {
	case got := <-trapw:
		if got != ".1.3.6.1.4.1.935.0.4" {
			t.Errorf("received trap oid = %q, want %q", got, ".1.3.6.1.4.1.935.0.4")
		}
	default:
		t.Error("expected a trap to be sent")
	}
}

func TestTrapHandler_NotSend(t *testing.T) {
	logger := newTestLogger(t)
	defer logger.Close()

	trapw := make(chan string, 1)
	trapAcl := net.ParseIP("192.168.1.1")
	addr := &net.UDPAddr{IP: net.ParseIP("192.168.1.1")}

	// communicationLost has send=false
	packet := &g.SnmpPacket{
		Variables: []g.SnmpPDU{
			{Name: ".1.3.6.1.4.1.935.0.1", Type: g.OctetString, Value: []byte("communicationLost")},
		},
	}

	trapHandler(packet, addr, trapAcl, logger, trapw)

	select {
	case <-trapw:
		t.Error("expected no trap to be sent for non-send trap")
	default:
	}
}
