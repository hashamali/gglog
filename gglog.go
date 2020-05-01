package gglog

// GGLog is an interface for logging.
type GGLog interface {
	With(metadata map[string]interface{}) GGLog
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}
