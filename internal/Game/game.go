package game

import (
	personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"
	"github.com/TalesPalma/Game2dGolang/internal/scenario"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	personagem *personagem.Personagem
	scenario   *scenario.Scenario
}

func (g *Game) Update() error {
	err := g.personagem.Behavior()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.scenario.Background, nil)
	if g.personagem != nil {
		g.personagem.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
