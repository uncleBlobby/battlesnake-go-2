package main

import (
	"testing"
)

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 2, Y: 0},
		Body: []Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

func TestRightHandWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 10, Y: 0},
		Body: []Coord{{X: 10, Y: 0}, {X: 9, Y: 0}, {X: 8, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move right
		if nextMove.Move == "right" {
			t.Errorf("snake moved into the right-hand wall, %s", nextMove.Move)
		}
	}
}

func TestLeftHandWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 0, Y: 0},
		Body: []Coord{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move right
		if nextMove.Move == "left" {
			t.Errorf("snake moved into the left-hand wall, %s", nextMove.Move)
		}
	}
}

func TestBottomWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 5, Y: 0},
		Body: []Coord{{X: 5, Y: 0}, {X: 4, Y: 0}, {X: 3, Y: 0}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move right
		if nextMove.Move == "down" {
			t.Errorf("snake moved into the bottom wall, %s", nextMove.Move)
		}
	}
}

func TestTopWallAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 5, Y: 10},
		Body: []Coord{{X: 5, Y: 10}, {X: 4, Y: 10}, {X: 3, Y: 10}},
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move right
		if nextMove.Move == "up" {
			t.Errorf("snake moved into the top wall, %s", nextMove.Move)
		}
	}
}

// TODO: More GameState test cases!
