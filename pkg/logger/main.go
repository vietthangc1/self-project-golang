package logger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
)

const (
	LogDebugLevel = 1
	LogInfoLevel  = 0
	LogWarnLevel  = -1
	LogErrorLevel = -2
	LogFatalLevel = -3
)

type Logger = logr.Logger

func Factory(name string) Logger {
	return zapr.NewLogger(global)
}
