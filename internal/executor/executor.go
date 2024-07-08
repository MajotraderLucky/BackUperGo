package executor

import (
	"backupergo/internal/config"
	"fmt"
	"os/exec"
	"strings"
)

func GetPathsFromCommand(cfg config.Config) ([]string, error) {
	cmd := exec.Command(cfg.MysqlPath, "-Nse", cfg.DBQuery, cfg.DBName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %s", err)
	}
	return strings.Split(strings.TrimSpace(string(output)), "\n"), nil
}
