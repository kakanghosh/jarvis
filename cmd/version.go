package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", utils.GetAppName()),
	Long:  fmt.Sprintf("All software has versions. This is %s's", utils.GetAppName()),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v1.0.1\n", utils.GetAppName())
	},
}
