package libnodejs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// PackageJSON represents the contents of a package.json file.
type PackageJSON struct {
	Engines struct {
		Node string `json:"node"`
	} `json:"engines"`
	Scripts struct {
		PostStart string
		PreStart  string
		Start     string
	}
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`

	AllScripts map[string]string `json:"scripts"`
}

// ParsePackageJSON parses the contents of a package.json file.
func ParsePackageJSON(path string) (PackageJSON, error) {
	file, err := os.Open(filepath.Join(path, "package.json"))
	if err != nil {
		return PackageJSON{}, err
	}

	var pkg PackageJSON
	err = json.NewDecoder(file).Decode(&pkg)
	if err != nil {
		return PackageJSON{}, fmt.Errorf("unable to decode package.json %w", err)
	}

	if fileCloseErr := file.Close(); fileCloseErr != nil {
		return PackageJSON{}, fileCloseErr
	}

	startScriptName := os.Getenv(StartScriptNameEnvName)
	if startScriptName == "" {
		startScriptName = "start"
	} else {
		if pkg.AllScripts[startScriptName] == "" {
			return PackageJSON{}, fmt.Errorf("no script entry with name \"%s\" exists", startScriptName)
		}
	}

	pkg.Scripts.Start = pkg.AllScripts[startScriptName]
	pkg.Scripts.PreStart = pkg.AllScripts["prestart"]
	pkg.Scripts.PostStart = pkg.AllScripts["poststart"]

	return pkg, nil
}

// HasStartScript indicates the presence of a start script in the package.json
// file or as defined by .
func (pj PackageJSON) HasStartScript() bool {
	return pj.Scripts.Start != ""
}

func (pj PackageJSON) GetVersion() string {
	return pj.Engines.Node
}
