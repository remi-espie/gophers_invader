package screen

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

type GameOverEntity struct {
	*tl.Entity
	Game     *tl.Game
	Name     string
	Score    int
	Duration float32
	NameBox  *tl.FpsText
	MainMenu *tl.BaseLevel
}

func GameOver(game *tl.Game, score int, duration float32, mainMenuLevel *tl.BaseLevel) {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	level.AddEntity(tl.NewText(10, 5, "Game Over", tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(10, 7, "Score: "+strconv.Itoa(score), tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(10, 8, "Time: "+fmt.Sprintf("%.2f", duration)+"s", tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(10, 10, "Enter your name and press ENTER to save score", tl.ColorRed, tl.ColorBlack))
	level.AddEntity(tl.NewText(10, 15, "Name:", tl.ColorRed, tl.ColorBlack))

	nameBox := tl.NewFpsText(15, 15, tl.ColorWhite, tl.ColorBlack, 60)
	nameBox.SetText("_____")
	name := ""
	level.AddEntity(nameBox)
	GameOverEntity := GameOverEntity{
		Entity:   tl.NewEntity(1, 1, 1, 1),
		Game:     game,
		Name:     name,
		Score:    score,
		Duration: duration,
		NameBox:  nameBox,
		MainMenu: mainMenuLevel,
	}

	level.AddEntity(&GameOverEntity)

	game.Screen().SetLevel(level)
	//game.Start()
}

func (entity *GameOverEntity) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		if event.Key == tl.KeyEnter {
			AddScore(entity.Name, entity.Score, entity.Duration)
			entity.Game.Screen().SetLevel(entity.MainMenu)
		}
		entity.Name = entity.Name + string(event.Ch)
		entity.NameBox.SetText(entity.Name)
		return
	}
}
