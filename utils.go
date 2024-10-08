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
		cmd.Run()
	} else {
		cmd := exec.Command("clear") // Command to clear the screen on Unix-like systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// renderFrame renders the entire framebuffer to the console
func renderFrame() {
	// Clear the screen first
	clearScreen()

	// Render the entire framebuffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Print the character at (x, y)
			os.Stdout.WriteString(fmt.Sprintf("\033[%d;%dH%c", y+1, x+1, framebuffer[y][x]))
		}
	}
}

// updateFrame updates the framebuffer with the current game state
func updateFrame() {
	// Example: Fill the framebuffer with a character (for testing)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			framebuffer[y][x] = ' ' // Clear the framebuffer
			// game logic
			if x == 5 && y == 5 { // Example: set a character at (5, 5)
				framebuffer[y][x] = 'X'
			}
		}
	}
}
