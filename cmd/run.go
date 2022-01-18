package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Long:  `This will run pre-setuped application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application is begin to run...")
		go show("hello", 5)
		show("world", 5)
	},
}

func show(message string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s %d\n", message, i)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
