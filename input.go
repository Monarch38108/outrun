package main

import (
	"fmt"
)

// checkInput handles keyboard input
func checkInput() {
	var input string
	for running {
		fmt.Scan(&input) // Read input from the console
		switch input {
		case "q":
			running = false // Quit the game
		case "a":
			moveX(-1) // Move left on 'a'
		case "d":
			moveX(1) // Move right on 'd'
		}
	}
}

// moveX updates the position of 'X' based on input
func moveX(direction int) {
	mu.Lock()
	// Update xPos based on direction
	if direction == -1 && xPos > 0 { // Move left
		xPos--
		hasChanged = true
	} else if direction == 1 && xPos < width-1 { // Move right
		xPos++
		hasChanged = true
	}
	mu.Unlock()
}
