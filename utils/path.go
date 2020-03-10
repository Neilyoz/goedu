package utils

import (
	"os"
	"strings"
)

// GetCurrentPath 获取当前工作路径
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}
