package cmd

import (
	"fmt"
	"github/com/kakanghosh/jarvis/utils"
	"github/com/kakanghosh/jarvis/utils/color"
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	extension   string
	fileTidyCmd = &cobra.Command{
		Use:   "file-tidy",
		Short: "Tidy up the random files in folder",
		Long:  "Tidy up the random files in folder",
	}
)

func init() {
	fileTidyCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(extension) == 0 {
			return fmt.Errorf("extension is required")
		}
		files, err := utils.GetFilesInWorkingDirectory()
		if err != nil {
			return err
		}
		fileList := utils.FilterFilesByExtension(files, extension)
		if len(fileList) > 0 {
			showListOfFilesFoundWithExtenion(fileList, extension)

			defaultNewDirectoryName := fmt.Sprintf("%s-%s", utils.GetTodaysDateString(), extension)
			showOptionForChoosingFolderName(defaultNewDirectoryName)

			text := utils.ReadStringFromTerminal()
			option, err := strconv.Atoi(text)
			if err != nil {
				return err
			}

			newDirectoryName := ""
			switch option {
			case 1:
				newDirectoryName = defaultNewDirectoryName
			case 2:
				fmt.Printf("Enter valid directory name: ")
				text = utils.ReadStringFromTerminal()
				if utils.ValidateDirectoryName(text) {
					newDirectoryName = text
				} else {
					return fmt.Errorf("invalid directory name: %s", text)
				}
			default:
				log.Fatalf("Invalid option %d", option)
			}
			workingDirectory, _ := os.Getwd()
			utils.MoveFilesToNewLocation(workingDirectory, newDirectoryName, fileList)
			fmt.Printf("%s\n", color.GreenText("File moving completed"))
		} else {
			fmt.Printf("%s\n", color.RedText("Files not found!"))
		}
		return nil
	}
	fileTidyCmd.PersistentFlags().StringVarP(&extension, "extension", "e", "", "Extension of the file example pdf, jpeg, png, xlxs etc")
	rootCmd.AddCommand(fileTidyCmd)
}

func showOptionForChoosingFolderName(defaultNewDirectoryName string) {
	fmt.Printf("Choose option for folder name to move into\n")
	fmt.Printf("1. Default (%s)\n", color.YellowText(defaultNewDirectoryName))
	fmt.Printf("2. Custom\n")
}

func showListOfFilesFoundWithExtenion(fileList []fs.FileInfo, extension string) {
	fmt.Printf("List of file found with extension %s\n", color.GreenText(extension))
	for index, file := range fileList {
		fmt.Printf("%d. %s\n", index+1, color.BlueText(file.Name()))
	}
}
