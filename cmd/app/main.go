package main

import (
	"backupergo/internal/config"
	"backupergo/internal/filemanager"
	"fmt"
	"os"
	"path/filepath"
)

func ensureDirectoriesExist(dirPath string, validPaths map[string]bool) error {
	for dirName := range validPaths {
		fullPath := filepath.Join(dirPath, dirName)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Если директория не существует, создаем ее
			if err := os.Mkdir(fullPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", fullPath, err)
			}
			fmt.Printf("Created directory: %s\n", fullPath)
		}
	}
	return nil
}

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
	err = ensureDirectoriesExist(backupDir, validPaths)
	if err != nil {
		fmt.Printf("Failed to ensure directories exist: %v\n", err)
		return
	}
}
