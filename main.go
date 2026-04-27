package main

import (
	"fmt"
	"math"

	"github.com/corentings/chess/v2"
)

func getBestMove(game *chess.Game) chess.Move {
	bestScore := math.Inf(-1)
	validMoves := game.ValidMoves()


}

func main() {
	game := chess.NewGame()
	fmt.Println(game.Position().Board().Draw())
	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		// select a random move
		// moves := game.ValidMoves()
		move := getBestMove(game)
		if err := game.Move(&move, nil); err != nil {
			panic(err) // Should not happen with valid moves
		}
	}
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}
