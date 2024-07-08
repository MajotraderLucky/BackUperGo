package main

import (
	"backupergo/internal/config"
	"backupergo/internal/filemanager"
	"backupergo/internal/service"
	"fmt"
)

func main() {
	configPath := "config/config.json"

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	pathsFile := cfg.PathsFile
	newPaths, err := service.LoadAndProcessPaths(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = filemanager.UpdatePathsFile(pathsFile, newPaths)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Paths file has been updated successfully.")
}
