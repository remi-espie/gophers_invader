package screen

import (
	_ "embed"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/entities"
	"time"
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
	aliens.Draw()

	level.AddEntity(tl.NewRectangle(-40, -12, 80, 1, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-40, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(40, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-40, 13, 80, 1, tl.ColorWhite))
	level.AddEntity(tl.NewText(25, 15, "Score:", tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(-25, 15, "Time:", tl.ColorGreen, tl.ColorBlack))

	timer := tl.NewFpsText(-25, 16, tl.ColorGreen, tl.ColorBlack, 60)
	level.AddEntity(timer)

	game.Screen().SetLevel(level)
	go Timer(timer)
	game.Start()
}

func Timer(timer *tl.FpsText) {
	deltaTime := 0.0
	for {
		deltaTime += 0.01
		timer.SetText(fmt.Sprintf("%f", deltaTime))
		time.Sleep(time.Millisecond)
	}
}
