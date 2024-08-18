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
	fmt.Println("")
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
		return "", fmt.Errorf("error running command: %v \n Output: %s", err, out)
	}
	return string(out), nil
}

func main() {

	writeSection("Updating pacman packages...")
	commandOutput, err := runCommand("sudo", "pacman", "-Syu", "--noconfirm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	writeSection("Updating flatpak packages...")
	commandOutput, err = runCommand("flatpak", "update", "--force")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	writeSection("Removing orphan pacman packages...")
	commandOutput, _ = runCommand("sudo", "pacman", "-Qdt", "--noconfirm")
	if commandOutput != "" {
		commandOutput, err = runCommand("sudo", "pacman", "-Rns", commandOutput)
	} else {
		fmt.Println("No orphan packages to remove.")
	}

	writeSection("Checking for latest BIOS version...")
	commandOutput, err = runCommand("python", "/home/mitchfen/dev/scripts_and_configs/scripts/getLatestBiosVersion.py")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	writeSection("Getting current package counts...")
	commandOutput, err = runCommand("sudo", "pacman", "-Q")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pacman packages: %d", strings.Count(commandOutput, "\n"))
	commandOutput, err = runCommand("flatpak", "list")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nFlatpak packages: %d", strings.Count(commandOutput, "\n"))

	// TODO: Update aur packages using git pull
	// TODO: Update ohmyposh by after checking ~/.oh-my-posh-version.txt
}
