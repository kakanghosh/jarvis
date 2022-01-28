package cmd

import (
	"bufio"
	"fmt"
	"github/com/kakanghosh/jarvis/service"
	"github/com/kakanghosh/jarvis/utils/emoji"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
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

		cmdStr := ""
		if len(app.WorkingDirectory) > 0 {
			cmdStr = fmt.Sprintf("cd %s; %s", app.WorkingDirectory, app.Command)
		} else {
			cmdStr = app.Command
		}

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

		fmt.Printf("Starting app %s %s\n", app.Name, emoji.PERSON_RUNNING)
		if err := execCommand.Start(); err != nil {
			log.Fatal(err)
		}

		if err := execCommand.Wait(); err != nil {
			log.Fatal(err)
		}
		return nil
	}
	rootCmd.AddCommand(runCmd)
}
