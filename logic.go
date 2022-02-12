package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
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
		Color:      "#8a2be2",     // TODO: Personalize
		Head:       "shades",      // TODO: Personalize
		Tail:       "bolt",        // TODO: Personalize
		Version:    "0.1.4-beta",  // updated feb 12 wrapped hack
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

	AvoidNeck(state, scoredMoves)
	if state.Game.Ruleset.Name != "wrapped" {
		AvoidWalls(state, scoredMoves)
	}
	AvoidOwnBody(state, scoredMoves)
	AvoidOtherSnakes(state, scoredMoves)
	PreferNotSaucyMoves(state, scoredMoves)
	FindCloseFood(state, FindMissingHealth(state), scoredMoves)
	TurnAwayFromLargerSnakes(state, scoredMoves)
	TurnTowardSmallerSnakes(state, scoredMoves)

	viewDistances := CountCellsUntilYouSeeASnake(state)

	if state.You.Health <= 15 {
		FindDistantFood(state, FindMissingHealth(state), scoredMoves)
	}

	//fmt.Println(CountCellsUntilYouSeeASnake(state))

	return BattlesnakeMoveResponse{
		Move: ChooseMove(scoredMoves, viewDistances),
	}
}
