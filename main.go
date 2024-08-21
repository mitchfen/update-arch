package main

import (
	"fmt"
	"github.com/mitchfen/update-arch/helpers"
	"strings"
)

const biosVersionScriptPath = "/home/mitchfen/dev/scripts_and_configs/scripts/getLatestBiosVersion.py"
const aurPath = "/home/mitchfen/aur"

func main() {
	helpers.WriteSection("Updating pacman packages...")
	commandOutput, err := helpers.RunCommand("sudo", "pacman", "-Syu", "--noconfirm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Removing orphan pacman packages...")
	commandOutput, _ = helpers.RunCommand("sudo", "pacman", "-Qdt", "--noconfirm")
	if commandOutput != "" {
		fmt.Println("Removing orphan packages: ", commandOutput)
		commandOutput, err = helpers.RunCommand("sudo", "pacman", "-Rns$(sudo pacman -Qdt)")
	} else {
		fmt.Println("No orphan packages to remove.")
	}

	helpers.WriteSection("Updating flatpak packages...")
	commandOutput, err = helpers.RunCommand("flatpak", "update", "--noninteractive")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Checking for latest BIOS version...")
	commandOutput, err = helpers.RunCommand("python", biosVersionScriptPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(commandOutput)

	helpers.WriteSection("Updating AUR packages...")
	err = helpers.UpdateAurPackages(aurPath)
	if err != nil {
		fmt.Println(err)
	}

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
}
