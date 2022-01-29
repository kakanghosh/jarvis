package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("AUTHOR", "Kakan Ghosh <kakanghosh69@gmail.com>")
	viper.SetDefault("APP_NAME", "jarvis")
	viper.SetDefault("APP_VERSION", "2.0.0")
	viper.SetDefault("REPO_OWNER_NAME", "kakanghosh")
	viper.SetDefault("@YEAR", time.Now().Year())
	viper.SetDefault("ROOT_DIRECTORY", fmt.Sprintf("%s/.jarvis", os.Getenv("HOME")))
	viper.SetDefault("TASKS_FILE_LOCATION", fmt.Sprintf("%s/%s", viper.GetString("ROOT_DIRECTORY"), "tasks.json"))
	viper.SetDefault("CHECK_UPDATE_FILE_LOCATION", fmt.Sprintf("%s/%s", viper.GetString("ROOT_DIRECTORY"), "check_update.json"))
}

func GetAppName() string {
	return viper.GetString("APP_NAME")
}

func RootDirectory() string {
	return viper.GetString("ROOT_DIRECTORY")
}

func TasksFileLocation() string {
	return viper.GetString("TASKS_FILE_LOCATION")
}

func GetVersion() string {
	return viper.GetString("APP_VERSION")
}

func GetAuthor() string {
	return viper.GetString("AUTHOR")
}

func GetCopyRightYear() int {
	return viper.GetInt("@YEAR")
}

func CheckUpdateFileLocation() string {
	return viper.GetString("CHECK_UPDATE_FILE_LOCATION")
}

func GetRepoOwnerName() string {
	return viper.GetString("REPO_OWNER_NAME")
}
