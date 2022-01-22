package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/utils"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   utils.GetAppName(),
		Short: "J.A.R.V.I.S",
		Long:  `JUST A REALLY VERY INTELLIGENT SYSTEM`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	utils.CreateDirectoryifNotExist(utils.RootDirectory())
	utils.CreateFileIfNotExist(utils.AppsFileLocation(), `[]`)
}
