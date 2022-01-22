package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"

	"github.com/spf13/cobra"
)

var rmAppCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove application",
	Long:  `This will remove application`,
}

func init() {
	rmAppCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("[name] expected for %s", rmAppCmd.Use)
		}
		appName := args[0]
		if app := service.GetAppByName(appName); app != nil {
			if err := service.RemoveApp(app); err == nil {
				fmt.Printf("App removed %s\n", appName)
			} else {
				return fmt.Errorf("failed to remove app %s", appName)
			}
		} else {
			return fmt.Errorf("'%s' app not found", appName)
		}
		return nil
	}
	rootCmd.AddCommand(rmAppCmd)
}
