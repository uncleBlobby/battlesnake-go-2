package main

import "fmt"

func ChooseMove(scoredMoves map[string]int, viewDistances map[string]int) string {
	var nextMove string

	scoredMoves["left"] += viewDistances["left"]
	scoredMoves["right"] += viewDistances["right"]
	scoredMoves["up"] += viewDistances["up"]
	scoredMoves["down"] += viewDistances["down"]

	fmt.Println(scoredMoves)

	// Default move to the right.
	nextMove = "right"
	temporaryMaximum := scoredMoves["right"]

	for move, score := range scoredMoves {
		if score > temporaryMaximum {
			temporaryMaximum = score
			nextMove = move
		}
	}
	fmt.Printf("Chose: %v Score: %v.\n", nextMove, temporaryMaximum)
	return nextMove
}
