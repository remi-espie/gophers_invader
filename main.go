package main

import (
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/screen"
)

func main() {

	game := tl.NewGame()
	game.Screen().SetFps(60)

	screen.MainMenu(game)
}
