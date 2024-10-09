// game_setup.go
package game

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"
	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/TalesPalma/Game2dGolang/internal/scenario"
)

func InitGame() *Game {
	animation := animation.NewAnimation(config.IdleImagePath, config.FrameWidth, config.FrameHeight, config.NumFrames)
	personagem := &personagem.Personagem{
		X:         100,
		Y:         100,
		Animation: animation,
	}
	scenario := scenario.InitScenario()
	return &Game{
		personagem: personagem,
		scenario:   scenario,
	}
}
