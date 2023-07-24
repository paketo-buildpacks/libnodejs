package libnodejs_test

import (
	"testing"

	"github.com/paketo-buildpacks/libnodejs"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testResolveNodeVersion(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		planner draft.Planner
	)

	it.Before(func() {
		planner = draft.NewPlanner()
	})

	it("Should respect the priorities and return the proper Node.js version", func() {
		entriesTests := []struct {
			Entries []packit.BuildpackPlanEntry
		}{
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": "package.json"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "=16", "version-source": "BP_NODE_VERSION"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "=16", "version-source": "package.json"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "=16", "version-source": ".node-version"},
					},
				},
			},
		}

		for _, tt := range entriesTests {
			entry, _ := libnodejs.ResolveNodeVersion(
				planner.Resolve,
				packit.BuildpackPlan{
					Entries: tt.Entries,
				})
			Expect(entry).To(Equal(tt.Entries[len(tt.Entries)-1]))
		}
	})

	it("Should respect the priorities and return the proper Node.js version - reversed", func() {
		entriesTests := []struct {
			Entries []packit.BuildpackPlanEntry
		}{
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "=16", "version-source": "BP_NODE_VERSION"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": "package.json"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": "package.json"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
					},
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
					},
				},
			},
			{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name:     "node",
						Metadata: map[string]interface{}{"version": "=16", "version-source": ".node-version"},
					},
				},
			},
		}

		for _, tt := range entriesTests {
			entry, _ := libnodejs.ResolveNodeVersion(
				planner.Resolve,
				packit.BuildpackPlan{
					Entries: tt.Entries,
				})
			Expect(entry).To(Equal(tt.Entries[0]))
		}
	})

	it("Returns entry with no name no plan entry", func() {
		entry, _ := libnodejs.ResolveNodeVersion(
			planner.Resolve,
			packit.BuildpackPlan{
				Entries: []packit.BuildpackPlanEntry{},
			})
		Expect(entry.Name).To(Equal(""))
	})

	it("Returns candidates correctly", func() {
		plan := packit.BuildpackPlan{
			Entries: []packit.BuildpackPlanEntry{
				{
					Name:     "node",
					Metadata: map[string]interface{}{"version": "=16", "version-source": "BP_NODE_VERSION"},
				},
				{
					Name:     "node",
					Metadata: map[string]interface{}{"version": "<0", "version-source": "package.json"},
				},
				{
					Name:     "node",
					Metadata: map[string]interface{}{"version": "<0", "version-source": ".nvmrc"},
				},
				{
					Name:     "node",
					Metadata: map[string]interface{}{"version": "<0", "version-source": ".node-version"},
				}},
		}

		_, candidates := libnodejs.ResolveNodeVersion(
			planner.Resolve,
			plan,
		)
		Expect(candidates).To(Equal(plan.Entries))
	})
}
