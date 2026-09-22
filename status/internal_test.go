package status

import (
	"testing"

	"git.vsu.ru/ai/ups-agent/card"
)

func TestGetLastDiagnosticsDate(t *testing.T) {
	st := map[string]any{
		card.UpsSmartTestLastDiagnosticsDate: "2025-06-15",
	}
	if got := getLastDiagnosticsDate(st); got != "2025-06-15" {
		t.Errorf("getLastDiagnosticsDate() = %q, want %q", got, "2025-06-15")
	}

	if got := getLastDiagnosticsDate(map[string]any{}); got != "" {
		t.Errorf("getLastDiagnosticsDate() empty = %q, want empty", got)
	}

	// Non-string value should be ignored
	stBad := map[string]any{card.UpsSmartTestLastDiagnosticsDate: 12345}
	if got := getLastDiagnosticsDate(stBad); got != "" {
		t.Errorf("getLastDiagnosticsDate() non-string = %q, want empty", got)
	}
}

func TestGetOutputCurrent(t *testing.T) {
	st := map[string]any{
		card.UpsOutputCurrent: 15, // 1.5 A
	}
	if got := getOutputCurrent(st); got != 1.5 {
		t.Errorf("getOutputCurrent() = %v, want %v", got, 1.5)
	}

	if got := getOutputCurrent(map[string]any{}); got != 0.0 {
		t.Errorf("getOutputCurrent() empty = %v, want 0.0", got)
	}
}

func TestGetOutputFrequency(t *testing.T) {
	st := map[string]any{
		card.UpsSmartOutputFrequency: 500, // 50.0 Hz
	}
	if got := getOutputFrequency(st); got != "50.0 Гц" {
		t.Errorf("getOutputFrequency() = %q, want %q", got, "50.0 Гц")
	}

	if got := getOutputFrequency(map[string]any{}); got != "" {
		t.Errorf("getOutputFrequency() empty = %q, want empty", got)
	}
}

func TestGetOutputVoltageVolt(t *testing.T) {
	st := map[string]any{
		card.UpsSmartOutputVoltage: 2200, // 220.0 V
	}
	if got := getOutputVoltageVolt(st); got != 220.0 {
		t.Errorf("getOutputVoltageVolt() = %v, want %v", got, 220.0)
	}

	if got := getOutputVoltageVolt(map[string]any{}); got != 0.0 {
		t.Errorf("getOutputVoltageVolt() empty = %v, want 0.0", got)
	}
}
