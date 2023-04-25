package libnodejs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/libnodejs"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testFindProjectPath(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
	)

	it.Before(func() {
		workingDir = t.TempDir()
		t.Setenv("BP_NODE_PROJECT_PATH", "custom/path")

		err := os.MkdirAll(filepath.Join(workingDir, "custom", "path"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
	})

	it("returns the project path", func() {
		result, err := libnodejs.FindProjectPath(workingDir)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(filepath.Join(workingDir, "custom", "path")))
	})

	context("failure cases", func() {
		context("when the project path subdirectory isn't accessible", func() {
			it.Before(func() {
				Expect(os.Chmod(workingDir, 0000)).To(Succeed())
			})

			it.After(func() {
				Expect(os.Chmod(workingDir, os.ModePerm)).To(Succeed())
			})

			it("returns an error", func() {
				_, err := libnodejs.FindProjectPath(workingDir)
				Expect(err).To(MatchError(ContainSubstring("permission denied")))
			})
		})

		context("when the project path subdirectory does not exist", func() {
			it.Before(func() {
				t.Setenv("BP_NODE_PROJECT_PATH", "some-garbage")
			})

			it("returns an error", func() {
				_, err := libnodejs.FindProjectPath(workingDir)
				Expect(err).To(MatchError(ContainSubstring("could not find project path")))
				Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
			})
		})
	})
}
