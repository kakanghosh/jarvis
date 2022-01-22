package service

import (
	"encoding/json"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"io/ioutil"
	"os"
)

func GetApplist() ([]model.AddAppFlags, error) {
	appList := make([]model.AddAppFlags, 0)
	fileData, err := ioutil.ReadFile(utils.AppsFileLocation())
	if err != nil {
		return appList, err
	}
	err = json.Unmarshal(fileData, &appList)
	return appList, err
}

func GetAppByName(appName string) *model.AddAppFlags {
	appList, _ := GetApplist()
	for _, app := range appList {
		if app.Name == appName {
			return &app
		}
	}
	return nil
}

func RemoveApp(app *model.AddAppFlags) error {
	appList, _ := GetApplist()
	newAppList := make([]model.AddAppFlags, 0)
	for _, eachApp := range appList {
		if eachApp.Name != app.Name {
			newAppList = append(newAppList, eachApp)
		}
	}
	jsonFile, _ := json.MarshalIndent(newAppList, "", "  ")
	return os.WriteFile(utils.AppsFileLocation(), jsonFile, 0644)
}

func AddApp(app *model.AddAppFlags) error {
	appList, _ := GetApplist()
	appList = append(appList, *app)
	jsonFile, _ := json.MarshalIndent(appList, "", "  ")
	return os.WriteFile(utils.AppsFileLocation(), jsonFile, 0644)
}
