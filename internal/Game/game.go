package game

import (
	"image/color"

	personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	personagem *personagem.Personagem
}

func (g *Game) Update() error {
	err := g.personagem.Behavior()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if g.personagem != nil {
		g.personagem.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, "Hello world")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
