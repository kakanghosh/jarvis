package utils

import (
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils/color"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func CreateFileIfNotExist(fileLocation string, defaultValue string) {
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		os.WriteFile(fileLocation, []byte(defaultValue), 0644)
	}
}

func CreateDirectoryifNotExist(directoryPath string) {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err := os.Mkdir(directoryPath, 0744)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetFilesInWorkingDirectory() ([]fs.FileInfo, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadDir(workingDirectory)
}

func GetFileExtensionFromName(fileName string) string {
	extension := ""
	if parts := strings.Split(fileName, "."); len(parts) > 0 {
		extension = parts[len(parts)-1]
	}
	return extension
}

func GetExtensionsFromFiles(files []fs.FileInfo) []model.FileExtensionCounter {
	extensionMap := make(map[string]int)
	for _, file := range files {
		if !file.IsDir() {
			extension := GetFileExtensionFromName(file.Name())
			_, ok := extensionMap[extension]
			if ok {
				extensionMap[extension] += 1
			} else {
				extensionMap[extension] = 1
			}
		}
	}
	fileExtensionCounterList := make([]model.FileExtensionCounter, 0)
	for extension, counter := range extensionMap {
		fileExtensionCounterList = append(fileExtensionCounterList, model.FileExtensionCounter{
			Extension: extension,
			Counter:   counter,
		})
	}
	return fileExtensionCounterList
}

func FilterFilesByExtension(files []fs.FileInfo, extension string) []fs.FileInfo {
	fileList := make([]fs.FileInfo, 0)
	for _, file := range files {
		fileExtension := GetFileExtensionFromName(file.Name())
		if !file.IsDir() && strings.Compare(fileExtension, extension) == 0 {
			fileList = append(fileList, file)
		}
	}
	return fileList
}

func MoveFilesToNewLocation(workingDirectory string, newDirectoryName string, fileList []fs.FileInfo) {
	newDirectoryPath := workingDirectory + string(os.PathSeparator) + newDirectoryName
	CreateDirectoryifNotExist(newDirectoryPath)
	for _, file := range fileList {
		oldFilePath := workingDirectory + string(os.PathSeparator) + file.Name()
		newFilePath := newDirectoryPath + string(os.PathSeparator) + file.Name()
		os.Rename(oldFilePath, newFilePath)
		fmt.Printf("%s moved to %s\n", color.YellowText(oldFilePath), color.GreenText(newFilePath))
	}
}

func ValidateDirectoryName(directoryName string) bool {
	r, _ := regexp.Compile(`^(\w+[-]?)+$`)
	return r.MatchString(directoryName)
}
