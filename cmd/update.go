package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils"
	"github/com/kakanghosh/jarvis/utils/color"
	"strings"

	"github.com/spf13/cobra"
)

var selfUpdateCommand = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("Update %s", utils.GetAppName()),
	Long:  fmt.Sprintf("Update %s", utils.GetAppName()),
}

func doSelfUpdate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 || strings.Compare(utils.GetAppName(), args[0]) != 0 {
		return fmt.Errorf("invalid args run $ %s", color.CyanText(service.GetUpdateCommand()))
	}
	service.DoSelfUpdate()
	return nil
}

func init() {
	selfUpdateCommand.RunE = doSelfUpdate
	rootCmd.AddCommand(selfUpdateCommand)
}
