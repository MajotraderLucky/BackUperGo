package main

import (
	"backupergo/internal/config"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	MysqlPath string `json:"mysqlPath"`
	DBQuery   string `json:"dbQuery"`
	DBName    string `json:"dbName"`
}

func getPathsFromCommand(config Config) ([]string, error) {
	cmd := exec.Command(config.MysqlPath, "-Nse", config.DBQuery, config.DBName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %s", err)
	}
	return strings.Split(strings.TrimSpace(string(output)), "\n"), nil
}

// updatePathsFile обновляет файл с путями, если есть изменения
func updatePathsFile(filePath string, newPaths []string) error {
	existingPaths := make(map[string]bool)

	// Чтение существующих путей из файла
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

	// Определение изменений
	newPathsMap := make(map[string]bool)
	for _, path := range newPaths {
		newPathsMap[path] = true
	}

	// Проверка необходимости обновления файла
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
		// Обновление файла
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

func main() {
	configPath := "config/config.json" // Путь к файлу конфигурации
	pathsFile := "paths.txt"           // Путь к файлу, в который записываются пути

	// Загрузка конфигурации
	config, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Получение новых путей
	newPaths, err := getPathsFromCommand(Config(config))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Обновление файла с путями
	err = updatePathsFile(pathsFile, newPaths)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Paths file has been updated successfully.")
}
