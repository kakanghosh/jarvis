package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/color"
	"github/com/kakanghosh/jarvis/utils/emoji"
	"strconv"

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
		fmt.Printf("%s %s\n\n", color.YellowText("Showing app list"), emoji.MAGNIFIED_GLASS_RIGHT)
		if len(appList) == 0 {
			fmt.Printf("%s\n\n", color.RedText("No app found!"))
		} else {
			showAppListInOutput(appList, showDetails)
		}

		return nil
	}
	appListCmd.PersistentFlags().BoolVarP(&showDetails, "details", "d", false, "Show list in details")
	rootCmd.AddCommand(appListCmd)
}

func showAppListInOutput(appList []model.AppFlags, showDetails bool) {
	for i, app := range appList {
		showSingleAppInOutput(i+1, &app, showDetails)
		fmt.Println()
	}
}

func showSingleAppInOutput(position int, appFlags *model.AppFlags, showDetails bool) {
	fmt.Printf("%s. %s\n", color.BlueText(strconv.Itoa(position)), color.GreenText(appFlags.Name))
	if showDetails {
		if len(appFlags.WorkingDirectory) == 0 {
			fmt.Printf("   %s  %s \n", emoji.OPEN_FILE_FOLDER, emoji.RED_CROSS_MARK)
		} else {
			fmt.Printf("   %s  %s\n", emoji.OPEN_FILE_FOLDER, color.CyanText(appFlags.WorkingDirectory))
		}
		fmt.Printf("   %s  %s\n", emoji.DOLLAR, color.YellowText(appFlags.Command))
	}
}
