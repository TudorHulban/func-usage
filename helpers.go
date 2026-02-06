package funcusage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getModulePath(root string) (string, error) {
	modPath := filepath.Join(root, "go.mod")

	data, err := os.ReadFile(modPath) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	lines := strings.Split(string(data), "\n") //nolint:modernize

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	return "",
		errors.New("module path not found in go.mod")
}
