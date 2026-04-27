package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*

 */

type Game struct {
	x float32
	y float32
}

func (g* Game) Update() error {
	return nil
}

func (g* Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, g.x, g.y, 40, 20, color.White, false)
}

func (g* Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// outsideWith, Height provided by engine (parameters)
	// returns int, int
	return 320, 240
}

func main() {
	game := &Game{};
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("My Chess Engine")

	if err := ebiten.RunGame(game); err != nil{
		log.Fatal(err)
	} 

}