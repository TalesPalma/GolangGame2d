package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	personagem *Personagem
}

type Personagem struct {
	X      float64
	Y      float64
	assets *ebiten.Image
}

func (g *Game) Update() error {

	keys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if len(keys) == 0 {
		return nil
	}

	switch keys[0] {
	case ebiten.KeyRight:
		newImage, err := loadImage("assets/personagem02.png")
		if err != nil {
			log.Fatal(err)
		}
		g.personagem.X += 10
		g.personagem.assets = newImage
	case ebiten.KeyLeft:
		newImage, err := loadImage("assets/personagem03.png")
		if err != nil {
			log.Fatal(err)
		}
		g.personagem.X -= 10
		g.personagem.assets = newImage
	default:
		newImage, err := loadImage("assets/personagem01.png")
		if err != nil {
			log.Fatal(err)
		}
		g.personagem.assets = newImage
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if g.personagem.assets != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.personagem.X, g.personagem.Y)
		screen.DrawImage(g.personagem.assets, op)
	}

	ebitenutil.DebugPrint(screen, "Hello world")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello Tales")

	image, err := loadImage("assets/personagem01.png")
	if err != nil {
		log.Fatal("Error ao carregar a imagem do personagem ", err)
	}
	game := &Game{
		personagem: &Personagem{
			X:      0,
			Y:      0,
			assets: image,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func loadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
