package envconfig

import (
	"os"
	"strconv"
)

func parseString(envName string) (string, bool) {
	if env := os.Getenv(envName); len(env) > 0 {
		return env, true
	}

	return "", false
}

func parseInt(envName string) (int, bool) {
	if val, ok := parseInt64(envName); ok {
		return int(val), true
	}
	return 0, false
}

func parseInt16(envName string) (int16, bool) {
	if val, ok := parseInt64(envName); ok {
		return int16(val), true
	}
	return 0, false
}

func parseInt32(envName string) (int32, bool) {
	if val, ok := parseInt64(envName); ok {
		return int32(val), true
	}
	return 0, false
}

func parseInt64(envName string) (int64, bool) {
	if envInt := os.Getenv(envName); len(envInt) > 0 {
		if i64, err := strconv.ParseInt(envInt, 10, 64); err == nil {
			return i64, true
		}
	}
	return 0, false
}

func parseBool(envName string) (bool, bool) {
	if envb := os.Getenv(envName); len(envb) > 0 {
		if b, err := strconv.ParseBool(envb); err == nil {
			return b, true
		}
	}
	return false, false
}
