package entities

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	Game     *tl.Game
	prevX    int
	Level    *tl.BaseLevel
	GameOver *bool
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	player.Level.SetOffset(screenWidth/2, screenHeight/2)
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
		*player.GameOver = true
	}
	if _, ok := collision.(*Laser); ok {
		*player.GameOver = true
	}
}

func (player *Player) Shoot() {

	entities := player.Level.Entities

	for _, entity := range entities {
		if _, ok := entity.(*Laser); ok {
			laser := entity.(*Laser)
			if laser.ShootByPlayer {
				return
			}
		}
	}

	x, y := player.Position()
	laser := Laser{
		Entity:        tl.NewEntity(x+1, y-1, 1, 1),
		ShootByPlayer: true,
		Level:         player.Level,
		Game:          player.Game,
		WaitingTime:   0.1,
		TimeDelta:     0,
	}

	laser.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '|'})
	player.Level.AddEntity(&laser)
}
