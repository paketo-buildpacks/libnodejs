package libnodejs

import (
	"fmt"
	"os"
	"strconv"
)

// ShouldLaunchWithTini reports whether BP_LAUNCH_WITH_TINI is set to a true
// boolean value. Returns false when the variable is unset.
func ShouldLaunchWithTini() (bool, error) {
	if value, found := os.LookupEnv(LaunchWithTiniEnvName); found {
		enable, err := strconv.ParseBool(value)
		if err != nil {
			return false, fmt.Errorf("failed to parse %s value %s: %w", LaunchWithTiniEnvName, value, err)
		}
		return enable, nil
	}
	return false, nil
}
