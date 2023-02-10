package screen

import (
	_ "embed"
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/entities"
)

//go:embed canvas/player.txt
var playerBytes []byte

//go:embed canvas/aliens/alien_low.txt
var alienLowBytes []byte

//go:embed canvas/aliens/alien_middle.txt
var alienMiddleBytes []byte

//go:embed canvas/aliens/alien_high.txt
var alienHighBytes []byte

func NewGame() {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	player := entities.Player{
		Entity: tl.NewEntityFromCanvas(0, 10, entities.CreateCanvas(playerBytes)),
		Level:  level,
		Game:   game,
	}
	level.AddEntity(&player)

	aliens := entities.AlienCluster{
		Level:           level,
		Game:            game,
		AlienMiddleByte: alienMiddleBytes,
		AlienLowByte:    alienLowBytes,
		AlienHighByte:   alienHighBytes,
		WaitingTime:     1.0,
	}
	aliens.CreateCluster()

	level.AddEntity(tl.NewRectangle(-40, -12, 80, 1, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(-40, -12, 1, 30, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(40, -12, 1, 30, tl.ColorBlue))
	game.Screen().SetLevel(level)
	game.Start()
}
