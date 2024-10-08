package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

var (
	running bool
	mu      sync.Mutex
)

func checkInput() {
	// Terminal in den Raw-Modus versetzen
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState) // Stelle die alte Terminal-Konfiguration wieder her

	for {
		// Lese ein Byte
		var b [1]byte
		_, err := os.Stdin.Read(b[:])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		mu.Lock()
		// Überprüfe, ob das eingegebene Zeichen 'q' ist
		if b[0] == 'q' {
			running = false
		}
		mu.Unlock()
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

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
