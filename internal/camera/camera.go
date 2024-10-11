package camera

import personagem "github.com/TalesPalma/Game2dGolang/internal/Personagem"

func (c *Camera) CameraBehavior(player *personagem.Personagem) {
	c.X = player.X - (c.Width / 2)
	c.Y = player.Y - (c.Height / 2)
}
