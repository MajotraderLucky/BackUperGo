package config

import (
	"bufio"
	"os"
	"path/filepath"
)

// LoadPaths читает пути из указанного файла и возвращает карту базовых имен директорий.
func LoadPaths(filePath string) (map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	paths := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		base := filepath.Base(path)
		paths[base] = true
	}

	return paths, scanner.Err()
}
