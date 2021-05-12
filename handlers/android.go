package handlers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type AndroidVersionHandler struct {
	Directory string
	AppName   string
}

var parsedGradleFile string

func (handler AndroidVersionHandler) Setup() error {
	var err error
	parsedGradleFile, err = readGradleFile(handler.Directory + "/build.gradle")
	if err != nil {
		return err
	}
	return nil
}

func (handler AndroidVersionHandler) GetCurrentVersion() string {
	tokens := parseToken(parsedGradleFile)
	result := contructAST(tokens)
	version := result.Find("versionName")
	return version.Children[0].Name
}

func (handler AndroidVersionHandler) SetCurrentVersion(versionString string) error {
	currentVersion := handler.GetCurrentVersion()
	parsedGradleFile = strings.Replace(parsedGradleFile, ("versionName " + currentVersion), ("versionName " + versionString), 1)
	return ioutil.WriteFile(handler.Directory+"/build.gradle", []byte(parsedGradleFile), os.ModePerm)
}

func (handler AndroidVersionHandler) GetCurrentBuild() string {
	tokens := parseToken(parsedGradleFile)
	result := contructAST(tokens)
	version := result.Find("versionCode")
	return version.Children[0].Name
}

func (handler AndroidVersionHandler) IncrementCurrentBuild(versionString string) error {
	currentBuild := handler.GetCurrentBuild()
	asInteger, err := strconv.ParseInt(currentBuild, 10, 32)
	if err != nil {
		return err
	}
	parsedGradleFile = strings.Replace(parsedGradleFile, ("versionCode " + currentBuild), ("versionCode " + fmt.Sprintf("%v", asInteger+1)), 1)
	return ioutil.WriteFile(handler.Directory+"/build.gradle", []byte(parsedGradleFile), os.ModePerm)
}
