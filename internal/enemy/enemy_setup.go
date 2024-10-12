package enemy

import (
	animation "github.com/TalesPalma/Game2dGolang/internal/Animation"
	"github.com/TalesPalma/Game2dGolang/internal/config"
)

type Enemy struct {
	X         float64
	Y         float64
	animation *animation.Animation
}

func NewEnemy(x float64, y float64) *Enemy {

	animation := animation.NewAnimation(config.IdleEnemyPath, config.FrameWidth, config.FrameHeight, config.NumFrames)

	return &Enemy{X: x, Y: y, animation: animation}
}
