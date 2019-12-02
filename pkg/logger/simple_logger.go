package logger

import (
	"io"
	"log"
	"os"
	"strings"
)

type simpleLogger struct {
	Logger   *log.Logger
	LogLevel LogLevel
}

func NewSimpleLogger(name string, out io.Writer) Logger {
	return &simpleLogger{
		Logger:   log.New(out, name+" ", log.LstdFlags|log.Lmicroseconds),
		LogLevel: logLevelFromEnvironment(),
	}
}

func (l *simpleLogger) Errorf(f string, v ...interface{}) {
	if l.LogLevel <= LogError {
		l.Logger.Printf("ERROR: "+f, v...)
	}
}

func (l *simpleLogger) Warningf(f string, v ...interface{}) {
	if l.LogLevel <= LogWarn {
		l.Logger.Printf("WARNING: "+f, v...)
	}
}

func (l *simpleLogger) Infof(f string, v ...interface{}) {
	if l.LogLevel <= LogInfo {
		l.Logger.Printf("INFO: "+f, v...)
	}
}

func (l *simpleLogger) Debugf(f string, v ...interface{}) {
	if l.LogLevel <= LogDebug {
		l.Logger.Printf("DEBUG: "+f, v...)
	}
}

func logLevelFromEnvironment() LogLevel {
	logLevel, _ := os.LookupEnv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "error":
		return LogError
	case "warn":
		return LogWarn
	case "info":
		return LogInfo
	case "debug":
		return LogDebug
	case "disabled":
		return LogDisabled
	}
	return LogInfo
}
