package entities

import (
	tl "github.com/JoelOtter/termloop"
	"os"
)

type Player struct {
	*tl.Entity
	Game  *tl.Game
	prevX int
	Level *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.prevX = x
			player.SetPosition(x+1, y)
		case tl.KeyArrowLeft:
			player.prevX = x
			player.SetPosition(x-1, y)
		case tl.KeySpace:
			player.Shoot()
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, +10)
	}
	if _, ok := collision.(*Alien); ok {
		os.Exit(0)
	}
	if _, ok := collision.(*Laser); ok {
		os.Exit(0)
	}
}

func (player *Player) Shoot() {

	entities := player.Level.Entities

	for _, entity := range entities {
		if _, ok := entity.(*Laser); ok {
			return
		}
	}

	x, y := player.Position()
	laser := Laser{
		Entity:      tl.NewEntity(x+1, y-1, 1, 1),
		Level:       player.Level,
		Game:        player.Game,
		WaitingTime: 0.1,
		TimeDelta:   0,
	}

	laser.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '|'})
	player.Level.AddEntity(&laser)
}
