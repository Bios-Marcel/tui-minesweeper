package state

import (
	"sync"
)

// GameSession holds a renderable game state.
type GameSession struct {
	mutex                     *sync.Mutex
	renderNotificationChannel chan bool
}

// Lock implements sync.Locker.
func (session *GameSession) Lock() {
	session.mutex.Lock()
}

// Unlock implements sync.Locker.
func (session *GameSession) Unlock() {
	session.mutex.Unlock()
}

// NewGameSession produces a ready-to-use session state. The
// renderNotificationChannel is used to notify the renderer that the game
// state has changed and a rerender is required. This prevents having to
// constantly rerender without changes.
func NewGameSession(renderNotificationChannel chan bool) *GameSession {
	session := &GameSession{
		mutex:                     &sync.Mutex{},
		renderNotificationChannel: renderNotificationChannel,
	}
	return session
}
