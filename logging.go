package irapi

import (
	"io"

	"github.com/riccardotornesello/irapi-go/client/logger"
)

// LogLevel represents the logging level for the iRacing API client.
type LogLevel = logger.Level

// Log level constants for configuring the logger.
const (
	LogLevelDebug = logger.LevelDebug
	LogLevelInfo  = logger.LevelInfo
	LogLevelWarn  = logger.LevelWarn
	LogLevelError = logger.LevelError
)

// LoggerOptions configures the logger behavior.
type LoggerOptions struct {
	// Level sets the minimum log level to output. Messages below this level are discarded.
	// Default is LogLevelInfo.
	Level LogLevel
	// Output sets the destination for log output. Defaults to os.Stderr if nil.
	Output io.Writer
	// JSON enables JSON-formatted output when true. Defaults to text format.
	JSON bool
}

// ConfigureLogger sets up the global logger with the provided options.
// If opts is nil, default settings are used (Info level, text format, stderr output).
//
// Example:
//
//	// Enable debug logging
//	irapi.ConfigureLogger(&irapi.LoggerOptions{
//	    Level: irapi.LogLevelDebug,
//	})
//
//	// Enable JSON logging to a custom writer
//	irapi.ConfigureLogger(&irapi.LoggerOptions{
//	    Level:  irapi.LogLevelInfo,
//	    Output: myWriter,
//	    JSON:   true,
//	})
func ConfigureLogger(opts *LoggerOptions) {
	if opts == nil {
		logger.Configure(nil)
		return
	}
	logger.Configure(&logger.Options{
		Level:  opts.Level,
		Output: opts.Output,
		JSON:   opts.JSON,
	})
}

// DisableLogging silences all log output from the iRacing API client.
func DisableLogging() {
	logger.Disable()
}
