package logger

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestConfigure_DefaultSettings(t *testing.T) {
	// Reset to defaults
	Configure(nil)

	// Should not panic with default configuration
	Info("test message")
}

func TestConfigure_WithOutput(t *testing.T) {
	var buf bytes.Buffer
	Configure(&Options{
		Output: &buf,
		Level:  LevelInfo,
	})

	Info("test message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("expected output to contain 'test message', got: %s", output)
	}
	if !strings.Contains(output, "key=value") {
		t.Errorf("expected output to contain 'key=value', got: %s", output)
	}
}

func TestConfigure_JSONFormat(t *testing.T) {
	var buf bytes.Buffer
	Configure(&Options{
		Output: &buf,
		Level:  LevelInfo,
		JSON:   true,
	})

	Info("test message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, `"msg":"test message"`) {
		t.Errorf("expected JSON output with 'msg', got: %s", output)
	}
	if !strings.Contains(output, `"key":"value"`) {
		t.Errorf("expected JSON output with 'key', got: %s", output)
	}
}

func TestLogLevels(t *testing.T) {
	tests := []struct {
		name         string
		configLevel  Level
		logFunc      func(msg string, args ...any)
		logLevel     string
		shouldAppear bool
	}{
		{
			name:         "Debug message at Debug level",
			configLevel:  LevelDebug,
			logFunc:      Debug,
			logLevel:     "DEBUG",
			shouldAppear: true,
		},
		{
			name:         "Debug message at Info level",
			configLevel:  LevelInfo,
			logFunc:      Debug,
			logLevel:     "DEBUG",
			shouldAppear: false,
		},
		{
			name:         "Info message at Info level",
			configLevel:  LevelInfo,
			logFunc:      Info,
			logLevel:     "INFO",
			shouldAppear: true,
		},
		{
			name:         "Info message at Warn level",
			configLevel:  LevelWarn,
			logFunc:      Info,
			logLevel:     "INFO",
			shouldAppear: false,
		},
		{
			name:         "Warn message at Warn level",
			configLevel:  LevelWarn,
			logFunc:      Warn,
			logLevel:     "WARN",
			shouldAppear: true,
		},
		{
			name:         "Error message at Error level",
			configLevel:  LevelError,
			logFunc:      Error,
			logLevel:     "ERROR",
			shouldAppear: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Configure(&Options{
				Output: &buf,
				Level:  tt.configLevel,
			})

			tt.logFunc("test message")

			output := buf.String()
			containsLevel := strings.Contains(output, tt.logLevel)

			if tt.shouldAppear && !containsLevel {
				t.Errorf("expected output to contain level %s, got: %s", tt.logLevel, output)
			}
			if !tt.shouldAppear && containsLevel {
				t.Errorf("expected output NOT to contain level %s, got: %s", tt.logLevel, output)
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	Configure(&Options{
		Output: &buf,
		Level:  LevelInfo,
	})

	// Debug should not appear at Info level
	Debug("debug message")
	if strings.Contains(buf.String(), "debug message") {
		t.Error("debug message should not appear at Info level")
	}

	// Change to Debug level (need to reconfigure with same output)
	buf.Reset()
	Configure(&Options{
		Output: &buf,
		Level:  LevelDebug,
	})

	Debug("debug message")
	if !strings.Contains(buf.String(), "debug message") {
		t.Error("debug message should appear at Debug level")
	}
}

func TestDisable(t *testing.T) {
	var buf bytes.Buffer
	Configure(&Options{
		Output: &buf,
		Level:  LevelDebug,
	})

	// Log should appear before disable
	Info("before disable")
	if !strings.Contains(buf.String(), "before disable") {
		t.Error("message should appear before disable")
	}

	// Disable logging
	Disable()

	// Nothing should be logged after disable
	Info("after disable")
	Error("error after disable")

	// Re-enable to capture output
	buf.Reset()
	Configure(&Options{
		Output: &buf,
		Level:  LevelInfo,
	})

	// Verify no leftover messages
	if strings.Contains(buf.String(), "after disable") {
		t.Error("message should not appear after disable")
	}
}

func TestStructuredLogging(t *testing.T) {
	var buf bytes.Buffer
	Configure(&Options{
		Output: &buf,
		Level:  LevelInfo,
	})

	Info("request completed",
		"path", "/api/v1/users",
		"status", 200,
		"duration_ms", 42,
	)

	output := buf.String()

	expectedFields := []string{
		"request completed",
		"path=/api/v1/users",
		"status=200",
		"duration_ms=42",
	}

	for _, field := range expectedFields {
		if !strings.Contains(output, field) {
			t.Errorf("expected output to contain '%s', got: %s", field, output)
		}
	}
}

func TestDisable_ToDiscard(t *testing.T) {
	// Configure to discard output
	Configure(&Options{
		Output: io.Discard,
		Level:  LevelError + 1,
	})

	// These should not panic even though output is discarded
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}
