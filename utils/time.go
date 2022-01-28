package utils

import (
	"fmt"
	"time"
)

func GetTodaysDateString() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	return fmt.Sprintf("%d-%d-%d", day, month, year)
}
