package libnodejs_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitGoBuild(t *testing.T) {
	suite := spec.New("libnodejs", spec.Report(report.Terminal{}), spec.Sequential())
	suite("ProjectPathParser", testProjectPathParser)
	suite("PackageJsonParser", testPackageJsonParser)
	suite.Run(t)
}
