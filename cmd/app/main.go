package main

import (
	"backupergo/internal/config"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/robfig/cron/v3"
)

func runBackup(pathsFile, backupDir string) {
	file, err := os.Open(pathsFile)
	if err != nil {
		fmt.Println("Error opening paths file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		base := filepath.Base(path)
		dest := filepath.Join(backupDir, base)

		// Вызов rsync для каждого пути
		cmd := exec.Command("rsync", "-avz", path, dest)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to rsync %s: %v\n", base, err)
		} else {
			fmt.Printf("Successfully backed up %s to %s\n", base, dest)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading paths file:", err)
	}
}

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	c := cron.New()
	c.AddFunc("@daily", func() { runBackup(cfg.PathsFile, cfg.BackUpDir) })
	c.Start()

	// Блокирует основной поток, пока не будет завершен
	select {}
}
