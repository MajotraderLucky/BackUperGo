package filemanager

import (
	"bufio"
	"os"
)

func UpdatePathsFile(filePath string, newPaths []string) error {
	existingPaths := make(map[string]bool)

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

	newPathsMap := make(map[string]bool)
	for _, path := range newPaths {
		newPathsMap[path] = true
	}

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
