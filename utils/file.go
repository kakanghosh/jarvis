package utils

import (
	"log"
	"os"
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
