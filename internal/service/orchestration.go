package service

import (
	"backupergo/internal/config"
	"backupergo/internal/executor"
	"backupergo/internal/filemanager"
	"backupergo/internal/util"
	"fmt"
)

// LoadAndProcessPaths загружает конфигурацию, преобразует ее и получает новые пути.
func LoadAndProcessPaths(configPath string) ([]string, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	executorCfg := util.ConvertConfig(cfg)
	newPaths, err := executor.GetPathsFromCommand(executorCfg)
	if err != nil {
		return nil, fmt.Errorf("error getting new paths: %w", err)
	}

	return newPaths, nil
}

// ManageDirectories encapsulates the directory management logic
func ManageDirectories(cfg config.Config) error {
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
