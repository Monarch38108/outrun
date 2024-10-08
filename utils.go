package main

import (
	"fmt"
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

// renderFrame renders the framebuffer to the console if it has changed
func renderFrame() {
	// Check for changes and render only if needed
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if framebuffer[y][x] != lastFrame[y][x] {
				os.Stdout.WriteString(fmt.Sprintf("\033[%d;%dH%c", y+1, x+1, framebuffer[y][x]))
				lastFrame[y][x] = framebuffer[y][x] // Update lastFrame
			}
		}
	}
}
