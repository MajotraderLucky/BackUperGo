package main

import (
	"backupergo/internal/config"
	"backupergo/internal/executor"
	"bufio"
	"fmt"
	"os"
)

func updatePathsFile(filePath string, newPaths []string) error {
	existingPaths := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			existingPaths[scanner.Text()] = true
		}
		file.Close()
	}

	newPathsMap := make(map[string]bool)
	for _, path := range newPaths {
		newPathsMap[path] = true
	}

	changed := len(existingPaths) != len(newPathsMap)
	if !changed {
		for path := range existingPaths {
			if !newPathsMap[path] {
				changed = true
				break
			}
		}
	}

	if changed {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		for path := range newPathsMap {
			_, err := file.WriteString(path + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

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

	err = updatePathsFile(pathsFile, newPaths)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Paths file has been updated successfully.")
}
