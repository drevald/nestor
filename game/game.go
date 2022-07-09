package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"image/color"
	"os"
	"fmt"
	"bytes"
	"image/png"
	"math/rand"
)

// var r image.Rectangle

type Game struct {
	pos int
	bukva Bukva
}

type Bukva struct {
	bounds []int
	char []rune
	offsetX	int
	offsetY	int
}

var bukvy = []Bukva{
	{bounds:[]int{68, 37, 103, 109}, char:[]rune{'\u0430'}}, // азъ
	{bounds:[]int{107, 41, 129, 96}, char:[]rune{'\u0430'}}, // азъ
	{bounds:[]int{58, 177, 119, 234}, char:[]rune{'\u0431'}}, // буки
}

func (g *Game) Init () {
	g.pos = 0;
}

func (g *Game) Layout (int, int) (int, int) {
	return 640, 480
}

func (g *Game) Draw (screen *ebiten.Image) {
	if g.bukva.bounds == nil {
		g.bukva = bukvy[rand.Intn(3)]
		width, _ := ebiten.WindowSize()
		g.bukva.offsetX = rand.Intn(width-50)
		g.bukva.offsetY = 0
	}
    screen.Fill(color.RGBA{255, 255, 255, 255})
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.bukva.offsetX), float64(g.bukva.offsetY%500))
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
	subImage := bukImage.SubImage(image.Rect(g.bukva.bounds[0], g.bukva.bounds[1], g.bukva.bounds[2], g.bukva.bounds[3])).(*ebiten.Image)
	screen.DrawImage(subImage, options)
	g.bukva.offsetY++

}

func (g *Game) Update () error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) && g.bukva.char[0] == '\u0430' {
		g.bukva.bounds = nil
		g.bukva.char = nil
	}	
	if inpututil.IsKeyJustPressed(ebiten.KeyComma) && g.bukva.char[0] == '\u0431' {
		g.bukva.bounds = nil
		g.bukva.char = nil
	}	
    return nil
}