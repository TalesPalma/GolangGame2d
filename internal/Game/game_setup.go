// game_setup.go
package game

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"
	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/TalesPalma/Game2dGolang/internal/enemy"
	"github.com/TalesPalma/Game2dGolang/internal/scenario"
)

func InitGame() *Game {
	animation := animation.NewAnimation(config.IdleImagePath, config.FrameWidth, config.FrameHeight, config.NumFrames)
	personagem := personagem.NewPersonagem(animation)
	scenario := scenario.InitScenario()
	enemys := []enemy.Enemy{
		*enemy.NewEnemy(100),
		*enemy.NewEnemy(100),
	}
	return &Game{
		personagem: personagem,
		scenario:   scenario,
		enemy:      enemys,
	}
}
