package entities

import (
	tl "github.com/JoelOtter/termloop"
)

type AlienCluster struct {
	Game            *tl.Game
	Level           *tl.BaseLevel
	AlienLowByte    []byte
	AlienMiddleByte []byte
	AlienHighByte   []byte
	WaitingTime     *float64
	alienArray      []Alien
}

func (alienCluster *AlienCluster) CreateCluster() {
	alienCluster.AddAliens(alienCluster.AlienHighByte, 5, alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienMiddleByte, 4, alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienMiddleByte, 3, alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienLowByte, 2, alienCluster.WaitingTime)
	alienCluster.AddAliens(alienCluster.AlienLowByte, 1, alienCluster.WaitingTime)
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
