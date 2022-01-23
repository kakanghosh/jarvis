package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/service"

	"github.com/spf13/cobra"
)

var addAppCmd = &cobra.Command{
	Use:   "add-app",
	Short: "Add new application",
	Long:  `This will setuped application`,
}

var addAppFlags = model.AppFlags{
	Name:             "",
	WorkingDirectory: "",
	Command:          "",
}

func init() {
	addAppCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if !addAppFlags.IsFlagsValid() {
			return &addAppFlags
		}
		return nil
	}

	addAppCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if app := service.GetAppByName(addAppFlags.Name); app != nil {
			return fmt.Errorf("app exist. name [%s]", addAppFlags.Name)
		}
		return service.AddApp(&addAppFlags)
	}

	addAppCmd.PostRun = func(cmd *cobra.Command, args []string) {
		fmt.Printf("App %s added successfully\n", addAppFlags.Name)
	}

	addAppCmd.PersistentFlags().StringVarP(&addAppFlags.Name, "name", "n", "", "Program name")
	addAppCmd.PersistentFlags().StringVarP(&addAppFlags.WorkingDirectory, "directory", "d", "", "Working directory")
	addAppCmd.PersistentFlags().StringVarP(&addAppFlags.Command, "cmd", "c", "", "Command to start the program")
	rootCmd.AddCommand(addAppCmd)
}
