package main

import (
	"log"

	game "github.com/TalesPalma/Game2dGolang/internal/Game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tales Of Palma")
	if err := ebiten.RunGame(game.InitGame()); err != nil {
		log.Fatal(err)
	}
}
