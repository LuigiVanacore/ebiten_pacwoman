package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_pacwoman"
	"github.com/hajimehoshi/ebiten/v2"
)




func main() {
	ebiten.SetWindowSize(ebiten_pacwoman.SCREEN_WIDTH, ebiten_pacwoman.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Sprite Example")

	
	
	game := ebiten_pacwoman.NewGame(false)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}