package utils

import (
	"os"
	"strconv"
)

func GetStringFromEnv(key string, v string) string {
	value := os.Getenv(key)
	if value == "" {
		value = v
	}

	return value
}

func GetIntFromEnv(key string, v int) int {
	env := os.Getenv(key)

	// 未设置，则返回默认值
	if env == "" {
		return v
	}

	val, err := strconv.Atoi(env)
	// 转换失败，返回默认值
	if err != nil {
		return v
	}

	return val
}
