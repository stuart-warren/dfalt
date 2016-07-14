package dfalt

import (
	"os"
	"strconv"
	"time"
)

func EnvString(envKey, def string) string {
	ret := os.Getenv(envKey)
	if ret == "" {
		ret = def
	}
	return ret
}

func EnvInt(envKey string, def int) int {
	ret, err := strconv.Atoi(os.Getenv(envKey))
	if err != nil {
		ret = def
	}
	return ret
}

func EnvBool(envKey string, def bool) bool {
	ret, err := strconv.ParseBool(os.Getenv(envKey))
	if err != nil {
		ret = def
	}
	return ret
}

func EnvDuration(envKey string, def time.Duration) time.Duration {
	ret, err := time.ParseDuration(os.Getenv(envKey))
	if err != nil {
		ret = def
	}
	return ret
}

func EnvFloat64(envKey string, def float64) float64 {
	ret, err := strconv.ParseFloat(os.Getenv(envKey), 64)
	if err != nil {
		ret = def
	}
	return ret
}

func EnvInt64(envKey string, def int64) int64 {
	ret, err := strconv.ParseInt(os.Getenv(envKey), 10, 64)
	if err != nil {
		ret = def
	}
	return ret
}

func EnvUint64(envKey string, def uint64) uint64 {
	ret, err := strconv.ParseUint(os.Getenv(envKey), 10, 64)
	if err != nil {
		ret = def
	}
	return ret
}
