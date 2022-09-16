package utils

//logrus
import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log
var Log = logrus.New()

func init() {
	Log.SetReportCaller(true)

	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

// SetLogLevel
//
// # Set log level
//
// Parameters
//
//	level: log level
func SetLogLevel(level string) {
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}

// SetLogFormat
//
// # Set log format
//
// Parameters
//
//	format: log format
func SetLogFormat(format string) {
	switch format {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		Log.SetFormatter(&logrus.TextFormatter{})
	default:
		Log.SetFormatter(&logrus.TextFormatter{})
	}
}

// SetLogOutput
//
// # Set log output
//
// Parameters
//
//	output: log output
func SetLogOutput(output string) {
	switch output {
	case "stdout":
		Log.SetOutput(os.Stdout)
	case "stderr":
		Log.SetOutput(os.Stderr)
	default:
		Log.SetOutput(os.Stdout)
	}
}
