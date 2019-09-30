package envconfig

import (
	"os"
	"strconv"
	"strings"
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

func parseSliceString(envName string) ([]string, bool) {
	if env := os.Getenv(envName); len(env) > 0 {
		strs := strings.Split(env, ",")
		return strs, true
	}
	return nil, false
}

func parseSliceInt64(envName string) ([]int64, bool) {
	if env := os.Getenv(envName); len(env) > 0 {
		strs := strings.Split(env, ",")
		ints := make([]int64, 0, len(strs))
		for _, v := range strs {
			val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, false
			}
			ints = append(ints, val)
		}
		return ints, true
	}
	return nil, false
}

func parseSliceInt32(envName string) ([]int32, bool) {
	if arr, ok := parseSliceInt64(envName); ok {
		arr32 := make([]int32, 0, len(arr))
		for _, v := range arr {
			arr32 = append(arr32, int32(v))
		}

		return arr32, true
	}
	return nil, false
}

func parseSliceInt16(envName string) ([]int16, bool) {
	if arr, ok := parseSliceInt64(envName); ok {
		arr16 := make([]int16, 0, len(arr))
		for _, v := range arr {
			arr16 = append(arr16, int16(v))
		}
		return arr16, true
	}
	return nil, false
}

func parseSliceInt(envName string) ([]int, bool) {
	if arr, ok := parseSliceInt64(envName); ok {
		arr8 := make([]int, 0, len(arr))
		for _, v := range arr {
			arr8 = append(arr8, int(v))
		}
		return arr8, true
	}
	return nil, false
}

func parseSliceBool(envName string) ([]bool, bool) {
	if env := os.Getenv(envName); len(env) > 0 {
		strs := strings.Split(env, ",")
		arr := make([]bool, 0, len(strs))
		for _, v := range strs {
			b, err := strconv.ParseBool(v)
			if err != nil {
				return nil, false
			}
			arr = append(arr, b)
		}
		return arr, true
	}
	return nil, false
}
