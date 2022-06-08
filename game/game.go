package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Init () {

}

func (g *Game) Layout (int, int) (int, int) {
	return 100, 100
}

func (g *Game) Draw (*ebiten.Image) {

}

func (g *Game) Update () error {
	return nil
}

type Game struct {}