package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(filePath string, url string) (written int64, err error) {
	out, _ := os.Create(filePath)
	defer out.Close()
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	return io.Copy(out, response.Body)
}
