package terminal

import (
	"github.com/Bios-Marcel/tcell-game-template/state"
	"github.com/gdamore/tcell/v2"
)

// Renderer draws the state.GameSession onto a tcell.Screen. While the Renderer
// can hold state, it's recommended to keep it as state free as possible.
type Renderer struct {
}

// NewRenderer creates a ready to use Renderer.
func NewRenderer() *Renderer {
	return nil
}

// Draw renders the passed state.GameSession onto the passed tcell.Screen.
func (renderer *Renderer) Draw(screen tcell.Screen, session *state.GameSession) {
	//Draw

	screen.Show()
}
