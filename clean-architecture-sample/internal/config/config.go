package config

import (
	"os"
)

type Config struct {
	LogLevel string `json:"log_level"`
	DataFile string `json:"data_file"`
}

func LoadConfig() *Config {
	return &Config{
		LogLevel: getEnv("LOG_LEVEL", "info"),
		DataFile: getEnv("USER_DATA_FILE", "data/users.json"),
	}
}

// 環境変数を取得し、デフォルト値を設定するヘルパー関数
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
