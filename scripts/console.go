package ardent

import (
	"strings"
	"bufio"
	"os"
	"runtime"
	"os/exec"
)

// Get input from console
func ConsoleRead() (string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	return text
}

// Clear console
func ConsoleClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}