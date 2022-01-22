package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"

	"github.com/spf13/cobra"
)

var appListCmd = &cobra.Command{
	Use:   "app-list",
	Short: "Get app list",
	Long:  `This will show list of app`,
}

func init() {
	appListCmd.RunE = func(cmd *cobra.Command, args []string) error {
		appList, err := service.GetApplist()
		if err != nil {
			return err
		}
		if len(appList) == 0 {
			fmt.Println("No app found!")
		} else {
			for i := 0; i < len(appList); i++ {
				fmt.Printf("%d. %s [%s] [%s]\n", i+1, appList[i].Name, appList[i].WorkingDirectory, appList[i].Command)
			}
		}

		return nil
	}
	rootCmd.AddCommand(appListCmd)
}
