package utils

import (
	"fmt"
	"os"
)

var (
	appName          = "jarvis"
	version          = "1.0.1"
	appsFileName     = "apps.json"
	rootDirectory    = fmt.Sprintf("%s/.jarvis", os.Getenv("HOME"))
	appsFileLocation = fmt.Sprintf("%s/%s", rootDirectory, appsFileName)
)

func GetAppName() string {
	return appName
}

func RootDirectory() string {
	return rootDirectory
}

func AppsFileLocation() string {
	return appsFileLocation
}

func GetVersion() string {
	return version
}
