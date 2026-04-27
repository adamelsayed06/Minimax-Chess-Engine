package main

import(
	"github.com/hajimehoshi/ebiten/v2"
)

/*

*/

type Game struct {}

func (g* Game) Update() error {
	return nil
}

func (g* Game) Draw(screen *ebiten.Image) error {
	return nil;
}

func (g* Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// outsideWith, Height provided by engine (parameters)
	// returns int, int
	return 320, 460
}

func main() {

}