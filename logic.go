package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"fmt"
	"log"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "uncleBlobby", // TODO: Your Battlesnake username
		Color:      "#888888",     // TODO: Personalize
		Head:       "default",     // TODO: Personalize
		Tail:       "default",     // TODO: Personalize
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	scoredMoves := map[string]int{
		"up":    0,
		"down":  0,
		"left":  0,
		"right": 0,
	}
	/*
		possibleMoves := map[string]bool{
			"up":    true,
			"down":  true,
			"left":  true,
			"right": true,
		}
	*/

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
	if myNeck.X < myHead.X {
		scoredMoves["left"] += -100
	} else if myNeck.X > myHead.X {
		scoredMoves["right"] += -100
	} else if myNeck.Y < myHead.Y {
		scoredMoves["down"] += -100
	} else if myNeck.Y > myHead.Y {
		scoredMoves["up"] += -100
	}

	// TODO: Step 1 - Don't hit walls. // only in non-wrapped mode
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if state.Game.Ruleset.Name != "wrapped" {
		if myHead.X == boardWidth-1 {
			scoredMoves["right"] += -100
		}
		if myHead.X == 0 {
			scoredMoves["left"] += -100
		}
		if myHead.Y == boardHeight-1 {
			scoredMoves["up"] += -100
		}
		if myHead.Y == 0 {
			scoredMoves["down"] += -100
		}
	}

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	mybody := state.You.Body

	for i := 0; i < len(mybody); i++ {
		if (myHead.X+1 == mybody[i].X) && (myHead.Y == mybody[i].Y) {
			scoredMoves["right"] += -100
		}
		if (myHead.X-1 == mybody[i].X) && (myHead.Y == mybody[i].Y) {
			scoredMoves["left"] += -100
		}
		if (myHead.X == mybody[i].X) && (myHead.Y+1 == mybody[i].Y) {
			scoredMoves["up"] += -100
		}
		if (myHead.X == mybody[i].X) && (myHead.Y-1 == mybody[i].Y) {
			scoredMoves["down"] += -100
		}
	}

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.
	snakes := state.Board.Snakes

	for i := 0; i < len(snakes); i++ {
		for j := 0; j < int(snakes[i].Length); j++ {
			if (myHead.X-1 == snakes[i].Body[j].X) && (myHead.Y == snakes[i].Body[j].Y) {
				scoredMoves["left"] += -100
			}
			if (myHead.X+1 == snakes[i].Body[j].X) && (myHead.Y == snakes[i].Body[j].Y) {
				scoredMoves["right"] += -100
			}
			if (myHead.X == snakes[i].Body[j].X) && (myHead.Y-1 == snakes[i].Body[j].Y) {
				scoredMoves["down"] += -100
			}
			if (myHead.X == snakes[i].Body[j].X) && (myHead.Y+1 == snakes[i].Body[j].Y) {
				scoredMoves["up"] += -100
			}
		}
	}

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	// Default move will be right

	nextMove = "right"
	temporaryMaximum := scoredMoves["right"]
	fmt.Println(state.Game.Ruleset.Name)
	fmt.Println(scoredMoves)

	for move, score := range scoredMoves {
		if score > temporaryMaximum {
			temporaryMaximum = score
			nextMove = move
		}
	}

	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
