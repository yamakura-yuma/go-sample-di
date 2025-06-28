package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("USER_DATA_FILE", "data/test_config.json")
	cfg := LoadConfig()
	if cfg.LogLevel != "debug" {
		t.Errorf("expected debug, got %s", cfg.LogLevel)
	}
	if cfg.DataFile != "data/test_config.json" {
		t.Errorf("expected data/test_config.json, got %s", cfg.DataFile)
	}
}
