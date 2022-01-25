package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"

	"github.com/spf13/cobra"
)

var updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Update existing application",
	Long:  `This will update existing application`,
}

var updateAppFlags = model.UpdateAppFlags{}

func init() {
	updateAppCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if updateAppFlags.Serial <= 0 {
			return fmt.Errorf("[serial] expected for %s", updateAppCmd.Use)
		}
		app := service.GetAppBySerial(updateAppFlags.Serial)
		if app == nil {
			return fmt.Errorf("app not found")
		}
		_, err := service.UpdateApp(&updateAppFlags)
		return err
	}

	updateAppCmd.PostRun = func(cmd *cobra.Command, args []string) {
		fmt.Printf("App updated successfully %s\n", emoji.GREEN_CHECK_MARK)
	}
	updateAppCmd.PersistentFlags().IntVarP(&updateAppFlags.Serial, "serial", "s", 0, "Serial number of the app")
	updateAppCmd.PersistentFlags().StringVarP(&updateAppFlags.Name, "name", "n", "", "Program name")
	updateAppCmd.PersistentFlags().StringVarP(&updateAppFlags.WorkingDirectory, "directory", "d", "-1", "Working directory")
	updateAppCmd.PersistentFlags().StringVarP(&updateAppFlags.Command, "cmd", "c", "", "Command to start the program")

	rootCmd.AddCommand(updateAppCmd)
}
