package utils

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
