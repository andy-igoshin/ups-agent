package status

import (
	"testing"

	"git.vsu.ru/ai/ups-agent/card"
)

func TestGetBatteryCapacity(t *testing.T) {
	st := map[string]any{
		card.UpsSmartBatteryCapacity: 85,
	}
	got := getBatteryCapacity(st)
	want := "85%"
	if got != want {
		t.Errorf("getBatteryCapacity() = %q, want %q", got, want)
	}

	gotEmpty := getBatteryCapacity(map[string]any{})
	if gotEmpty != "" {
		t.Errorf("getBatteryCapacity() with empty map = %q, want empty string", gotEmpty)
	}
}

func TestGetBatteryLastReplaceDate(t *testing.T) {
	st := map[string]any{
		card.UpsBaseBatteryLastReplaceDate: "2025-01-01",
	}
	got := getBatteryLastReplaceDate(st)
	want := "2025-01-01"
	if got != want {
		t.Errorf("getBatteryLastReplaceDate() = %q, want %q", got, want)
	}
}

func TestGetBatteryLevel(t *testing.T) {
	got := getBatteryLevel(15)
	want := "15%"
	if got != want {
		t.Errorf("getBatteryLevel() = %q, want %q", got, want)
	}
}

func TestGetBatteryReplaceIndicator(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Замена не требуется (1)"},
		{2, "Замена требуется (2)"},
		{3, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartBatteryReplaceIndicator: tt.val}
		got := getBatteryReplaceIndicator(st)
		if got != tt.want {
			t.Errorf("getBatteryReplaceIndicator(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetBatteryStatus(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Неизвестно (1)"},
		{2, "АКБ в норме (2)"},
		{3, "Низкий заряд АКБ (3)"},
		{4, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsBaseBatteryStatus: tt.val}
		got := getBatteryStatus(st)
		if got != tt.want {
			t.Errorf("getBatteryStatus(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetBatteryTemperature(t *testing.T) {
	st := map[string]any{
		card.UpsSmartBatteryTemperature: 256, // 25.6 C
	}
	got := getBatteryTemperature(st)
	want := "25.6 ℃ "
	if got != want {
		t.Errorf("getBatteryTemperature() = %q, want %q", got, want)
	}
}

func TestGetBatteryVoltage(t *testing.T) {
	st := map[string]any{
		card.UpsSmartBatteryVoltage: 136, // 13.6 V
	}
	got := getBatteryVoltage(st)
	want := "13.6 В"
	if got != want {
		t.Errorf("getBatteryVoltage() = %q, want %q", got, want)
	}
}

func TestGetDiagnosticsResults(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Ок (1)"},
		{2, "Ошибка (2)"},
		{3, "Недействительный результат (3)"},
		{4, "Тест выполняется (4)"},
		{5, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartTestDiagnosticsResults: tt.val}
		got := getDiagnosticsResults(st)
		if got != tt.want {
			t.Errorf("getDiagnosticsResults(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetDiagnosticSchedule(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Неизвестно (1)"},
		{2, "Раз в 2 недели (2)"},
		{3, "Еженедельно (3)"},
		{4, "Никогда (4)"},
		{5, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartTestDiagnosticSchedule: tt.val}
		got := getDiagnosticSchedule(st)
		if got != tt.want {
			t.Errorf("getDiagnosticSchedule(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetEnvHumidity(t *testing.T) {
	st := map[string]any{
		card.UpsEnvHumidity:      45,
		card.UpsEnvUnderHumidity: 20,
		card.UpsEnvOverHumidity:  80,
	}
	got := getEnvHumidity(st)
	want := "45% (20/80)"
	if got != want {
		t.Errorf("getEnvHumidity() = %q, want %q", got, want)
	}
}

func TestGetEnvTemperature(t *testing.T) {
	st := map[string]any{
		card.UpsEnvTemperature:      225, // 22.5 C
		card.UpsEnvUnderTemperature: 150, // 15.0 C
		card.UpsEnvOverTemperature:  350, // 35.0 C
	}
	got := getEnvTemperature(st)
	want := "22.5 ℃  (15.0/35.0 ℃ )"
	if got != want {
		t.Errorf("getEnvTemperature() = %q, want %q", got, want)
	}
}

func TestGetFwAndFwCard(t *testing.T) {
	st := map[string]any{
		card.UpsSmartIdentFirmwareRevision:      "  v1.2.3  ",
		card.UpsSmartIdentAgentFirmwareRevision: "  v4.5.6  ",
	}
	if got := getFw(st); got != "v1.2.3" {
		t.Errorf("getFw() = %q, want 'v1.2.3'", got)
	}
	if got := getFwCard(st); got != "v4.5.6" {
		t.Errorf("getFwCard() = %q, want 'v4.5.6'", got)
	}
}

func TestGetLineFailCause(t *testing.T) {
	tests := []struct {
		val  int
		want string
	}{
		{1, "Причина отсутствует (1)"},
		{2, "Высокое напряжение (2)"},
		{3, "Низкое напряжение (3)"},
		{4, "Отключение питания (4)"},
		{5, "Кратковременное снижение напряжения (5)"},
		{6, "Длительное снижение напряжения (6)"},
		{7, "Кратковременный малый импульс (7)"},
		{8, "Кратковременный большой импульс (8)"},
		{9, ""},
	}

	for _, tt := range tests {
		st := map[string]any{card.UpsSmartInputLineFailCause: tt.val}
		got := getLineFailCause(st)
		if got != tt.want {
			t.Errorf("getLineFailCause(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetLineFreqAndVoltage(t *testing.T) {
	st := map[string]any{
		card.UpsSmartInputFrequency:      501, // 50.1 Hz
		card.UpsSmartInputLineVoltage:    2205,
		card.UpsSmartInputMinLineVoltage: 2180,
		card.UpsSmartInputMaxLineVoltage: 2250,
	}

	if got := getLineFreq(st); got != "50.1 Гц" {
		t.Errorf("getLineFreq() = %q, want '50.1 Гц'", got)
	}

	if got := getLineVoltage(st); got != "220.5 В (218.0/225.0 В)" {
		t.Errorf("getLineVoltage() = %q, want '220.5 В (218.0/225.0 В)'", got)
	}
}

func TestGetLowBatteryRunTimeAndMinReturnCapacity(t *testing.T) {
	st := map[string]any{
		card.UpsSmartConfigLowBatteryRunTime: 120, // 120 seconds
		card.UpsSmartConfigMinReturnCapacity: 10,
	}

	if got := getLowBatteryRunTime(st); got != "2m0s" {
		t.Errorf("getLowBatteryRunTime() = %q, want '2m0s'", got)
	}

	if got := getMinReturnCapacity(st); got != "10%" {
		t.Errorf("getMinReturnCapacity() = %q, want '10%%'", got)
	}
}

func TestGetOutputStatusAndLoad(t *testing.T) {
	st := map[string]any{
		card.UpsBaseOutputStatus:   2, // Online
		card.UpsSmartOutputLoad:    35,
		card.UpsOutputCurrent:      15,   // 1.5 A
		card.UpsSmartOutputVoltage: 2200, // 220.0 V
	}

	if got := getOutputStatus(st); got != "Онлайн (2)" {
		t.Errorf("getOutputStatus() = %q, want 'Онлайн (2)'", got)
	}

	if got := getOutputVoltage(st); got != "220.0 В" {
		t.Errorf("getOutputVoltage() = %q, want '220.0 В'", got)
	}

	// 1.5 A * 220.0 V = 330.0 VA
	if got := getOutputLoad(st); got != "35% (1.5 А, 330.0 ВА)" {
		t.Errorf("getOutputLoad() = %q, want '35%% (1.5 А, 330.0 ВА)'", got)
	}
}

func TestGetDelaysAndRunTime(t *testing.T) {
	st := map[string]any{
		card.UpsSmartConfigReturnDelay:       60,
		card.UpsSmartConfigShutoffDelay:      180,
		card.UpsSmartBatteryRunTimeRemaining: 3600,
	}

	if got := getReturnDelay(st); got != "1m0s" {
		t.Errorf("getReturnDelay() = %q, want '1m0s'", got)
	}

	if got := getShutoffDelay(st); got != "3m0s" {
		t.Errorf("getShutoffDelay() = %q, want '3m0s'", got)
	}

	if got := getRunTimeRemaining(st); got != "1h0m0s" {
		t.Errorf("getRunTimeRemaining() = %q, want '1h0m0s'", got)
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
		got := getSensitivity(st)
		if got != tt.want {
			t.Errorf("getSensitivity(%d) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestGetTransferVoltage(t *testing.T) {
	st := map[string]any{
		card.UpsSmartConfigLowTransferVolt:  1800,
		card.UpsSmartConfigHighTransferVolt: 2400,
	}
	got := getTransferVoltage(st)
	want := "180.0/240.0 В"
	if got != want {
		t.Errorf("getTransferVoltage() = %q, want %q", got, want)
	}
}

func TestGetUpsModelAndName(t *testing.T) {
	st := map[string]any{
		card.UpsBaseIdentModel:   "  Model-X  ",
		card.UpsBaseIdentUpsName: "  UPS-Y  ",
		card.SysDescr:            "  System Description  ",
	}

	// If config overrides, use config
	if got := getUpsModel(st, "ConfigModel"); got != "ConfigModel" {
		t.Errorf("getUpsModel with override = %q, want 'ConfigModel'", got)
	}
	if got := getUpsModel(st, ""); got != "Model-X" {
		t.Errorf("getUpsModel from status = %q, want 'Model-X'", got)
	}

	if got := getUpsName(st, "ConfigName"); got != "ConfigName" {
		t.Errorf("getUpsName with override = %q, want 'ConfigName'", got)
	}
	if got := getUpsName(st, ""); got != "UPS-Y" {
		t.Errorf("getUpsName from status = %q, want 'UPS-Y'", got)
	}

	// Test fallback to SysDescr
	stNoName := map[string]any{
		card.UpsBaseIdentUpsName: "",
		card.SysDescr:            "SysDescrFallback",
	}
	if got := getUpsName(stNoName, ""); got != "SysDescrFallback" {
		t.Errorf("getUpsName fallback = %q, want 'SysDescrFallback'", got)
	}
}
