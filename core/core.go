// Package core provides an entry point to use Xray core functionalities.
//
// Xray makes it possible to accept incoming network connections with certain
// protocol, process the data, and send them through another connection with
// the same or a difference protocol on demand.
//
// It may be configured to work with multiple protocols at the same time, and
// uses the internal router to tunnel through different inbound and outbound
// connections.
package core

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/xtls/xray-core/common/serial"
)

var (
	Version_x byte = 26
	Version_y byte = 9
	Version_z byte = 9
)

var (
	build               = "Custom"
	distributionVersion = "1.0.0"
)

const distributionName = "Auray"

func init() {
	// Manually injected
	if build != "Custom" {
		return
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	var isDirty bool
	var foundBuild bool
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			if len(setting.Value) < 7 {
				return
			}
			build = setting.Value[:7]
			foundBuild = true
		case "vcs.modified":
			isDirty = setting.Value == "true"
		}
	}
	if isDirty && foundBuild {
		build += "-dirty"
	}
}

// Version returns Xray's version as a string, in the form of "x.y.z" where x, y and z are numbers.
// ".z" part may be omitted in regular releases.
func Version() string {
	return fmt.Sprintf("%v.%v.%v", Version_x, Version_y, Version_z)
}

// DistributionName returns the name of this Xray-core distribution.
func DistributionName() string {
	return distributionName
}

// DistributionVersion returns Auray's own release version.
// It is intentionally separate from Version(), which remains the upstream
// Xray-core version used by compatibility-sensitive code.
func DistributionVersion() string {
	return distributionVersion
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() []string {
	return []string{
		serial.Concat(DistributionName(), " ", DistributionVersion(), " ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
	}
}
