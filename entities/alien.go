package entities

import tl "github.com/JoelOtter/termloop"

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
	}
}

func (alien *Alien) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*Laser); ok {
		*alien.WaitingTime = *alien.WaitingTime - 0.02
		alien.Level.RemoveEntity(alien)
	}
}
