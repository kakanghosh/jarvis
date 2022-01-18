package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of JARVIS",
	Long:  `All software has versions. This is JARVIS's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("JARVIS v1.0.0")
	},
}
