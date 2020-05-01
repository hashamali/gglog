package stlog

// Log is an interface for logging.
type Log interface {
	With(metadata interface{}) Log
	Info(message string)
	Infof(format string, v ...interface{})
	Infow(metadata interface{}, message string)
	Infofw(metadata interface{}, format string, v ...interface{})
	Error(message string)
	Errorf(format string, v ...interface{})
	Errorw(metadata interface{}, message string)
	Errorfw(metadata interface{}, format string, v ...interface{})
	V(verbosity int) bool
}
