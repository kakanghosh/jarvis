package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/utils"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", utils.GetAppName()),
	Long:  fmt.Sprintf("All software has versions. This is %s's", utils.GetAppName()),
	Run: func(cmd *cobra.Command, args []string) {
		user := utils.GetAuthor()
		currentYear := utils.GetCopyRightYear()
		fmt.Printf("%s %s @ %d\nMaintain by %s\n", utils.GetAppName(), utils.GetVersion(), currentYear, user)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
