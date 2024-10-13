package enemy

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Enemy) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(e.X, e.Y)

	currentFrame := e.animation.GetCurrentFrame()

	screen.DrawImage(currentFrame, op)
}

func (e *Enemy) Behavior() error {
	return nil
}

func (e *Enemy) GetPosition() (float64, float64) {
	return e.X, e.Y
}

func (e *Enemy) SetPosition(x, y float64) {
	e.X = x
	e.Y = y
}

func (e *Enemy) SetAnimation(animation *animation.Animation) {
	e.animation = animation
}
