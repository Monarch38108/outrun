package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	running bool       // Flag to indicate if the game is running
	mu      sync.Mutex // Mutex to synchronize access to shared variables
)

func main() {
	running = true  // Set the game state to running
	go checkInput() // Start the input checking in a separate goroutine

	// Main Gameloop
	for running {
		mu.Lock()
		fmt.Println("Game is running...") // Display the game status (debug)
		mu.Unlock()

		time.Sleep(100 * time.Millisecond) // Control the game update rate
		clearScreen()                      // Clear the Screen to show the next frame
	}

	fmt.Println("Game ended.") // Display end message (debug)
}
