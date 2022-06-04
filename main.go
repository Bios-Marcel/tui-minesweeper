package main

import (
	"os"

	"github.com/Bios-Marcel/tcell-game-template/state"
	"github.com/Bios-Marcel/tcell-game-template/terminal"
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, screenCreationError := terminal.CreateScreen()
	if screenCreationError != nil {
		panic(screenCreationError)
	}

	//Cleans up the terminal buffer and returns it to the shell.
	defer screen.Fini()

	//renderer used for drawing the board and the menu.
	renderer := terminal.NewRenderer()

	renderNotificationChannel := make(chan bool)
	gameSession := state.NewGameSession(renderNotificationChannel)

	//Gameloop; We draw whenever there's a frame-change. This means we
	//don't have any specific frame-rates and it could technically happen
	//that we don't draw for a while. The first frame is drawn without
	//waiting for a change, so that the screen doesn't stay empty.

	//Listen for key input.
	go func() {
		for {
			switch event := screen.PollEvent().(type) {
			case *tcell.EventKey:
				//Allow closing the game cleanly.
				if event.Key() == tcell.KeyCtrlC {
					screen.Fini()
					os.Exit(0)
				} else if true {
					//Here you can check for all your supports keybindings and
					//accordingly manipulate the gameSession object.

					gameSession.Lock()

					//Invoke public methods from gameSession that might cause state changes.

					gameSession.Unlock()
				}
			case *tcell.EventResize:
				//If the renderer supports resizing, or different sizes in
				//general, we invoke the renderer here to make sure the current
				//screens size is used as good as we can.
				gameSession.Lock()
				screen.Clear()
				gameSession.Unlock()
				renderNotificationChannel <- true
			default:
				//Unsupported or irrelevant event
			}
		}
	}()

	for {
		//We lock before rendering in order to avoid drawing a state that's currently under manipulation.
		gameSession.Lock()
		renderer.Draw(screen, gameSession)
		gameSession.Unlock()

		<-renderNotificationChannel
	}
}
