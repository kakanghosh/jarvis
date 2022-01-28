package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadStringFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
