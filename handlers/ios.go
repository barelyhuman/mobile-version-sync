package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"howett.net/plist"
)

type IOSVersionHandler struct {
	Directory string
	AppName   string
}

type IOSPlistProps struct {
	Props map[string]interface{}
}

var plistDetails IOSPlistProps

func (handler IOSVersionHandler) Setup() error {
	data, err := ioutil.ReadFile(handler.Directory + "/" + handler.AppName + "/Info.plist")
	buf := bytes.NewReader(data)
	if err != nil {
		log.Fatal(err)
	}

	decoder := plist.NewDecoder(buf)

	err = decoder.Decode(&plistDetails.Props)
	if err != nil {
		return err
	}
	return nil
}

func (handler IOSVersionHandler) GetCurrentVersion() string {
	return fmt.Sprintf("%v", plistDetails.Props["CFBundleShortVersionString"])
}

func (handler IOSVersionHandler) SetCurrentVersion(versionString string) error {
	plistDetails.Props["CFBundleShortVersionString"] = versionString

	modifiedPlist := bytes.NewBuffer([]byte(""))
	encoder := plist.NewEncoder(modifiedPlist)
	encoder.Indent("\t")
	err := encoder.Encode(plistDetails.Props)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(handler.Directory+"/"+handler.AppName+"/Info.plist", modifiedPlist.Bytes(), os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func (handler IOSVersionHandler) GetCurrentBuild() string {
	return fmt.Sprintf("%v", plistDetails.Props["CFBundleVersion"])
}

func (handler IOSVersionHandler) IncrementCurrentBuild(versionString string) error {
	versionCode, err := strconv.ParseInt(versionString, 10, 32)
	if err != nil {
		return err
	}
	versionCode += 1

	plistDetails.Props["CFBundleVersion"] = fmt.Sprintf("%v", versionCode)

	modifiedPlist := bytes.NewBuffer([]byte(""))
	encoder := plist.NewEncoder(modifiedPlist)
	encoder.Indent("\t")
	err = encoder.Encode(plistDetails.Props)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(handler.Directory+"/"+handler.AppName+"/Info.plist", modifiedPlist.Bytes(), os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}
