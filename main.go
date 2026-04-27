package main

import (
	"fmt"

	"github.com/corentings/chess/v2"
)

func main() {
	game := chess.NewGame()
	fmt.Println(game.Position().Board().Draw())
	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		// select a random move
		moves := game.ValidMoves()
		move := getBestMove(moves, game.Position)
		if err := game.Move(&move, nil); err != nil {
			panic(err) // Should not happen with valid moves
		}
	}
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}
