package model

import "fmt"

type AddAppFlags struct {
	Name             string `json:"name"`
	WorkingDirectory string `json:"workingDirectory"`
	Command          string `json:"command"`
}

func (flags *AddAppFlags) IsFlagsValid() bool {
	return len(flags.Name) > 0 && len(flags.WorkingDirectory) > 0 && len(flags.Command) > 0
}

func (flags *AddAppFlags) Error() string {
	missingFlags := ""
	if len(flags.Name) == 0 {
		missingFlags += "[name]"
	}

	if len(flags.WorkingDirectory) == 0 {
		missingFlags += "[directory]"
	}

	if len(flags.Command) == 0 {
		missingFlags += "[cmd]"
	}
	return fmt.Sprintf("flags %s expected", missingFlags)
}
