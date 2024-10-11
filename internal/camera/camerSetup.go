package camera

type Camera struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewCamera(x, y, width, height float64) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
