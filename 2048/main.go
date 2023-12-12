package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tomatoSister/2048/core/game"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
