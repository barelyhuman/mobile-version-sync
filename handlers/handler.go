package handlers

type VersionHandler interface {
	Setup() error
	GetCurrentVersion(projectDirectory string) string
	GetCurrentBuild(projectDirectory string) string
	SetCurrentVersion(versionString string) error
	IncrementCurrentBuild(buildNumber string) error
}
