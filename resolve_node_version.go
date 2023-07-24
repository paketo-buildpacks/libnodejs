package libnodejs

import (
	packit "github.com/paketo-buildpacks/packit/v2"
)

// Resolve the Node.js version to be used based on the possible sources in
// priority order
type nodeResolver func(name string, entries []packit.BuildpackPlanEntry, priorities []interface{}) (packit.BuildpackPlanEntry, []packit.BuildpackPlanEntry)

func ResolveNodeVersion(resolver nodeResolver, plan packit.BuildpackPlan) (packit.BuildpackPlanEntry, []packit.BuildpackPlanEntry) {

	priorities := []interface{}{
		"BP_NODE_VERSION",
		"package.json",
		".nvmrc",
		".node-version",
	}

	return resolver("node", plan.Entries, priorities)
}
