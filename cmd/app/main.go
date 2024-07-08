package main

import (
	"backupergo/internal/config"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func loadPaths(filePath string) (map[string]bool, error) {
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

func filterAndCleanDirectories(dirPath string, validPaths map[string]bool) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			dirName := file.Name()
			if !validPaths[dirName] && dirName != "web" && dirName != "mysql" {
				err := os.RemoveAll(filepath.Join(dirPath, dirName))
				if err != nil {
					return fmt.Errorf("failed to remove directory %s: %v", dirName, err)
				}
				fmt.Printf("Removed directory: %s\n", dirName)
			} else {
				fmt.Printf("Kept directory: %s\n", dirName)
			}
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

	validPaths, err := loadPaths(pathsFile)
	if err != nil {
		fmt.Printf("Failed to load paths: %v\n", err)
		return
	}

	err = filterAndCleanDirectories(backupDir, validPaths)
	if err != nil {
		fmt.Printf("Failed to filter and clean directories: %v\n", err)
		return
	}
}
