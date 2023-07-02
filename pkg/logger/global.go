// provide a global opinionated logger using zap
package logger

// global instance
var (
	logDefault   = Factory("ko")
	warnDefault  = logDefault.V(LogWarnLevel)
	debugDefault = logDefault.V(LogDebugLevel)
)

func Fatal(err error, msg string, keysAndValues ...interface{}) {
	logDefault.Error(err, msg, keysAndValues...)
	panic(err)
}

// FatalIf panic on error not empty
func FatalIf(err error, msg string, keysAndValues ...interface{}) {
	if err == nil {
		return
	}
	logDefault.Error(err, msg, keysAndValues...)
	panic(err)
}

// the default logger use Level INFO
func LOG() Logger {
	return logDefault
}

// the default logger use Level WARN
func WARN() Logger {
	return warnDefault
}

// the default logger use Level DEBUG
func DEBUG() Logger {
	return debugDefault
}

var Info = logDefault.Info
var Error = logDefault.Error
var Debug = debugDefault.Info
var Warn = warnDefault.Info
