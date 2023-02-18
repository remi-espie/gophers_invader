package screen

import (
	_ "embed"
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/entities"
	"os"
)

//go:embed canvas/art.txt
var titleArt []byte

func MainMenu(game *tl.Game) {

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	buttonPlay := entities.Button{
		Rectangle: tl.NewRectangle(-8, 4, 16, 3, tl.ColorGreen),
		Level:     level,
		Text:      "New Game",
		Action: func() {
			NewGame(game, level)
		},
	}

	buttonExit := entities.Button{
		Rectangle: tl.NewRectangle(16, 4, 16, 3, tl.ColorRed),
		Level:     level,
		Text:      "Exit",
		Action: func() {
			os.Exit(0)
		},
	}

	buttonScoreboard := entities.Button{
		Rectangle: tl.NewRectangle(-32, 4, 16, 3, tl.ColorYellow),
		Level:     level,
		Text:      "Scoreboard",
		Action: func() {
			Scoreboard(game, level)
		},
	}

	level.AddEntity(tl.NewEntityFromCanvas(-80, -10, entities.CreateCanvas(titleArt)))

	level.AddEntity(tl.NewRectangle(-100, -12, 200, 1, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-100, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(100, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-100, 13, 200, 1, tl.ColorWhite))

	level.AddEntity(tl.NewText(-25, 0, "Press ← or → to move !", tl.ColorBlue, tl.ColorBlack))
	level.AddEntity(tl.NewText(25, 0, "Press ⎵ to shoot !", tl.ColorBlue, tl.ColorBlack))

	level.AddEntity(&buttonPlay)

	level.AddEntity(&buttonScoreboard)

	level.AddEntity(&buttonExit)

	player := entities.Player{
		Entity: tl.NewEntityFromCanvas(0, 10, entities.CreateCanvas(playerBytes)),
		Level:  level,
		Game:   game,
	}
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	game.Start()
}
