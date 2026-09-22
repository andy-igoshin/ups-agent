package cardconfig

import (
	"testing"

	"git.vsu.ru/ai/ups-agent/card"
)

func TestGetAlarm(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "По времени (1)"},
		{2, "При низком заряде АКБ (2)"},
		{3, "Никогда (3)"},
		{4, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartConfigAlarm: tt.val}
		if got := getAlarm(st); got != tt.want {
			t.Errorf("getAlarm(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetAlarmTimer(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigAlarmTimer: 30}
	if got := getAlarmTimer(st); got != "30s" {
		t.Errorf("getAlarmTimer() = %q, want %q", got, "30s")
	}

	if got := getAlarmTimer(map[string]any{}); got != "" {
		t.Errorf("getAlarmTimer() empty = %q, want empty", got)
	}
}

func TestGetHighTransferVolt(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigHighTransferVolt: 2640}
	if got := getHighTransferVolt(st); got != "264.0 В" {
		t.Errorf("getHighTransferVolt() = %q, want %q", got, "264.0 В")
	}
}

func TestGetLowTransferVolt(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigLowTransferVolt: 1880}
	if got := getLowTransferVolt(st); got != "188.0 В" {
		t.Errorf("getLowTransferVolt() = %q, want %q", got, "188.0 В")
	}
}

func TestGetLowBatteryRunTime(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigLowBatteryRunTime: 120}
	if got := getLowBatteryRunTime(st); got != "2m0s" {
		t.Errorf("getLowBatteryRunTime() = %q, want %q", got, "2m0s")
	}
}

func TestGetMinReturnCapacity(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigMinReturnCapacity: 25}
	if got := getMinReturnCapacity(st); got != "25%" {
		t.Errorf("getMinReturnCapacity() = %q, want %q", got, "25%")
	}
}

func TestGetRatedOutputVoltage(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigRatedOutputVoltage: 2300}
	if got := getRatedOutputVoltage(st); got != "230.0 В" {
		t.Errorf("getRatedOutputVoltage() = %q, want %q", got, "230.0 В")
	}
}

func TestGetReturnDelay(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigReturnDelay: 60}
	if got := getReturnDelay(st); got != "1m0s" {
		t.Errorf("getReturnDelay() = %q, want %q", got, "1m0s")
	}
}

func TestGetSensitivity(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Авто (1)"},
		{2, "Низкая (2)"},
		{3, "Средняя (3)"},
		{4, "Высокая (4)"},
		{5, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartConfigSensitivity: tt.val}
		if got := getSensitivity(st); got != tt.want {
			t.Errorf("getSensitivity(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetShutoffDelay(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigShutoffDelay: 180}
	if got := getShutoffDelay(st); got != "3m0s" {
		t.Errorf("getShutoffDelay() = %q, want %q", got, "3m0s")
	}
}

func TestGetUpsSleepTime(t *testing.T) {
	st := map[string]any{card.UpsSmartConfigUpsSleepTime: 2}
	if got := getUpsSleepTime(st); got != "2m0s" {
		t.Errorf("getUpsSleepTime() = %q, want %q", got, "2m0s")
	}
}
