package main

import (
	"fmt"
	"os"
)

// renderFrame updates only the changed characters on the console
func renderFrame() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Move the cursor to (x, y) position and print the character
			os.Stdout.WriteString(fmt.Sprintf("\033[%d;%dH%c", y+1, x+1, framebuffer[y][x]))
		}
	}
}
