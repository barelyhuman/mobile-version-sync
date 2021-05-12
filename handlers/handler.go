package handlers

type VersionHandler interface {
	Setup() error
	GetCurrentVersion() string
	GetCurrentBuild() string
	SetCurrentVersion(versionString string) error
	IncrementCurrentBuild(buildNumber string) error
}
