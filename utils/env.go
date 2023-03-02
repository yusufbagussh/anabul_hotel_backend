package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvVar(key string, defaultVal string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
