package model

import "fmt"

type TaskFlags struct {
	Serial           int    `json:"-"`
	Name             string `json:"name"`
	WorkingDirectory string `json:"workingDirectory"`
	Command          string `json:"command"`
}

func (flags *TaskFlags) IsFlagsValid() bool {
	return len(flags.Name) > 0 && len(flags.Command) > 0
}

func (flags *TaskFlags) Error() string {
	missingFlags := ""
	if len(flags.Name) == 0 {
		missingFlags += "[name]"
	}
	if len(flags.Command) == 0 {
		missingFlags += "[cmd]"
	}
	return fmt.Sprintf("flags %s expected", missingFlags)
}

type UpdateTaskFlags struct {
	TaskFlags
}

type FileExtensionCounter struct {
	Extension string
	Counter   int
}

type CheckUpdate struct {
	LastChecked string `json:"lastChecked"`
}
