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
	Short: "Execute task",
	Long:  `This will execute pre-setuped task`,
}

func runTaskE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("[name] expected for %s", runCmd.Use)
	}
	identifier := args[0]
	task := service.GetTaskByNameOrSerial(identifier)
	if task == nil {
		return fmt.Errorf("task not found [%s]", identifier)
	}

	cmdStr := ""
	if len(task.WorkingDirectory) > 0 {
		cmdStr = fmt.Sprintf("cd %s; %s", task.WorkingDirectory, task.Command)
	} else {
		cmdStr = task.Command
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

	fmt.Printf("Executing task %s %s\n", task.Name, emoji.PERSON_RUNNING)
	if err := execCommand.Start(); err != nil {
		log.Fatal(err)
	}

	if err := execCommand.Wait(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func init() {
	runCmd.RunE = runTaskE
	rootCmd.AddCommand(runCmd)
}
