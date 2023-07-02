package logger

import (
	"log"
	"os"
	"strings"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[string]int8{
	"DEBUG": -int8(LogDebugLevel),
	"INFO":  -int8(LogInfoLevel),
}

var cfg = func() zap.Config {
	var config zap.Config
	if os.Getenv("PROD") == "yes" {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(-LogInfoLevel)
		config.EncoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
			zapcore.LowercaseLevelEncoder(l, pae)
		}
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
			zapcore.CapitalColorLevelEncoder(l, pae)
		}
		config.Level = zap.NewAtomicLevelAt(zapcore.Level(-LogDebugLevel))
	}
	logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	if v, ok := levelMap[logLevel]; ok {
		log.Printf("Settings level: %s, int %v", logLevel, v)
		config.Level = zap.NewAtomicLevelAt(zapcore.Level(v))
	}
	return config
}()

var global *zap.Logger = func() *zap.Logger {
	var e error
	g, e := cfg.Build()
	if e != nil {
		log.Fatalf("Unable to construct logger %s", e.Error())
	}
	return g
}()

func ZapDefault() *zap.Logger {
	return global
}

func ZapNoStack() *zap.Logger {
	c := cfg
	c.DisableStacktrace = true
	z, e := c.Build()
	if e != nil {
		panic(e)
	}
	return z
}

func ZapConfigDefault() zap.Config {
	return cfg
}

func WithZap(z *zap.Logger) Logger {
	return zapr.NewLogger(z)
}
