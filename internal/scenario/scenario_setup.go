package scenario

import (
	"log"

	"github.com/TalesPalma/Game2dGolang/internal/config"
	"github.com/TalesPalma/Game2dGolang/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scenario struct {
	Background *ebiten.Image
}

func InitScenario() *Scenario {
	image, err := utils.LoadImage(config.BackgroundImage)
	if err != nil {
		log.Fatal("Erro ao carregar imagem de fundo", err)
	}
	return &Scenario{
		Background: image,
	}
}
