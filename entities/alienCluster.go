package entities

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
	"time"
)

type AlienCluster struct {
	*tl.Entity
	Game            *tl.Game
	Level           *tl.BaseLevel
	AlienLowByte    []byte
	AlienMiddleByte []byte
	AlienHighByte   []byte
	WaitingTime     float64
	alienArray      []Alien
	scoreText       *tl.FpsText
	timeText        *tl.FpsText
}

func (alienCluster *AlienCluster) Draw() {
	alienCluster.scoreText = tl.NewFpsText(25, 16, tl.ColorRed, tl.ColorBlack, 60)
	alienCluster.scoreText.SetText("0")
	alienCluster.Level.AddEntity(alienCluster.scoreText)

	alienCluster.timeText = tl.NewFpsText(-25, 16, tl.ColorGreen, tl.ColorBlack, 60)
	alienCluster.Level.AddEntity(alienCluster.timeText)

	alienCluster.CreateCluster()
	go alienCluster.Loop()

}

func (alienCluster *AlienCluster) Loop() {
	for {
		score := 55
		for _, entity := range alienCluster.Level.Entities {
			if _, ok := entity.(*Alien); ok {
				score -= 1
			}
		}

		bonus := int((1.0 / alienCluster.WaitingTime) * 100)

		alienCluster.scoreText.SetText(strconv.Itoa(score * bonus))
		time.Sleep(60 * time.Millisecond)
	}
}

func (alienCluster *AlienCluster) CreateCluster() {
	alienCluster.AddAliens(alienCluster.AlienHighByte, 5, &alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienMiddleByte, 4, &alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienMiddleByte, 3, &alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienLowByte, 2, &alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienLowByte, 1, &alienCluster.WaitingTime)
}

func (alienCluster *AlienCluster) AddAliens(alien []byte, lign int, waitingTime *float64) {
	for j := -22; j < 22; j = j + 4 {
		alien := Alien{
			Entity:      tl.NewEntityFromCanvas(j, -lign*2, CreateCanvas(alien)),
			Level:       alienCluster.Level,
			Game:        alienCluster.Game,
			WaitingTime: waitingTime,
			Direction:   true,
			X:           j,
		}
		alienCluster.Level.AddEntity(&alien)
	}
}
