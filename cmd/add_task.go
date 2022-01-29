package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add-task",
	Short: "Add new task",
	Long:  `This will setuped new task`,
}

var addTaskFlags = model.TaskFlags{}

func addTaskPreRun(cmd *cobra.Command, args []string) error {
	if !addTaskFlags.IsFlagsValid() {
		return &addTaskFlags
	}
	return nil
}

func addTaskRun(cmd *cobra.Command, args []string) error {
	if task := service.GetTaskByName(addTaskFlags.Name); task != nil {
		return fmt.Errorf("task exist. name [%s]", addTaskFlags.Name)
	}
	return service.AddTask(&addTaskFlags)
}

func addTaskPostRun(cmd *cobra.Command, args []string) {
	fmt.Printf("%s added successfully %s\n", addTaskFlags.Name, emoji.GREEN_CHECK_MARK)
}

func init() {
	addTaskCmd.PreRunE = addTaskPreRun
	addTaskCmd.RunE = addTaskRun
	addTaskCmd.PostRun = addTaskPostRun

	addTaskCmd.PersistentFlags().StringVarP(&addTaskFlags.Name, "name", "n", "", "Task name")
	addTaskCmd.PersistentFlags().StringVarP(&addTaskFlags.WorkingDirectory, "directory", "d", "", "Working directory")
	addTaskCmd.PersistentFlags().StringVarP(&addTaskFlags.Command, "cmd", "c", "", "Command to execute the task")
	rootCmd.AddCommand(addTaskCmd)
}
