package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func writeSection(msg string) {
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

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v", err)
	}
	return string(out), nil
}

func main() {
	//scriptsPath := "~/dev/scripts_and_configs/scripts"

	writeSection("Updating pacman packages...")
	commandOutput, err := runCommand("sudo", "pacman", "-Syu")
	fmt.Println(commandOutput)
	if err != nil {
		fmt.Println(err)
	}

	writeSection("Updating flatpak packages...")
	commandOutput, err = runCommand("flatpak", "update")
	fmt.Println(commandOutput)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: Update aur packages using git pull
	// TODO: Update ohmyposh by after checking ~/.oh-my-posh-version.txt
	// TODO: Remove unused pacman packages
	// TODO: Run python script to check for BIOS update
	// TODO: Print out latest package counts

}
