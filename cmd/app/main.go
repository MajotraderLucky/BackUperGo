package main

import (
	"backupergo/internal/config"
	"backupergo/internal/executor"
	"backupergo/internal/filemanager"
	"fmt"
)

func convertConfig(src config.Config) config.Config {
	return config.Config{
		MysqlPath: src.MysqlPath,
		DBQuery:   src.DBQuery,
		DBName:    src.DBName,
	}
}

func main() {
	configPath := "config/config.json"
	pathsFile := "paths.txt"

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	executorCfg := convertConfig(cfg)
	newPaths, err := executor.GetPathsFromCommand(executorCfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = filemanager.UpdatePathsFile(pathsFile, newPaths)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Paths file has been updated successfully.")
}
