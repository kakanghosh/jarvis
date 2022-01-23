package cmd

import (
	"bufio"
	"fmt"
	"github/com/kakanghosh/jarvis/service"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	new string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Long:  `This will run pre-setuped application`,
}

func init() {
	runCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("[name] expected for %s", runCmd.Use)
		}
		identifier := args[0]
		app := service.GetAppByNameOrSerial(identifier)
		if app == nil {
			return fmt.Errorf("app not found [%s]", identifier)
		}

		cmdStr := fmt.Sprintf("cd %s; %s", app.WorkingDirectory, app.Command)
		execCommand := exec.Command("sh", "-c", cmdStr)
		stdout, err := execCommand.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(stdout)
		go func() {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()

		if err := execCommand.Start(); err != nil {
			log.Fatal(err)
		}

		if err := execCommand.Wait(); err != nil {
			log.Fatal(err)
		}
		return nil
	}
	runCmd.PersistentFlags().StringVarP(&new, "new", "n", "", "test child args")
	rootCmd.AddCommand(runCmd)
}
