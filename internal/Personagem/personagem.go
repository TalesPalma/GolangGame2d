package personagem

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Direction int

const (
	Right Direction = iota
	Left
)

type Personagem struct {
	X         float64
	Y         float64
	Animation *animation.Animation
	Attacked  bool
}

func NewPersonagem(animation *animation.Animation) *Personagem {
	return &Personagem{
		X:         100,
		Y:         100,
		Animation: animation,
		Attacked:  false,
	}
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

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		p.Attack()
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

	if p.Attacked {
		p.Attacked = false
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) || inpututil.IsKeyJustReleased(ebiten.KeyLeft) || inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		p.Animation = animation.NewAnimation(config.IdleImagePath, 120, 80, 10)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.movimentPersonage(Right)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.movimentPersonage(Left)
	}

}

func (p *Personagem) movimentPersonage(direction Direction) {
	if direction == Right {
		p.Animation = animation.NewAnimation(config.RunImageRightPath, 120, 80, 10)
	} else if direction == Left {
		p.Animation = animation.NewAnimation(config.RunImageLeftPath, 120, 80, 10)
	}
}

func (p *Personagem) Attack() {
	if !p.Attacked {
		p.Animation = animation.NewAnimation(config.AttackImage, 120, 80, 10)
		p.Attacked = true // Define que o ataque está em andamento
	}
}
