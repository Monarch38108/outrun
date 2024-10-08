package main

import (
	"golang.org/x/term"
	"os"
)

func checkInput() {
	// Terminal in den Raw-Modus versetzen
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState) // Stelle die alte Terminal-Konfiguration wieder her

	for {
		// Lese ein Byte
		var b [1]byte
		_, err := os.Stdin.Read(b[:])
		if err != nil {
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
