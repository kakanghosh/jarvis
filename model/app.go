package model

import "fmt"

type AppFlags struct {
	Serial           int    `json:"-"`
	Name             string `json:"name"`
	WorkingDirectory string `json:"workingDirectory"`
	Command          string `json:"command"`
}

func (flags *AppFlags) IsFlagsValid() bool {
	return len(flags.Name) > 0 && len(flags.Command) > 0
}

func (flags *AppFlags) Error() string {
	missingFlags := ""
	if len(flags.Name) == 0 {
		missingFlags += "[name]"
	}
	if len(flags.Command) == 0 {
		missingFlags += "[cmd]"
	}
	return fmt.Sprintf("flags %s expected", missingFlags)
}

type UpdateAppFlags struct {
	AppFlags
}

type CheckUpdate struct {
	LastChecked string `json:"lastChecked"`
}
