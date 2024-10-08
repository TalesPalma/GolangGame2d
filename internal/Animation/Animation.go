package animation

import (
	"image"

	"github.com/TalesPalma/Game2dGolang/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Image       *ebiten.Image
	frameWidth  int
	frameHeight int
	numFrames   int
	frameIndex  int
}

func NewAnimation(imagePath string, frameWidth, frameHeight int, numFrames int) *Animation {
	image, err := utils.LoadImage(imagePath)
	if err != nil {
		panic(err)
	}

	return &Animation{
		Image:       image,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		numFrames:   numFrames,
	}
}

func (a *Animation) Update() {
	a.frameIndex++
	if a.frameIndex >= a.numFrames {
		a.frameIndex = 0
	}
}

func (a *Animation) GetCurrentFrame() *ebiten.Image {
	frameX := a.frameIndex * a.frameWidth
	frameReact := image.Rect(frameX, 0, frameX+a.frameWidth, a.frameHeight)
	return a.Image.SubImage(frameReact).(*ebiten.Image)
}

// func (a *Animation) FlipImage() {
// 	flippedImage := ebiten.NewImage(a.frameWidth, a.frameHeight)
//
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Scale(1, -1)
//
// 	a.Image = flippedImage
// }
