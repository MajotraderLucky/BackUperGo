package service

import (
	"backupergo/internal/config"
	"backupergo/internal/executor"
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
