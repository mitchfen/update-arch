package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

func WriteSection(msg string) {
	color := "\033[35m"  // Magenta color code
	color2 := "\033[36m" // Cyan color code
	divider := strings.Repeat("─", 80)
	fmt.Print(color) // Change color without showing anything
	fmt.Println(divider)
	fmt.Print(color2) // Change color without showing anything
	fmt.Println(msg)
	fmt.Print(color) // Change color without showing anything
	fmt.Println(divider)
	fmt.Print("\033[0m") // Reset color
}

func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v \n Output: %s", err, out)
	}
	return string(out), nil
}
