package controller

import (
	"backupergo/internal/config"
	"backupergo/internal/filemanager"
	"backupergo/internal/service"
	"fmt"
)

// processPaths обрабатывает пути, загружая конфигурацию, обрабатывая новые пути и обновляя файл.
func ProcessPaths(configPath string) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	pathsFile := cfg.PathsFile
	newPaths, err := service.LoadAndProcessPaths(configPath)
	if err != nil {
		return fmt.Errorf("error processing paths: %v", err)
	}

	err = filemanager.UpdatePathsFile(pathsFile, newPaths)
	if err != nil {
		return fmt.Errorf("error updating paths file: %v", err)
	}

	fmt.Println("Paths file has been updated successfully.")
	return nil
}
