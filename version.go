package main

var (
	// Version release version
	Version = "0.0.1"

	// Build will be overwritten automatically by the build system
	Build = "-dev"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

// FullVersion display the full version and build
func FullVersion() string {
	return Version + Build + " (" + GitCommit + ")"
}
