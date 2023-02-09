package entities

import tl "github.com/JoelOtter/termloop"

type Laser struct {
	*tl.Entity
	Game        *tl.Game
	Level       *tl.BaseLevel
	WaitingTime float64
	TimeDelta   float64
}

func (laser *Laser) Draw(screen *tl.Screen) {
	laser.Entity.Draw(screen)
}

func (laser *Laser) Tick(e tl.Event) {
	timeDelta := laser.Game.Screen().TimeDelta()
	laser.TimeDelta += timeDelta
	if laser.TimeDelta > laser.WaitingTime {
		laser.TimeDelta = 0
		x, y := laser.Position()
		laser.SetPosition(x, y-1)
	}
}

func (laser *Laser) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		laser.Level.RemoveEntity(laser)
	}
}
