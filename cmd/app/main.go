package main

import (
	"backupergo/internal/config"
	"backupergo/internal/filemanager"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}
	pathsFile := cfg.PathsFile
	backupDir := cfg.BackUpDir

	validPaths, err := config.LoadPaths(pathsFile)
	if err != nil {
		fmt.Printf("Failed to load paths: %v\n", err)
		return
	}

	err = filemanager.FilterAndCleanDirectories(backupDir, validPaths)
	if err != nil {
		fmt.Printf("Failed to filter and clean directories: %v\n", err)
		return
	}

	// Добавлен вызов ensureDirectoriesExist после очистки директорий
	err = filemanager.EnsureDirectoriesExist(backupDir, validPaths)
	if err != nil {
		fmt.Printf("Failed to ensure directories exist: %v\n", err)
		return
	}
}
