package service

import (
	"encoding/json"
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"io/ioutil"
	"os"
	"strconv"
)

func GetApplist() ([]model.AppFlags, error) {
	appList := make([]model.AppFlags, 0)
	fileData, err := ioutil.ReadFile(utils.AppsFileLocation())
	if err != nil {
		return appList, err
	}
	err = json.Unmarshal(fileData, &appList)
	return appList, err
}

func GetAppByName(appName string) *model.AppFlags {
	appList, _ := GetApplist()
	for index, app := range appList {
		if app.Name == appName {
			app.Serial = index + 1
			return &app
		}
	}
	return nil
}

func GetAppBySerial(serial int) *model.AppFlags {
	appList, _ := GetApplist()
	if serial >= 1 && len(appList) >= serial {
		appList[serial-1].Serial = serial
		return &appList[serial-1]
	}
	return nil
}

func GetAppByNameOrSerial(identifier string) *model.AppFlags {
	appList, _ := GetApplist()
	for _, app := range appList {
		if app.Name == identifier {
			return &app
		}
	}
	appNo, err := strconv.Atoi(identifier)
	if err != nil {
		return nil
	}
	if appNo >= 1 && len(appList) >= appNo {
		appList[appNo-1].Serial = appNo
		return &appList[appNo-1]
	}
	return nil
}

func RemoveApp(app *model.AppFlags) error {
	appList, _ := GetApplist()
	newAppList := make([]model.AppFlags, 0)
	for _, eachApp := range appList {
		if eachApp.Name != app.Name {
			newAppList = append(newAppList, eachApp)
		}
	}
	return saveUpdatedAppList(newAppList)
}

func AddApp(app *model.AppFlags) error {
	appList, _ := GetApplist()
	appList = append(appList, *app)
	return saveUpdatedAppList(appList)
}

func UpdateApp(updateAppFlags *model.UpdateAppFlags) (*model.AppFlags, error) {
	appList, _ := GetApplist()

	foundedIndex := updateAppFlags.Serial - 1

	if len(updateAppFlags.Name) > 0 {
		if appWithUpdatedName := GetAppByName(updateAppFlags.Name); appWithUpdatedName != nil {
			if appWithUpdatedName.Serial != updateAppFlags.Serial {
				return nil, fmt.Errorf("duplicate name [%s]", updateAppFlags.Name)
			}
		}
		appList[foundedIndex].Name = updateAppFlags.Name
	}
	if updateAppFlags.WorkingDirectory != "-1" {
		appList[foundedIndex].WorkingDirectory = updateAppFlags.WorkingDirectory
	}
	if len(updateAppFlags.Command) > 0 {
		appList[foundedIndex].Command = updateAppFlags.Command
	}
	saveUpdatedAppList(appList)
	return &appList[foundedIndex], nil
}

func saveUpdatedAppList(appList []model.AppFlags) error {
	jsonFile, _ := json.MarshalIndent(appList, "", "  ")
	return os.WriteFile(utils.AppsFileLocation(), jsonFile, 0644)
}
