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
		Use:   "tasks",
		Short: "Get task list",
		Long:  `This will show list of task`,
	}
)

func showTasksRunE(cmd *cobra.Command, args []string) error {
	taskList, err := service.GetTasklist()
	if err != nil {
		return err
	}
	fmt.Printf("%s %s\n\n", color.YellowText("Showing task list"), emoji.MAGNIFIED_GLASS_RIGHT)
	if len(taskList) == 0 {
		fmt.Printf("%s\n\n", color.RedText("No task found!"))
	} else {
		showAppListOutput(taskList, showDetails)
	}

	return nil
}

func init() {
	appListCmd.RunE = showTasksRunE
	appListCmd.PersistentFlags().BoolVarP(&showDetails, "details", "d", false, "Show list in details")
	rootCmd.AddCommand(appListCmd)
}

func showAppListOutput(taskList []model.TaskFlags, showDetails bool) {
	for i, task := range taskList {
		showSingleTaskOutput(i+1, &task, showDetails)
		fmt.Println()
	}
}

func showSingleTaskOutput(position int, taskFlags *model.TaskFlags, showDetails bool) {
	fmt.Printf("%s. %s\n", color.BlueText(strconv.Itoa(position)), color.GreenText(taskFlags.Name))
	if showDetails {
		if len(taskFlags.WorkingDirectory) == 0 {
			fmt.Printf("   %s  %s \n", emoji.OPEN_FILE_FOLDER, emoji.RED_CROSS_MARK)
		} else {
			fmt.Printf("   %s  %s\n", emoji.OPEN_FILE_FOLDER, color.CyanText(taskFlags.WorkingDirectory))
		}
		fmt.Printf("   %s  %s\n", emoji.DOLLAR, color.YellowText(taskFlags.Command))
	}
}
