package main

import (
	"fmt"
	"math"

	"github.com/corentings/chess/v2"
)

func search(game *chess.Game, depth int) float64 {
	return 0
}

func getBestMove(game *chess.Game, depth int) chess.Move {
	bestScore := math.Inf(-1)
	var bestMove *chess.Move
	validMoves := game.ValidMoves()

	for _, move := range validMoves {
		// _ indicates unused since we dont care about idx and go forces enumeration

		gameCopy := game.Clone()
		gameCopy.Move(&move, nil)
		
		final_position_score_after_move := search(gameCopy, depth - 1) // we played 1 move so depth -= 1
		// get final elo of playing said move

		if (final_position_score_after_move > bestScore) {
			m := move // copy so we dont affect the iterating variable
			bestMove = &m
			bestScore = final_position_score_after_move
		}
	}

	return *bestMove
}

func main() {
	game := chess.NewGame()
	fmt.Println(game.Position().Board().Draw())
	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		// select a random move
		// moves := game.ValidMoves()
		move := getBestMove(game, 5)
		if err := game.Move(&move, nil); err != nil {
			panic(err) // Should not happen with valid moves
		}
	}
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}
