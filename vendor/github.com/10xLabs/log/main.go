package log

import (
	filtered "github.com/Fs02/logrus-filtered-formatter"
	"github.com/sirupsen/logrus"
)

var filterFields = []string{"password"}

// Fields for using in `WithFields`
type Fields map[string]interface{}

// Config ...
type Config struct {
	// Set DebugEnabled to true to use DebugLevel logging and to include function caller in logs
	DebugEnabled bool
	PrettyPrint  bool
}

var log *logrus.Logger = logrus.New()

// Setup ...
func Setup(config Config) {
	logLevel := logrus.InfoLevel
	if config.DebugEnabled {
		logLevel = logrus.DebugLevel
	}

	log = logrus.New()
	formatter := filtered.New(filterFields, &logrus.JSONFormatter{PrettyPrint: config.PrettyPrint})
	log.SetFormatter(formatter)
	log.SetLevel(logLevel)
	log.SetReportCaller(config.DebugEnabled)
}

// WithFields adds a struct of fields to the log entry.
func WithFields(fields Fields) *logrus.Entry {
	f := logrus.Fields{}
	for k, v := range fields {
		f[k] = v
	}
	return log.WithFields(f)
}

// WithError adds an error as single field to the log entry.
func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Printf ...
func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Info ...
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
