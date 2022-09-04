package version

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

var (
	ver *version.Version

	// GITCOMMIT GITCOMMIT
	GITCOMMIT = "unknown"

	// GITBRANCH GITBRANCH
	GITBRANCH = "unknown"

	// GITSTATE GITSTATE
	GITSTATE = "unknown"

	// BUILDTIME BUILDTIME
	BUILDTIME = "unknown"

	// VERSION VERSION
	VERSION = "v0"
)

func init() {
	ver = version.Must(version.NewVersion(VERSION))
}

// Version Version
func Version() *version.Version {
	return ver
}

// GitCommit GitCommit
func GitCommit() string {
	return GITCOMMIT
}

// GitBranch GitBranch
func GitBranch() string {
	return GITBRANCH
}

// GitState GitState
func GitState() string {
	return GITSTATE
}

// BuildDate BuildDate
func BuildDate() string {
	return BUILDTIME
}

// String String
func String() (s string) {
	s += fmt.Sprintf("%s\n", Version().String())
	s += fmt.Sprintf("Build Date: %s\n", BuildDate())
	s += fmt.Sprintf("Git Branch: %s\n", GitBranch())
	s += fmt.Sprintf("Git Commit: %s", GitCommit())
	return
}
