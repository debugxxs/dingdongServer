package tools

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string   `json:"app_name"`
	AppMode  string   `json:"app_mode"`
	AppHost  string   `json:"app_host"`
	AppPort  string   `json:"app_port"`
	Database Database `json:"database"`
}

type Database struct {
	Drive   string `json:"drive"`
	DbName  string `json:"db_name"`
	DbUser  string `json:"db_user"`
	DbPass  string `json:"db_pass"`
	ShowSql bool   `json:"show_sql"`
	CharSet string `json:"char_set"`
}

var cfg *Config

//ParsingConfig 解析配置
func ParsingConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
