package animation

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	image       *ebiten.Image
	frameWidth  int
	frameHeight int
	numFrames   int
	frameIndex  int
}

func NewAnimation(image *ebiten.Image, frameWidth, frameHeight int, numFrames int) *Animation {
	return &Animation{
		image:       image,
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
	return a.image.SubImage(frameReact).(*ebiten.Image)
}
