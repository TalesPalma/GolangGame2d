package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

func LoadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
