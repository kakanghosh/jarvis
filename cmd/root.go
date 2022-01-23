package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
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
	utils.CreateDirectoryifNotExist(utils.RootDirectory())
	utils.CreateFileIfNotExist(utils.AppsFileLocation(), `[]`)
}
