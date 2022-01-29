package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"

	"github.com/spf13/cobra"
)

var rmTaskCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  `This will remove task`,
}

func removeRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("[name]/[serial] expected for %s", rmTaskCmd.Use)
	}
	identifier := args[0]
	if task := service.GetTaskByNameOrSerial(identifier); task != nil {
		if err := service.RemoveTask(task); err == nil {
			fmt.Printf("Task removed %s %s\n", task.Name, emoji.GREEN_CHECK_MARK)
		} else {
			return fmt.Errorf("failed to remove task %s", task.Name)
		}
	} else {
		return fmt.Errorf("'%s' task not found", identifier)
	}
	return nil
}

func init() {
	rmTaskCmd.RunE = removeRunE
	rootCmd.AddCommand(rmTaskCmd)
}
