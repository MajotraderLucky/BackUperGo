package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MysqlPath string `json:"mysqlPath"`
	DBQuery   string `json:"dbQuery"`
	DBName    string `json:"dbName"`
	PathsFile string `json:"pathsFile"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
