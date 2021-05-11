package handlers

type AndroidVersionHandler struct {
	Directory string
	AppName   string
}

func (handler AndroidVersionHandler) Setup() error {
	return nil
}

func (handler AndroidVersionHandler) GetCurrentVersion(projectDir string) string {
	return ""
}

func (handler AndroidVersionHandler) SetCurrentVersion(versionString string) error {
	return nil
}

func (handler AndroidVersionHandler) GetCurrentBuild(projectDir string) string {
	return ""
}

func (handler AndroidVersionHandler) IncrementCurrentBuild(versionString string) error {
	return nil
}
