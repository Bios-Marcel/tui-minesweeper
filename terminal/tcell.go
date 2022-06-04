package terminal

import "github.com/gdamore/tcell/v2"

// CreateScreen generates a ready to use screen. The screen has
// no cursor and doesn't support mouse eventing.
func CreateScreen() (tcell.Screen, error) {
	screen, screenCreationError := tcell.NewScreen()
	if screenCreationError != nil {
		return nil, screenCreationError
	}

	screenInitError := screen.Init()
	if screenInitError != nil {
		return nil, screenInitError
	}

	//Make sure it's disable, even though it should be by default.
	screen.DisableMouse()
	//Make sure cursor is hidden by default.
	screen.HideCursor()

	return screen, nil
}

// DrawFilledRectangle draws a rectangle with the specified style. This
// removes all runes previously written into the target area of the rectangle.
func DrawFilledRectangle(screen tcell.Screen, xStart, yStart, width, height int, style tcell.Style) {
	for y := yStart; y < yStart+height; y++ {
		for x := xStart; x < xStart+width; x++ {
			screen.SetContent(x, y, 0, nil, style)
		}
	}
}
