package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"

	"github.com/spf13/cobra"
)

var (
	showDetails bool
	appListCmd  = &cobra.Command{
		Use:   "app-list",
		Short: "Get app list",
		Long:  `This will show list of app`,
	}
)

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
				if showDetails {
					fmt.Printf("%d. %s [%s] [%s]\n\n", i+1, appList[i].Name, appList[i].WorkingDirectory, appList[i].Command)
				} else {
					fmt.Printf("%d. %s\n\n", i+1, appList[i].Name)
				}
			}
		}

		return nil
	}
	appListCmd.PersistentFlags().BoolVarP(&showDetails, "details", "d", false, "Show list in details")
	rootCmd.AddCommand(appListCmd)
}
