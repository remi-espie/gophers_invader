package entities

import tl "github.com/JoelOtter/termloop"

type GameOverZone struct {
	*tl.Entity
	EnteredZone bool
}

func (gameOverZone *GameOverZone) Collide(collision tl.Physical) {
	if _, ok := collision.(*Alien); ok {
		gameOverZone.EnteredZone = true
	}
}
