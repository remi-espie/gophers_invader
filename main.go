package main

import (
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/entities"
)

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	player := entities.Player{
		Entity: tl.NewEntity(0, 0, 3, 2),
		Level:  level,
		Game:   game,
	}

	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '█'})
	level.AddEntity(&player)

	level.AddEntity(tl.NewRectangle(10, -30, 50, 20, tl.ColorBlue))
	game.Screen().SetLevel(level)
	game.Start()
}
