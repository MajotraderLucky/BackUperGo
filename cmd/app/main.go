package main

import (
	"backupergo/internal/config"
	"backupergo/internal/service"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	err = service.ManageDirectories(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
