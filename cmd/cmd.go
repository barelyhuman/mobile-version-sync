package cmd

import (
	"flag"
	"fmt"

	"github.com/barelyhuman/mobile-version-sync/handlers"
)

type Options struct {
	platform      *string
	directory     *string
	version       *string
	bumpBuildCode *bool
}

var options Options

var flatAliases = map[string]string{
	"platform": "p",
	"dir":      "d",
	"version":  "v",
	"bump":     "b",
}

func Setup() bool {
	options.platform = flag.String("platform", "ios", "platform to run the version sync on [ios|android]")
	options.directory = flag.String("dir", "./ios/", "directory that contains the android or iOS structure")
	options.version = flag.String("version", "", "version string to sync with")
	options.bumpBuildCode = flag.Bool("bump", false, "bump the build/versionCode as well (default: false)")

	for from, to := range flatAliases {
		flagSet := flag.Lookup(from)
		flag.Var(flagSet.Value, to, fmt.Sprintf("alias to %s", flagSet.Name))
	}

	flag.Parse()

	return true
}

func Run() error {
	selectedPlatfrom := *options.platform
	var versionHandler handlers.VersionHandler

	switch selectedPlatfrom {
	case "android":
		{
			versionHandler = handlers.AndroidVersionHandler{}
			break
		}
	case "ios":
		{
			versionHandler = handlers.IOSVersionHandler{}
			break
		}
	default:
		{
			flag.Usage()
			return fmt.Errorf("invalid platform")
		}
	}

	IncreaseVersion(versionHandler)

	if !*options.bumpBuildCode {
		return nil
	}

	BumpBuild(versionHandler)

	return nil
}

func IncreaseVersion(versionHandler handlers.VersionHandler) error {
	currentVersion := versionHandler.GetCurrentVersion(*options.directory)
	if currentVersion == *options.version {
		return fmt.Errorf("version already in sync? Did you want to bump the version code/ build number instead?")
	}
	return versionHandler.SetCurrentVersion(*options.version)
}

func BumpBuild(versionHandler handlers.VersionHandler) error {
	return versionHandler.IncrementCurrentBuild(versionHandler.GetCurrentBuild(*options.directory) + 1)
}
