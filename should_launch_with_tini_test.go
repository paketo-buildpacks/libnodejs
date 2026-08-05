package libnodejs_test

import (
	"testing"

	"github.com/paketo-buildpacks/libnodejs"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testShouldLaunchWithTini(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	it("returns false when the env var is unset", func() {
		result, err := libnodejs.ShouldLaunchWithTini()
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(BeFalse())
	})

	context("when BP_LAUNCH_WITH_TINI is true", func() {
		it.Before(func() {
			t.Setenv("BP_LAUNCH_WITH_TINI", "true")
		})

		it("returns true", func() {
			result, err := libnodejs.ShouldLaunchWithTini()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeTrue())
		})
	})

	context("when BP_LAUNCH_WITH_TINI is false", func() {
		it.Before(func() {
			t.Setenv("BP_LAUNCH_WITH_TINI", "false")
		})

		it("returns false", func() {
			result, err := libnodejs.ShouldLaunchWithTini()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeFalse())
		})
	})

	context("when BP_LAUNCH_WITH_TINI is malformed", func() {
		it.Before(func() {
			t.Setenv("BP_LAUNCH_WITH_TINI", "not-a-bool")
		})

		it("returns an error", func() {
			_, err := libnodejs.ShouldLaunchWithTini()
			Expect(err).To(MatchError(ContainSubstring("failed to parse BP_LAUNCH_WITH_TINI value not-a-bool")))
		})
	})
}
