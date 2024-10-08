package main

import (
	"golang.org/x/term"
	"os"
)

// checkInput listens for user input in raw mode
func checkInput() {
	// Set the terminal to raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return // If there's an error, exit the function
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {

		}
	}(int(os.Stdin.Fd()), oldState) // Restore the old terminal settings when done

	for {
		// Read a single byte from standard input
		var b [1]byte
		_, err := os.Stdin.Read(b[:])
		if err != nil {
			return // Exit if there's an error reading input
		}

		mu.Lock()
		// Check if the input character is 'q'
		if b[0] == 'q' {
			running = false // Set running to false to exit the game loop
		}
		mu.Unlock()
	}
}
