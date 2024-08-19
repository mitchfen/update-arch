package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func WriteSection(msg string) {
	color := "\033[35m"  // Magenta color code
	color2 := "\033[36m" // Cyan color code
	divider := strings.Repeat("─", 80)
	fmt.Print("\n")  // Change color without showing anything
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

func UpdateAurPackages(aurPath string) error {
	err := os.Chdir(aurPath)
	if err != nil {
		return fmt.Errorf("failed to change to AUR directory: %w", err)
	}

	dirEntries, err := os.ReadDir(aurPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// Iterate over subdirectories
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			err = os.Chdir(dirEntry.Name())
			if err != nil {
				return fmt.Errorf("failed to change directory: %w", err)
			}

			// Execute `git pull`
			gitOutput, err := RunCommand("git", "pull")
			if err != nil {
				return fmt.Errorf("failed to execute git pull: %w", err)
			}

			// Check if the repository is up to date
			if strings.Contains(gitOutput, "Already up to date.") {
				fmt.Printf("%s is already up to date\n", dirEntry.Name())
			} else {
				fmt.Printf("%s is being updated...\n", dirEntry.Name())
				_, err = RunCommand("makepkg", "-si")
				if err != nil {
					return fmt.Errorf("failed to execute makepkg -si: %w", err)
				}

				// Execute `git clean -fxd`
				err = exec.Command("git", "clean", "-fxd").Run()
				if err != nil {
					return fmt.Errorf("failed to execute git clean -fxd: %w", err)
				}
			}

			err = os.Chdir(aurPath)
			if err != nil {
				return fmt.Errorf("failed to change back to parent directory: %w", err)
			}
		}
	}
	return err
}
