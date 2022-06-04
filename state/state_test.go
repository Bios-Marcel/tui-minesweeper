package state

import (
	"sync"
	"testing"
)

func TestGameSession_InterfaceCompliance(t *testing.T) {
	var _ sync.Locker = &GameSession{}
}
