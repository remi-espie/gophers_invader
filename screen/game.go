package screen

import (
	_ "embed"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"gophers_invader/entities"
	"strconv"
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

func NewGame(game *tl.Game, mainMenuLevel *tl.BaseLevel) {

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	playerGameOver := false

	player := entities.Player{
		Entity:   tl.NewEntityFromCanvas(0, 10, entities.CreateCanvas(playerBytes)),
		Level:    level,
		Game:     game,
		GameOver: &playerGameOver,
	}
	level.AddEntity(&player)

	waintingTime := 1.0

	aliens := entities.AlienCluster{
		Level:           level,
		Game:            game,
		AlienMiddleByte: alienMiddleBytes,
		AlienLowByte:    alienLowBytes,
		AlienHighByte:   alienHighBytes,
		WaitingTime:     &waintingTime,
	}
	aliens.CreateCluster()

	gameOverZone := entities.GameOverZone{
		Entity:      tl.NewEntity(-40, 0, 80, 1),
		EnteredZone: false,
	}
	level.AddEntity(&gameOverZone)

	level.AddEntity(tl.NewRectangle(-40, -12, 80, 1, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-40, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(40, -12, 1, 25, tl.ColorWhite))
	level.AddEntity(tl.NewRectangle(-40, 13, 80, 1, tl.ColorWhite))
	level.AddEntity(tl.NewText(25, 15, "Score:", tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(-25, 15, "Time:", tl.ColorGreen, tl.ColorBlack))

	timer := tl.NewFpsText(-25, 16, tl.ColorGreen, tl.ColorBlack, 60)
	level.AddEntity(timer)
	scoreText := tl.NewFpsText(25, 16, tl.ColorGreen, tl.ColorBlack, 60)
	level.AddEntity(scoreText)

	game.Screen().SetLevel(level)
	go Loop(timer, scoreText, &waintingTime, level, gameOverZone, game, mainMenuLevel, &playerGameOver)
	//game.Start()
}

func Loop(timer *tl.FpsText, scoreText *tl.FpsText, waintingTime *float64, level *tl.BaseLevel, gameOverZone entities.GameOverZone, game *tl.Game, mainMenuLevel *tl.BaseLevel, playerGameOver *bool) {
	deltaTime := float32(0.0)
	for {
		deltaTime += 0.001
		timer.SetText(fmt.Sprintf("%.4f", deltaTime))

		score := 55
		for _, entity := range level.Entities {
			if _, ok := entity.(*entities.Alien); ok {
				score -= 1
			}
		}

		bonus := int((1.0 / *waintingTime) * 100)
		scoreText.SetText(strconv.Itoa(score * bonus))

		if gameOverZone.EnteredZone || *playerGameOver {
			GameOver(game, score*bonus, deltaTime, mainMenuLevel)
			break
		}

		time.Sleep(time.Millisecond)
	}
}