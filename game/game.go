package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"os"
	"fmt"
	"bytes"
	"image/png"
)

// var r image.Rectangle

type Game struct {
	pos int
}

func (g *Game) Init () {
	g.pos = 0;
}

func (g *Game) Layout (int, int) (int, int) {
	return 640, 480
}

func (g *Game) Draw (screen *ebiten.Image) {
    screen.Fill(color.RGBA{255, 255, 255, 255})
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(100.0, float64(g.pos%500))

	data, err := os.ReadFile("game/bukvy.png")
	if err != nil {
		fmt.Println(err)
	}
	
	imageReader := bytes.NewReader(data)
	buk, err := png.Decode(imageReader)
	if err != nil {
		fmt.Println(err)
	}
	
	bukImage := ebiten.NewImageFromImage(buk)
	subImage := bukImage.SubImage(image.Rect(68, 37, 103, 109)).(*ebiten.Image)
	screen.DrawImage(subImage, options)
	g.pos++

}

func (g *Game) Update () error {
    return nil
}