package enemy

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Enemy) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(e.X, e.Y)

	currentFrame := e.animation.GetCurrentFrame()

	screen.DrawImage(currentFrame, op)
}
