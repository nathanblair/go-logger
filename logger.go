package logger

import (
	"io"
	"log"
)

// Log defines shared logging methods
type Log interface {
	Debug(format string, v interface{})
	Warn(format string, v interface{})
	Error(format string, v interface{})
	Info(format string, v interface{})
	Log(format string, v interface{})
}

// Logger is an implementation of the Log interface
type Logger struct {
	debugLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	infoLogger  *log.Logger
	logger      *log.Logger
}

// Debug print a message meant for the DEBUG channel
func (logger *Logger) Debug(format string, v interface{}) {
	logger.debugLogger.Printf(format, v)
}

// Warn print a message meant for the WARN channel
func (logger *Logger) Warn(format string, v interface{}) {
	logger.warnLogger.Printf(format, v)
}

// Error print a message meant for the ERROR channel
func (logger *Logger) Error(format string, v interface{}) {
	logger.errorLogger.Printf(format, v)
}

// Info print a message meant for the INFO channel
func (logger *Logger) Info(format string, v interface{}) {
	logger.infoLogger.Printf(format, v)
}

// Log print a message meant for the LOG channel
func (logger *Logger) Log(format string, v interface{}) {
	logger.logger.Printf(format, v)
}

// New creates a sane-defaults Logger instance
func New(writer io.Writer, errorWriter io.Writer) Logger {
	return Logger{
		debugLogger: log.New(writer, "[DEBUG] ", log.LstdFlags),
		warnLogger:  log.New(writer, "[WARN]  ", log.LstdFlags),
		errorLogger: log.New(errorWriter, "[ERROR] ", log.LstdFlags),
		infoLogger:  log.New(writer, "[INFO]  ", log.LstdFlags),
		logger:      log.New(writer, "", log.LstdFlags),
	}
}
