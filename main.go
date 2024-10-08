package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	width  = 40 // Width of the screen
	height = 20 // Height of the screen
)

var (
	running     bool
	mu          sync.Mutex
	framebuffer [height][width]rune // Framebuffer to hold the current state of the screen
	hasChanged  bool                // Flag to check if there are any changes
	xPos        int                 // Current x position of 'X'
	yPos        int                 // Current y position of 'X'
)

func main() {
	running = true
	xPos = 5 // Start position of X
	yPos = 5 // Start position of Y

	go checkInput() // Start the input checking in a separate goroutine

	// Main game loop
	for running {
		mu.Lock()
		updateFrame()   // Update the framebuffer with the current game state
		if hasChanged { // Only render if there are changes
			renderFrame()      // Render the framebuffer to the console
			hasChanged = false // Reset the change flag
		}
		mu.Unlock()

		time.Sleep(100 * time.Millisecond) // Control the game update rate
	}

	fmt.Println("Game ended.")
}
