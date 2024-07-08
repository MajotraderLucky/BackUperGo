package filemanager

import (
	"fmt"
	"os"
	"path/filepath"
)

// FilterAndCleanDirectories удаляет ненужные директории из заданного пути.
func FilterAndCleanDirectories(dirPath string, validPaths map[string]bool) error {
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
