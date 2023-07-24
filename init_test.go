package libnodejs_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitLibnodejs(t *testing.T) {
	suite := spec.New("libnodejs", spec.Report(report.Terminal{}), spec.Sequential())
	suite("FindProjectPath", testFindProjectPath)
	suite("PackageJSON", testPackageJSON)
	suite("FindNodeApplication", testFindNodeApplication)
	suite("ResolveNodeVersion", testResolveNodeVersion)
	suite.Run(t)
}
