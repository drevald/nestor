package game

import (
	"testing"
	"fmt"
	"os"
	"image/png"
	// _ "image/png"
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestReadExternalImage (t *testing.T) {
	data, err := os.ReadFile("bukvy.png")
	if err != nil {
		fmt.Println(err)
	}
	imageReader := bytes.NewReader(data)
	image, err := png.Decode(imageReader)
	if err != nil {
		fmt.Println(err)
	}
	bubbleImage := ebiten.NewImageFromImage(image)
	screen := ebiten.NewImage(100, 100)
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(bubbleImage, options)
}