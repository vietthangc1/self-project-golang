package envx

import (
	"os"
	"strconv"
	"strings"
)

func envClean(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func String(key string, defaultValue string) string {
	envValue := envClean(key)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}


func Int(key string, defaultValue int32) int32 {
	envValue := envClean(key)
	if envValue == "" {
		return defaultValue
	}
	envValueInt, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}
	return int32(envValueInt)
}

