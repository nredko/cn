package logger

type LogLevel int8

const (
	LogDisabled = iota
	LogDebug
	LogInfo
	LogWarn
	LogError
)

type Logger interface {
	Errorf(string, ...interface{})
	Warningf(string, ...interface{})
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
}
