package gsl

// Log is an interface for logging.
type Log interface {
	With(fields map[string]interface{}) Log
	Info(message string)
	Infof(format string, v ...interface{})
	Infow(fields map[string]interface{}, message string)
	Infofw(fields map[string]interface{}, format string, v ...interface{})
	Error(message string)
	Errorf(format string, v ...interface{})
	Errorw(fields map[string]interface{}, message string)
	Errorfw(fields map[string]interface{}, format string, v ...interface{})
	V(verbosity int) bool
}
