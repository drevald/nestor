package main

import (
	"github.com/drevald/nestor/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main () {
	g := &game.Game{}
	g.Init()
	ebiten.RunGame(g)	
}