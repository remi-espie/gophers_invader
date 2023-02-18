package entities

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Alien struct {
	*tl.Entity
	Game        *tl.Game
	Level       *tl.BaseLevel
	WaitingTime *float64
	TimeDelta   float64
	Direction   bool
	X           int
}

func (alien *Alien) Draw(screen *tl.Screen) {
	alien.Entity.Draw(screen)
}

func (alien *Alien) Tick(e tl.Event) {
	timeDelta := alien.Game.Screen().TimeDelta()
	alien.TimeDelta += timeDelta
	if alien.TimeDelta > *alien.WaitingTime {
		alien.TimeDelta = 0
		x, y := alien.Position()

		if x > alien.X+10 {
			alien.Direction = false
			y++
		}
		if x < alien.X-10 {
			alien.Direction = true
			y++
		}

		if alien.Direction {
			alien.SetPosition(x+1, y)
		} else {
			alien.SetPosition(x-1, y)
		}

		if rand.Intn(200) == 0 {
			alien.Shoot()
		}

	}
}

func (alien *Alien) Shoot() {
	x, y := alien.Position()

	laser := Laser{
		Entity:        tl.NewEntity(x+1, y+1, 1, 1),
		ShootByPlayer: false,
		Level:         alien.Level,
		Game:          alien.Game,
		WaitingTime:   0.2,
		TimeDelta:     0,
	}

	laser.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '|'})
	alien.Level.AddEntity(&laser)
}

func (alien *Alien) Collide(collision tl.Physical) {
	if _, ok := collision.(*Laser); ok {
		laser := collision.(*Laser)
		if laser.ShootByPlayer {
			*alien.WaitingTime = *alien.WaitingTime - 0.02
			alien.Level.RemoveEntity(alien)
		}
	}
}
