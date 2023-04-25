package libnodejs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/libnodejs"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testPackageJSON(t *testing.T, context spec.G, it spec.S) {
	Expect := NewWithT(t).Expect

	var (
		path       string
		filePath   string
		workingDir string
	)

	it.Before(func() {
		workingDir = t.TempDir()

		path = workingDir
		filePath = filepath.Join(workingDir, "package.json")
		Expect(os.WriteFile(filePath, []byte(`{
			"scripts": {
				"poststart": "echo \"poststart\"",
				"prestart": "echo \"prestart\"",
				"start": "echo \"start\" && node server.js"
			}
		}`), 0600)).To(Succeed())
	})

	context("when parsing a valid package.json with start scripts", func() {
		it("successfully extracts the scripts information", func() {
			pkg, err := libnodejs.ParsePackageJSON(path)
			Expect(err).NotTo(HaveOccurred())

			Expect(pkg.Scripts.Start).To(ContainSubstring(`echo "start" && node server.js`))
			Expect(pkg.Scripts.PreStart).To(Equal(`echo "prestart"`))
			Expect(pkg.Scripts.PostStart).To(Equal(`echo "poststart"`))
		})
	})

	context("failure cases", func() {
		context("when the package.json is not a valid json file", func() {
			it.Before(func() {
				Expect(os.WriteFile(filePath, []byte(`%%%`), 0600)).To(Succeed())
			})

			it("fails parsing", func() {
				_, err := libnodejs.ParsePackageJSON(path)
				Expect(err).To(HaveOccurred())
			})
		})

		context("when the path to package.json is invalid", func() {
			it("fails parsing", func() {
				_, err := libnodejs.ParsePackageJSON("/tmp/non-existent")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	context("HasStartScript", func() {
		context("when a start script is present", func() {
			it("indicates that the package.json file has a start script", func() {
				pkg, err := libnodejs.ParsePackageJSON(path)
				Expect(err).NotTo(HaveOccurred())

				Expect(pkg.HasStartScript()).To(BeTrue())
			})
		})

		context("when a start script is NOT present", func() {
			it.Before(func() {
				Expect(os.WriteFile(filePath, []byte(`{}`), 0600)).To(Succeed())
			})

			it("indicates that the package.json file does not have a start script", func() {
				pkg, err := libnodejs.ParsePackageJSON(path)
				Expect(err).NotTo(HaveOccurred())

				Expect(pkg.HasStartScript()).To(BeFalse())
			})
		})
	})
}
