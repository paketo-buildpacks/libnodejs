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
		PostStart string `json:"poststart"`
		PreStart  string `json:"prestart"`
		Start     string `json:"start"`
	} `json:"scripts"`
}

// ParsePackageJSON parses the contents of a package.json file.
func ParsePackageJSON(path string) (PackageJSON, error) {
	file, err := os.Open(filepath.Join(path, "package.json"))
	if err != nil {
		return PackageJSON{}, err
	}
	defer file.Close()

	var pkg PackageJSON
	err = json.NewDecoder(file).Decode(&pkg)
	if err != nil {
		return PackageJSON{}, fmt.Errorf("unable to decode package.json %w", err)
	}

	return pkg, nil
}

// HasStartScript indicates the presence of a start script in the package.json
// file.
func (pj PackageJSON) HasStartScript() bool {
	return pj.Scripts.Start != ""
}

func (pj PackageJSON) GetVersion() string {
	return pj.Engines.Node
}
