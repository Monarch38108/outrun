package main

import (
	"os"
	"os/exec"
	"runtime"
)

// clearScreen clears the terminal screen based on the operating system
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Command to clear the screen on Windows
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	} else {
		cmd := exec.Command("clear") // Command to clear the screen on Unix-like systems
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}
