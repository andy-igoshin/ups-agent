package config

import (
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: 60 * time.Second,
				},
				PowerCtl: struct {
					Poweroff     bool          `default:"true"`
					PoweroffCmd  string        `default:"/sbin/poweroff"`
					BatteryLevel int           `default:"10"`
					Minutes      time.Duration `default:"10m"`
					Timeout      time.Duration `default:"0s"`
					UpsPoweroff  bool
				}{
					BatteryLevel: 10,
					Minutes:      10 * time.Minute,
					Timeout:      0,
				},
			},
			wantErr: false,
		},
		{
			name: "empty snmp addr",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "",
					PollTime: 60 * time.Second,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid snmp addr",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "invalid-ip",
					PollTime: 60 * time.Second,
				},
			},
			wantErr: true,
		},
		{
			name: "negative poll time",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: -1 * time.Second,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid battery level low",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: 60 * time.Second,
				},
				PowerCtl: struct {
					Poweroff     bool          `default:"true"`
					PoweroffCmd  string        `default:"/sbin/poweroff"`
					BatteryLevel int           `default:"10"`
					Minutes      time.Duration `default:"10m"`
					Timeout      time.Duration `default:"0s"`
					UpsPoweroff  bool
				}{
					BatteryLevel: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid battery level high",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: 60 * time.Second,
				},
				PowerCtl: struct {
					Poweroff     bool          `default:"true"`
					PoweroffCmd  string        `default:"/sbin/poweroff"`
					BatteryLevel int           `default:"10"`
					Minutes      time.Duration `default:"10m"`
					Timeout      time.Duration `default:"0s"`
					UpsPoweroff  bool
				}{
					BatteryLevel: 101,
				},
			},
			wantErr: true,
		},
		{
			name: "negative minutes",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: 60 * time.Second,
				},
				PowerCtl: struct {
					Poweroff     bool          `default:"true"`
					PoweroffCmd  string        `default:"/sbin/poweroff"`
					BatteryLevel int           `default:"10"`
					Minutes      time.Duration `default:"10m"`
					Timeout      time.Duration `default:"0s"`
					UpsPoweroff  bool
				}{
					BatteryLevel: 10,
					Minutes:      -1 * time.Minute,
				},
			},
			wantErr: true,
		},
		{
			name: "negative timeout",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:     "192.168.1.1",
					PollTime: 60 * time.Second,
				},
				PowerCtl: struct {
					Poweroff     bool          `default:"true"`
					PoweroffCmd  string        `default:"/sbin/poweroff"`
					BatteryLevel int           `default:"10"`
					Minutes      time.Duration `default:"10m"`
					Timeout      time.Duration `default:"0s"`
					UpsPoweroff  bool
				}{
					BatteryLevel: 10,
					Minutes:      10 * time.Minute,
					Timeout:      -1 * time.Second,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid trap listener",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:         "192.168.1.1",
					PollTime:     60 * time.Second,
					TrapListener: "invalid-listener-no-port",
				},
			},
			wantErr: true,
		},
		{
			name: "valid trap listener",
			config: Config{
				Snmp: struct {
					Addr          string
					Community     string        `default:"public"`
					PollTime      time.Duration `default:"60s"`
					TrapCommunity string        `default:"public"`
					TrapListener  string
				}{
					Addr:         "192.168.1.1",
					PollTime:     60 * time.Second,
					TrapListener: "0.0.0.0:162",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(&tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCliCardConfigType(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"get", CliCardConfigGet},
		{"set", CliCardConfigSet},
		{"default", CliCardConfigDefault},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := cliCardConfigType(tt.input)
			if got != tt.want {
				t.Errorf("cliCardConfigType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestConfigMethods(t *testing.T) {
	c := &Config{cliType: CliCardConfigGet}
	if !c.IsCardConfig() {
		t.Error("IsCardConfig() should be true for CliCardConfigGet")
	}
	if c.CardOp() != CliCardConfigGet {
		t.Errorf("CardOp() = %v, want %v", c.CardOp(), CliCardConfigGet)
	}
	if c.IsStatus() {
		t.Error("IsStatus() should be false for CliCardConfigGet")
	}

	c2 := &Config{cliType: CliStatus}
	if c2.IsCardConfig() {
		t.Error("IsCardConfig() should be false for CliStatus")
	}
	if c2.CardOp() != CliUnknown {
		t.Errorf("CardOp() = %v, want %v", c2.CardOp(), CliUnknown)
	}
	if !c2.IsStatus() {
		t.Error("IsStatus() should be true for CliStatus")
	}
}
