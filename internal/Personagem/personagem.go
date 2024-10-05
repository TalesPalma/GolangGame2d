package personagem

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Personagem struct {
	X         float64
	Y         float64
	Animation *animation.Animation
}

func (p *Personagem) Moviment() error {
	p.Animation.Update()
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if len(keys) == 0 {
		return nil
	}
	switch keys[0] {
	case ebiten.KeyRight:
		p.X += 10
	case ebiten.KeyLeft:
		p.X -= 10
	default:
	}

	return nil
}

func (p *Personagem) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Animation.GetCurrentFrame(), op)
}
