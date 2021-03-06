package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"

	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Update existing task",
	Long:  `This will update existing task`,
}

var updateTaskFlags = model.UpdateTaskFlags{}

func updateTaskRunE(cmd *cobra.Command, args []string) error {
	if updateTaskFlags.Serial <= 0 {
		return fmt.Errorf("[serial] expected for %s", updateTaskCmd.Use)
	}
	task := service.GetTaskBySerial(updateTaskFlags.Serial)
	if task == nil {
		return fmt.Errorf("task not found")
	}
	_, err := service.UpdateTask(&updateTaskFlags)
	return err
}

func updateTaskPostRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Task updated successfully %s\n", emoji.GREEN_CHECK_MARK)
}

func init() {
	updateTaskCmd.RunE = updateTaskRunE
	updateTaskCmd.PostRun = updateTaskPostRun

	updateTaskCmd.PersistentFlags().IntVarP(&updateTaskFlags.Serial, "serial", "s", 0, "Serial number of the task")
	updateTaskCmd.PersistentFlags().StringVarP(&updateTaskFlags.Name, "name", "n", "", "Task name")
	updateTaskCmd.PersistentFlags().StringVarP(&updateTaskFlags.WorkingDirectory, "directory", "d", "-1", "Working directory")
	updateTaskCmd.PersistentFlags().StringVarP(&updateTaskFlags.Command, "cmd", "c", "", "Command to execute the task")

	rootCmd.AddCommand(updateTaskCmd)
}
