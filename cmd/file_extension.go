package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"github/com/kakanghosh/jarvis/utils/color"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	showFileCounter  bool
	fileExtensionCmd = &cobra.Command{
		Use:   "file-ext",
		Short: "Get all the file-extension in the working directory",
		Long:  "Get all the file-extension in the working directory",
	}
)

func init() {
	fileExtensionCmd.RunE = func(cmd *cobra.Command, args []string) error {
		files, err := utils.GetFilesInWorkingDirectory()
		if err != nil {
			return err
		}
		extensionCounterList := utils.GetExtensionsFromFiles(files)
		if len(extensionCounterList) > 0 {
			fmt.Printf("%s\n", color.GreenText("List of extension"))
			printExtensionList(extensionCounterList, showFileCounter)
		} else {
			fmt.Printf("%s\n", color.RedText("Extension not found!"))
		}
		return nil
	}
	fileExtensionCmd.PersistentFlags().BoolVarP(&showFileCounter, "counter", "c", false, "Show number of file found")
	rootCmd.AddCommand(fileExtensionCmd)
}

func printExtensionList(extensionCounterList []model.FileExtensionCounter, showFileCounter bool) {
	for index, item := range extensionCounterList {
		if showFileCounter {
			fmt.Printf("%d. %s (%s)\n", index+1, item.Extension, color.YellowText(strconv.Itoa(item.Counter)))
		} else {
			fmt.Printf("%d. %s \n", index+1, item.Extension)
		}
	}
}
