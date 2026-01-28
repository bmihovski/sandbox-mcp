package version

var (
	// These values are injected during build time using -ldflags
	Version   = "v1.0.3"
	CommitSHA = "none"
)

// GetVersion returns the current version of sandbox-mcp
func GetVersion() string {
	return Version
}

// GetCommitSHA returns the git commit hash of the build
func GetCommitSHA() string {
	return CommitSHA
}
