package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	width  = 40 // Width of the Screen
	height = 20 // Height if the Screen
)

var (
	running     bool                // Flag to indicate if the game is running
	mu          sync.Mutex          // Mutex to synchronize access to shared variables
	framebuffer [height][width]rune // Framebuffer to hold the current state of the screen
	lastFrame   [height][width]rune // Last rendered frame for comparison
)

func main() {
	running = true  // Set the game state to running
	go checkInput() // Start the input checking in a separate goroutine

	// Main Gameloop
	for running {
		mu.Lock()
		updateFrame()
		renderFrame()
		mu.Unlock()

		time.Sleep(100 * time.Millisecond) // Control the game update rate
		clearScreen()                      // Clear the Screen to show the next frame
	}

	fmt.Println("Game ended.") // Display end message (debug)
}
