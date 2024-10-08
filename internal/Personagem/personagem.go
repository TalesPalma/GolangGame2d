package personagem

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Personagem struct {
	X         float64
	Y         float64
	Animation *animation.Animation
	direction string
}

func (p *Personagem) Behavior() error {
	p.Animation.Update()
	p.setAnimation()

	keys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if len(keys) == 0 {
		return nil
	}
	switch keys[0] {
	case ebiten.KeyRight:
		p.X += 10
	case ebiten.KeyLeft:
		p.X -= 10
	}

	return nil
}

func (p *Personagem) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	// Obtenha a imagem atual
	currentFrame := p.Animation.GetCurrentFrame()
	op.GeoM.Translate(p.X, p.Y)

	// Desenhe a imagem no screen
	screen.DrawImage(currentFrame, op)
}

func (p *Personagem) setAnimation() {

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) || inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		if p.direction != "" {
			p.Animation = animation.NewAnimation(config.IdleImagePath, 120, 80, 10)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.movimentPersonage("right")
		p.direction = "right"
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.movimentPersonage("left")
		p.direction = "left"
	}

}

func (p *Personagem) movimentPersonage(direction string) {
	if direction == "right" {
		p.Animation = animation.NewAnimation(config.RunImageRightPath, 120, 80, 10)
	} else if direction == "left" {
		p.Animation = animation.NewAnimation(config.RunImageLeftPath, 120, 80, 10)
	}
}
