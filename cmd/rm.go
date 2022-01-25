package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"

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
			return fmt.Errorf("[name]/[serial] expected for %s", rmAppCmd.Use)
		}
		identifier := args[0]
		if app := service.GetAppByNameOrSerial(identifier); app != nil {
			if err := service.RemoveApp(app); err == nil {
				fmt.Printf("App removed %s %s\n", app.Name, emoji.GREEN_CHECK_MARK)
			} else {
				return fmt.Errorf("failed to remove app %s", app.Name)
			}
		} else {
			return fmt.Errorf("'%s' app not found", identifier)
		}
		return nil
	}
	rootCmd.AddCommand(rmAppCmd)
}
