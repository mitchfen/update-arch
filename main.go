package main

import (
	"fmt"
	"github.com/mitchfen/update-arch/helpers"
	"strings"
)

func main() {
	helpers.WriteSection("Updating pacman packages...")
	commandOutput, err := helpers.RunCommand("sudo", "pacman", "-Syu", "--noconfirm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Updating flatpak packages...")
	commandOutput, err = helpers.RunCommand("flatpak", "update", "--noninteractive")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Removing orphan pacman packages...")
	commandOutput, _ = helpers.RunCommand("sudo", "pacman", "-Qdt", "--noconfirm")
	if commandOutput != "" {
		commandOutput, err = helpers.RunCommand("sudo", "pacman", "-Rns", commandOutput)
	} else {
		fmt.Println("No orphan packages to remove.")
	}

	helpers.WriteSection("Checking for latest BIOS version...")
	commandOutput, err = helpers.RunCommand("python", "/home/mitchfen/dev/scripts_and_configs/scripts/getLatestBiosVersion.py")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Getting current package counts...")
	commandOutput, err = helpers.RunCommand("sudo", "pacman", "-Q")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pacman packages: %d", strings.Count(commandOutput, "\n"))
	commandOutput, err = helpers.RunCommand("flatpak", "list")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nFlatpak packages: %d", strings.Count(commandOutput, "\n"))

	// TODO: Update aur packages using git pull
	// TODO: Update ohmyposh by after checking ~/.oh-my-posh-version.txt
}
