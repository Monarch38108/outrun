package main

// updateFrame updates the framebuffer with the current game state
func updateFrame() {
	hasChanged = false // Reset change flag
	// Clear the framebuffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			framebuffer[y][x] = ' '
		}
	}

	// Set the character at the current position
	framebuffer[yPos][xPos] = 'X'
	hasChanged = true // Mark as changed
}
