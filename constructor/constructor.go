package constructor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/drevald/nestor/game"
	"fmt"
)

type Game struct {

}

var arr []string

func (g *Game) Draw(screen *ebiten.Image) {
	arr = []string{"аз", "буки", "веди", "глаголь"}
	b := &game.Bukva{}
	fmt.Print(arr)
	fmt.Print(b)
}

func (g *Game) Layout(x, y int) (xx, yy int) {
	return 100, 100
}

func (g *Game) Update () error {
	return nil
}