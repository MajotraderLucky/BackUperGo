package main

import (
	"backupergo/internal/controller"
	"fmt"
)

func main() {
	configPath := "config/config.json"

	if err := controller.ProcessPaths(configPath); err != nil {
		fmt.Println(err)
	}
}
