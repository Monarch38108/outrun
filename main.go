package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	running bool
	mu      sync.Mutex
)

func main() {
	running = true
	go checkInput()

	for running {
		mu.Lock()
		fmt.Println("Game is running...")
		mu.Unlock()

		time.Sleep(100 * time.Millisecond)
		clearScreen() // Bildschirm löschen, um das nächste Frame anzuzeigen
	}

	fmt.Println("Game ended.")
}
