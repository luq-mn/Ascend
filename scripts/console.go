package ardent

import (
	"bufio"
	"os"
)

func ConsoleRead() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}