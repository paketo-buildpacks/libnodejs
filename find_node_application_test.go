package libnodejs_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/libnodejs"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testFindNodeApplication(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
	)

	it.Before(func() {
		workingDir = t.TempDir()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("finds the server.js application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "server.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "server.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "server.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the server.js application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("server.js")))
		})
	})

	context("finds the server.cjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "app.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "server.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "server.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the server.cjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("server.cjs")))
		})
	})

	context("finds the server.mjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "app.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "server.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the server.mjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("server.mjs")))
		})
	})

	context("finds the app.js application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "app.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the app.js application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("app.js")))
		})
	})

	context("finds the app.cjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the app.cjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("app.cjs")))
		})
	})

	context("finds the app.mjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "app.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the app.mjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("app.mjs")))
		})
	})

	context("finds the main.js application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "main.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the main.js application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("main.js")))
		})
	})

	context("finds the main.cjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the main.cjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("main.cjs")))
		})
	})

	context("finds the main.mjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "main.mjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the main.mjs application entrypoint successfully", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("main.mjs")))
		})
	})

	context("finds the index.js application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "index.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the index.js application entrypoint", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("index.js")))
		})
	})

	context("finds the index.cjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "index.cjs"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the index.cjs application entrypoint", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("index.cjs")))
		})
	})

	context("finds the index.mjs application entrypoint", func() {
		it.Before(func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "index.mjs"), nil, 0600)).To(Succeed())
		})

		it("finds the index.mjs application entrypoint", func() {
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("index.mjs")))
		})
	})

	context("when there is a launchpoint", func() {
		it.Before(func() {
			Expect(os.Mkdir(filepath.Join(workingDir, "src"), os.ModePerm)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "src", "launchpoint.js"), nil, 0600)).To(Succeed())
		})

		context("when the launchpoint file exists", func() {
			it("returns the highest priority file", func() {
				t.Setenv("BP_LAUNCHPOINT", "./src/launchpoint.js")
				file, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).NotTo(HaveOccurred())
				Expect(file).To(Equal(filepath.Join("src", "launchpoint.js")))
			})
		})

		context("when the launchpoint file does not exist", func() {
			it("returns the empty string and no error", func() {
				t.Setenv("BP_LAUNCHPOINT", "./no-such-file.js")
				file, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).To(MatchError(ContainSubstring("expected value derived from BP_LAUNCHPOINT [./no-such-file.js] to be an existing file")))
				Expect(file).To(Equal(""))
			})

			it("returns the empty string and no error if BP_VERIFY_LAUNCHPOINT is true", func() {
				t.Setenv("BP_LAUNCHPOINT", "./no-such-file.js")
				t.Setenv("BP_VERIFY_LAUNCHPOINT", "true")
				file, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).To(MatchError(ContainSubstring("expected value derived from BP_LAUNCHPOINT [./no-such-file.js] to be an existing file")))
				Expect(file).To(Equal(""))
			})

			it("returns the highest priority file if BP_VERIFY_LAUNCHPOINT is false", func() {
				t.Setenv("BP_LAUNCHPOINT", "./gen/no-such-file-yet.js")
				t.Setenv("BP_VERIFY_LAUNCHPOINT", "false")
				file, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).NotTo(HaveOccurred())
				Expect(file).To(Equal(filepath.Join("gen", "no-such-file-yet.js")))
			})
		})
	})

	context("when there is a project path", func() {
		it.Before(func() {
			Expect(os.Mkdir(filepath.Join(workingDir, "frontend"), os.ModePerm)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "frontend", "server.js"), nil, 0600)).To(Succeed())
			Expect(os.WriteFile(filepath.Join(workingDir, "frontend", "app.js"), nil, 0600)).To(Succeed())
		})

		it("returns the highest priority file", func() {
			t.Setenv("BP_NODE_PROJECT_PATH", "frontend")
			file, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(file).To(Equal(filepath.Join("frontend", "server.js")))
		})
	})

	context("when there is a project path but value specified does not exist", func() {
		it("returns a failure", func() {
			t.Setenv("BP_NODE_PROJECT_PATH", "frontend")
			_, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
		})
	})

	context("when no application can be found", func() {
		it("returns that application could not be found", func() {
			_, err := libnodejs.FindNodeApplication(workingDir)
			Expect(err).To(MatchError(fmt.Errorf("could not find app in %s: expected one of server.js | server.cjs | server.mjs | app.js | app.cjs | app.mjs | main.js | main.cjs | main.mjs | index.js | index.cjs | index.mjs", workingDir)))
		})
	})

	context("failure cases", func() {
		context("when the launchpoint cannot be stat'd", func() {
			it.Before(func() {
				Expect(os.Chmod(workingDir, 0000)).To(Succeed())
			})

			it.After(func() {
				Expect(os.Chmod(workingDir, os.ModePerm)).To(Succeed())
			})

			it("fails with helpful error", func() {
				t.Setenv("BP_LAUNCHPOINT", "something.js")
				_, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).To(MatchError(ContainSubstring("permission denied")))
			})
		})

		context("when the working dir cannot be stat'd", func() {
			it.Before(func() {
				Expect(os.Chmod(workingDir, 0000)).To(Succeed())
			})

			it.After(func() {
				Expect(os.Chmod(workingDir, os.ModePerm)).To(Succeed())
			})

			it("fails with helpful error", func() {
				_, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).To(MatchError(ContainSubstring("permission denied")))
			})
		})

		context("when the project path cannot be stat'd", func() {
			it.Before(func() {
				Expect(os.Chmod(workingDir, 0000)).To(Succeed())
			})

			it.After(func() {
				Expect(os.Chmod(workingDir, os.ModePerm)).To(Succeed())
			})

			it("fails with helpful error", func() {
				t.Setenv("BP_NODE_PROJECT_PATH", "frontend")
				_, err := libnodejs.FindNodeApplication(workingDir)
				Expect(err).To(MatchError(ContainSubstring("permission denied")))
			})
		})
	})
}
