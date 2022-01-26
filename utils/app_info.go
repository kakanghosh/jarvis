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
	viper.SetDefault("APP_VERSION", "1.3.1")
	viper.SetDefault("@YEAR", time.Now().Year())
	viper.SetDefault("ROOT_DIRECTORY", fmt.Sprintf("%s/.jarvis", os.Getenv("HOME")))
	viper.SetDefault("APPS_FILE_LOCATION", fmt.Sprintf("%s/%s", viper.GetString("ROOT_DIRECTORY"), "apps.json"))
}

func GetAppName() string {
	return viper.GetString("APP_NAME")
}

func RootDirectory() string {
	return viper.GetString("ROOT_DIRECTORY")
}

func AppsFileLocation() string {
	return viper.GetString("APPS_FILE_LOCATION")
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
