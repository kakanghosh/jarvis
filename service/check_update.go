package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"github/com/kakanghosh/jarvis/utils/color"
	"github/com/kakanghosh/jarvis/utils/emoji"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/google/go-github/v42/github"
)

func getLastUpdateCheck() (*model.CheckUpdate, error) {
	checkUpdate := model.CheckUpdate{}
	fileData, err := ioutil.ReadFile(utils.CheckUpdateFileLocation())
	if err != nil {
		return &checkUpdate, err
	}
	err = json.Unmarshal(fileData, &checkUpdate)
	return &checkUpdate, err
}

func updateLastUpdateCheck() error {
	checkUpdate := model.CheckUpdate{
		LastChecked: utils.GetTodaysDateString(),
	}
	jsonFile, _ := json.MarshalIndent(checkUpdate, "", "  ")
	return os.WriteFile(utils.CheckUpdateFileLocation(), jsonFile, 0644)
}

func CheckUpdate() {
	if !isUpdateAlreadyChecked() {
		version := "v" + utils.GetVersion()
		latestRelease, _, _ := getLatestRelease()
		if latestRelease != nil && *latestRelease.TagName != version {
			fmt.Printf("Update available: %s\n", color.GreenText(*latestRelease.TagName))
			fmt.Printf("To update run `$ %s`\n", color.CyanText(GetUpdateCommand()))
		}
		updateLastUpdateCheck()
	}
}

func DoSelfUpdate() {
	version := "v" + utils.GetVersion()
	latestRelease, _, _ := getLatestRelease()
	if latestRelease == nil {
		log.Fatal(fmt.Sprintf("%s\n", color.GreenText("Release not found.")))
	}
	if *latestRelease.TagName != version {
		downloadLink := getLatestReleaseDownloadLink(latestRelease)
		workingDirectory, _ := os.Getwd()
		releaseDownloadFilePath := workingDirectory + string(os.PathSeparator) + utils.GetAppName()
		downloadLatestRelease(releaseDownloadFilePath, downloadLink)
		moveReleaseToCmdPath(releaseDownloadFilePath)
		fmt.Printf("%s from %s to %s %s\n", color.YellowText("Update completed"), color.RedText(version), color.GreenText(*latestRelease.TagName), emoji.GREEN_CHECK_MARK)
	} else {
		fmt.Printf("%s %s\n", color.GreenText("Release is upto date."), emoji.GREEN_CHECK_MARK)
	}
}

func moveReleaseToCmdPath(releaseDownloadFilePath string) {
	err := os.Chmod(releaseDownloadFilePath, 0775)
	if err != nil {
		log.Fatal(err)
	}

	cmdPath, err := exec.LookPath(utils.GetAppName())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(releaseDownloadFilePath, cmdPath)
	if err != nil {
		log.Fatal(err)
	}
}

func isUpdateAlreadyChecked() bool {
	lastUpdateCheck, _ := getLastUpdateCheck()
	return strings.Compare(lastUpdateCheck.LastChecked, utils.GetTodaysDateString()) == 0
}

func getLatestReleaseDownloadLink(latestRelease *github.RepositoryRelease) string {
	currentOS := runtime.GOOS
	if runtime.GOOS != "linux" {
		currentOS = fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	}
	if len(latestRelease.Assets) > 0 {
		for _, release := range latestRelease.Assets {
			if strings.Contains(*release.Name, currentOS) {
				return release.GetBrowserDownloadURL()
			}
		}
	}
	log.Fatalf("release not found for os %s\n", currentOS)
	return ""
}

func getLatestRelease() (*github.RepositoryRelease, *github.Response, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	return client.Repositories.GetLatestRelease(ctx, utils.GetRepoOwnerName(), utils.GetAppName())
}

func downloadLatestRelease(releaseDownloadFilePath, url string) {
	fmt.Printf("%s: %s\n", color.YellowText("Downloading release from"), color.BlueText(url))
	_, err := utils.DownloadFile(releaseDownloadFilePath, url)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUpdateCommand() string {
	return fmt.Sprintf("%s update %s", utils.GetAppName(), utils.GetAppName())
}
