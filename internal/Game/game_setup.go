// game_setup.go
package game

import (
	"log"

	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"
	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/TalesPalma/Game2dGolang/internal/utils"
)

func InitGame() *Game {
	image, err := utils.LoadImage(config.IdleImagePath)
	if err != nil {
		log.Fatal("Error ao carregar a imagem do personagem ", err)
	}

	animation := animation.NewAnimation(image, config.FrameWidth, config.FrameHeight, config.NumFrames)
	personagem := &personagem.Personagem{
		X:         0,
		Y:         0,
		Animation: animation,
	}
	return &Game{
		personagem: personagem,
	}
}
