// Package fortnitestats provides details for the Fortnite Stats applet.
package fortnitestats

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed fortnite_stats.star
var source []byte

// New creates a new instance of the Fortnite Stats applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "fortnite-stats",
		Name:        "Fortnite Stats",
		Author:      "naominori",
		Summary:     "View your Fortnite stats",
		Desc:        "Displays your Fortnite stats.",
		FileName:    "fortnite_stats.star",
		PackageName: "fortnitestats",
		Source:  source,
	}
}
