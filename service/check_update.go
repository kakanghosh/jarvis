package service

import (
	"encoding/json"
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"io/ioutil"
)

func GetLastUpdateCheck() (*model.CheckUpdate, error) {
	checkUpdate := model.CheckUpdate{}
	fileData, err := ioutil.ReadFile(utils.CheckUpdateFileLocation())
	if err != nil {
		return &checkUpdate, err
	}
	err = json.Unmarshal(fileData, &checkUpdate)
	return &checkUpdate, err
}

func CheckUpdate() {
	lastUpdateCheck, _ := GetLastUpdateCheck()
	fmt.Printf("last checked: %s\n", lastUpdateCheck.LastChecked)
	// version := "v" + utils.GetVersion()
	// client := github.NewClient(nil)
	// ctx := context.Background()
	// repoRelease, _, err := client.Repositories.GetLatestRelease(ctx, "kakanghosh", "jarvis")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if repoRelease != nil && *repoRelease.TagName != version {
	// 	fmt.Printf("Update available: %s\n", color.GreenText(*repoRelease.TagName))
	// 	currentOS := runtime.GOOS
	// 	if runtime.GOOS != "linux" {
	// 		currentOS = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	// 	}
	// 	if len(repoRelease.Assets) > 0 {
	// 		for _, release := range repoRelease.Assets {
	// 			if strings.Contains(*release.Name, currentOS) {
	// 				//fmt.Printf("%s\n", release.GetBrowserDownloadURL())
	// 				break
	// 			}
	// 		}
	// 	}

	// }
}
