package utils

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
