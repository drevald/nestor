package main

import (
	// "github.com/drevald/nestor/game"
    "github.com/drevald/nestor/constructor"
	"github.com/hajimehoshi/ebiten/v2"
)

func main () {
    // game := &game.Game{}
    // // Specify the window size as you like. Here, a doubled size is specified.
    // ebiten.SetWindowSize(640, 480)
    // ebiten.SetWindowTitle("Your game's title")
    // // Call ebiten.RunGame to start your game loop.
    // if err := ebiten.RunGame(game); err != nil {
    //     log.Fatal(err)
    // }

//    ebiten.RunGame(&game.Game{})
    ebiten.RunGame(&constructor.Game{})

}