package entities

import tl "github.com/JoelOtter/termloop"

type Button struct {
	*tl.Rectangle
	Level  *tl.BaseLevel
	Text   string
	Action func()
}

func (button *Button) Draw(screen *tl.Screen) {
	x, y := button.Position()
	button.Level.AddEntity(tl.NewText(x+8-len(button.Text)/2, y+1, button.Text, tl.ColorWhite, tl.ColorDefault))
	button.Rectangle.Draw(screen)
}

func (button *Button) Collide(collision tl.Physical) {
	if _, ok := collision.(*Laser); ok {
		collision.(*Laser).Level.RemoveEntity(collision.(*Laser))
		button.Action()
	}
}
