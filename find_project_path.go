package libnodejs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FindProjectPath will validate that project path exists and is valid relative to the
// working directory.
func FindProjectPath(workingDir string) (string, error) {
	projectPath := os.Getenv(ProjectPathEnvName)
	if projectPath == "" {
		return workingDir, nil
	}

	path := filepath.Clean(filepath.Join(workingDir, projectPath))
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("could not find project path %q: %w", path, err)
		}

		return "", err
	}

	return path, nil
}
