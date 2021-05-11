package handlers

type IOSVersionHandler struct {
	directory string
}

func (handler IOSVersionHandler) GetCurrentVersion(projectDir string) string {
	return ""
}

func (handler IOSVersionHandler) SetCurrentVersion(versionString string) error {
	return nil
}

func (handler IOSVersionHandler) GetCurrentBuild(projectDir string) string {
	return ""
}

func (handler IOSVersionHandler) IncrementCurrentBuild(versionString string) error {
	return nil
}
