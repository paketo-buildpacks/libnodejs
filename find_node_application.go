package libnodejs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindNodeApplication(workingDir string) (string, error) {

	projectPath, err := FindProjectPath(workingDir)
	if err != nil {
		return "", err
	}

	launchpoint := os.Getenv(LaunchPointEnvName)
	if launchpoint != "" {
		if _, err := os.Stat(filepath.Join(workingDir, launchpoint)); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return "", fmt.Errorf("expected value derived from BP_LAUNCHPOINT [%s] to be an existing file", launchpoint)
			}

			return "", err
		}

		return filepath.Clean(launchpoint), nil
	}

	files := []string{"server.js", "app.js", "main.js", "index.js"}
	for _, file := range files {
		_, err := os.Stat(filepath.Join(projectPath, file))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return "", err
		}

		return filepath.Join(os.Getenv(ProjectPathEnvName), file), nil
	}

	return "", fmt.Errorf("could not find app in %s: expected one of %s", filepath.Clean(projectPath), strings.Join(files, " | "))
}
