// Package logger provides a configurable structured logging system for the iRacing API client.
// It uses Go's standard library log/slog for professional-grade structured logging with
// support for multiple log levels and verbosity configuration.
package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
)

// Level represents the logging level.
type Level int

const (
	// LevelDebug is used for detailed debugging information.
	LevelDebug Level = Level(slog.LevelDebug)
	// LevelInfo is used for general operational information.
	LevelInfo Level = Level(slog.LevelInfo)
	// LevelWarn is used for warning messages.
	LevelWarn Level = Level(slog.LevelWarn)
	// LevelError is used for error messages.
	LevelError Level = Level(slog.LevelError)
	// LevelDisabled disables all logging output.
	LevelDisabled Level = Level(slog.LevelError + 1000)
)

// Options configures the logger behavior.
type Options struct {
	// Level sets the minimum log level to output. Messages below this level are discarded.
	Level Level
	// Output sets the destination for log output. Defaults to os.Stderr if nil.
	Output io.Writer
	// JSON enables JSON-formatted output when true. Defaults to text format.
	JSON bool
}

var (
	defaultLogger *slog.Logger
	mu            sync.RWMutex
)

func init() {
	// Initialize with default settings (Info level, text format, stderr output)
	defaultLogger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// Configure sets up the global logger with the provided options.
func Configure(opts *Options) {
	if opts == nil {
		opts = &Options{}
	}

	output := opts.Output
	if output == nil {
		output = os.Stderr
	}

	var handler slog.Handler
	handlerOpts := &slog.HandlerOptions{
		Level: slog.Level(opts.Level),
	}

	if opts.JSON {
		handler = slog.NewJSONHandler(output, handlerOpts)
	} else {
		handler = slog.NewTextHandler(output, handlerOpts)
	}

	mu.Lock()
	defaultLogger = slog.New(handler)
	mu.Unlock()
}

// Disable silences all log output by setting the level to LevelDisabled
// and directing output to io.Discard.
func Disable() {
	Configure(&Options{
		Output: io.Discard,
		Level:  LevelDisabled,
	})
}

// getLogger returns the current configured logger.
func getLogger() *slog.Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogger
}

// Debug logs a debug-level message with optional key-value pairs.
func Debug(msg string, args ...any) {
	getLogger().Debug(msg, args...)
}

// DebugContext logs a debug-level message with context and optional key-value pairs.
func DebugContext(ctx context.Context, msg string, args ...any) {
	getLogger().DebugContext(ctx, msg, args...)
}

// Info logs an info-level message with optional key-value pairs.
func Info(msg string, args ...any) {
	getLogger().Info(msg, args...)
}

// InfoContext logs an info-level message with context and optional key-value pairs.
func InfoContext(ctx context.Context, msg string, args ...any) {
	getLogger().InfoContext(ctx, msg, args...)
}

// Warn logs a warning-level message with optional key-value pairs.
func Warn(msg string, args ...any) {
	getLogger().Warn(msg, args...)
}

// WarnContext logs a warning-level message with context and optional key-value pairs.
func WarnContext(ctx context.Context, msg string, args ...any) {
	getLogger().WarnContext(ctx, msg, args...)
}

// Error logs an error-level message with optional key-value pairs.
func Error(msg string, args ...any) {
	getLogger().Error(msg, args...)
}

// ErrorContext logs an error-level message with context and optional key-value pairs.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	getLogger().ErrorContext(ctx, msg, args...)
}
