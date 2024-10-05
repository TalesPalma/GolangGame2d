package main

import (
	"log"

	game "github.com/TalesPalma/Game2dGolang/internal/Game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello Tales")
	if err := ebiten.RunGame(game.InitGame()); err != nil {
		log.Fatal(err)
	}
}
