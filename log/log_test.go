package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"git.vsu.ru/ai/ups-agent/config"
)

func TestGetLevel(t *testing.T) {
	l := &Logger{}

	tests := []struct {
		input string
		want  int
	}{
		{"emerg", log_EMERG},
		{"alert", log_ALERT},
		{"crit", log_CRIT},
		{"error", log_ERR},
		{"warn", log_WARNING},
		{"notice", log_NOTICE},
		{"info", log_INFO},
		{"debug", log_DEBUG},
		{"unknown", log_INFO}, // default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := l.getLevel(tt.input)
			if got != tt.want {
				t.Errorf("getLevel(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetLevelString(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{log_EMERG, "emerg"},
		{log_ALERT, "alert"},
		{log_CRIT, "crit"},
		{log_ERR, "error"},
		{log_WARNING, "warn"},
		{log_NOTICE, "notice"},
		{log_INFO, "info"},
		{log_DEBUG, "debug"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := getLevelString(tt.input)
			if got != tt.want {
				t.Errorf("getLevelString(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-log-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "test.log"
	conf.Logger.Level = "debug"

	logger, _ := New(conf)
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	defer logger.Close()

	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.Crit("crit message")
	logger.Alert("alert message")
	logger.Emerg("emerg message")

	logger.Close()

	logPath := filepath.Join(tmpDir, "test.log")
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	logStr := string(content)
	expectedSubstrings := []string{
		"[debug] debug message",
		"[info] info message",
		"[warn] warn message",
		"[error] error message",
		"[crit] crit message",
		"[alert] alert message",
		"[emerg] emerg message",
	}

	for _, sub := range expectedSubstrings {
		if !strings.Contains(logStr, sub) {
			t.Errorf("expected log to contain %q, but it didn't. Log content:\n%s", sub, logStr)
		}
	}
}

func TestLoggerReopen(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-log-reopen-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "reopen.log"
	conf.Logger.Level = "info"

	logger, _ := New(conf)
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	defer logger.Close()

	logger.Info("before reopen")

	logPath := filepath.Join(tmpDir, "reopen.log")
	rotatedPath := filepath.Join(tmpDir, "reopen.log.1")
	if err := os.Rename(logPath, rotatedPath); err != nil {
		t.Fatalf("failed to rename log file: %v", err)
	}

	logger.Reopen()

	logger.Info("after reopen")
	logger.Close()

	rotatedContent, err := os.ReadFile(rotatedPath)
	if err != nil {
		t.Fatalf("failed to read rotated log file: %v", err)
	}
	if !strings.Contains(string(rotatedContent), "before reopen") {
		t.Errorf("expected rotated log to contain 'before reopen'")
	}

	newContent, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read new log file: %v", err)
	}
	if !strings.Contains(string(newContent), "after reopen") {
		t.Errorf("expected new log to contain 'after reopen'")
	}
	if strings.Contains(string(newContent), "before reopen") {
		t.Errorf("expected new log NOT to contain 'before reopen'")
	}
}

func TestLoggerLevelFiltering(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "ups-agent-log-filter-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	conf := &config.Config{}
	conf.Logger.LogDir = tmpDir
	conf.Logger.FileName = "filter.log"
	conf.Logger.Level = "warn"

	logger, _ := New(conf)
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	defer logger.Close()

	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.Close()

	logPath := filepath.Join(tmpDir, "filter.log")
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	logStr := string(content)
	if strings.Contains(logStr, "[debug]") {
		t.Error("did not expect debug message to be logged")
	}
	if strings.Contains(logStr, "[info]") {
		t.Error("did not expect info message to be logged")
	}
	if !strings.Contains(logStr, "[warn] warn message") {
		t.Error("expected warn message to be logged")
	}
	if !strings.Contains(logStr, "[error] error message") {
		t.Error("expected error message to be logged")
	}
}
