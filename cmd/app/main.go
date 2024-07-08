package main

import (
	"backupergo/internal/config"
	"backupergo/internal/filemanager"
	"fmt"
)

// manageDirectories encapsulates the directory management logic
func manageDirectories(cfg config.Config) error {
	pathsFile := cfg.PathsFile
	backupDir := cfg.BackUpDir

	validPaths, err := config.LoadPaths(pathsFile)
	if err != nil {
		return fmt.Errorf("failed to load paths: %v", err)
	}

	err = filemanager.FilterAndCleanDirectories(backupDir, validPaths)
	if err != nil {
		return fmt.Errorf("failed to filter and clean directories: %v", err)
	}

	err = filemanager.EnsureDirectoriesExist(backupDir, validPaths)
	if err != nil {
		return fmt.Errorf("failed to ensure directories exist: %v", err)
	}

	return nil
}

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	err = manageDirectories(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
