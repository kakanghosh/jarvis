package utils

import (
	"fmt"
	"os"
)

var appName = "jarvis"
var appsFileName = "apps.json"
var rootDirectory = fmt.Sprintf("%s/.jarvis", os.Getenv("HOME"))
var appsFileLocation = fmt.Sprintf("%s/%s", rootDirectory, appsFileName)

func GetAppName() string {
	return appName
}

func RootDirectory() string {
	return rootDirectory
}

func AppsFileLocation() string {
	return appsFileLocation
}
