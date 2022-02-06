package main

import (
	"fmt"
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
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
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

func TestOwnBodyAvoidance1(t *testing.T) {
	// Arrange
	me := Battlesnake{

		//Snake is boxed into itself here, can only move up.
		//
		//       BHB
		//		 BBB
		//
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 4, Y: 5}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 6, Y: 5}},
		Length: 6,
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move != "up" {
			t.Errorf("snake moved into its own body, %s", nextMove.Move)
		}
	}
}

func TestOwnBodyAvoidance2(t *testing.T) {
	// Arrange
	me := Battlesnake{

		//Snake is boxed into itself here, can only move up.
		//
		//       BBB
		//		 BHB
		//
		Head:   Coord{X: 5, Y: 4},
		Body:   []Coord{{X: 5, Y: 4}, {X: 4, Y: 4}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 6, Y: 4}},
		Length: 6,
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move != "down" {
			t.Errorf("snake moved into its own body, %s", nextMove.Move)
		}
	}
}

func TestOwnBodyAvoidance3(t *testing.T) {
	// Arrange
	me := Battlesnake{

		//Snake is boxed into itself here, can only move up.
		//
		//       BB
		//		 HB
		//		 BB
		Head:   Coord{X: 5, Y: 5},
		Body:   []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 6, Y: 5}, {X: 6, Y: 6}, {X: 5, Y: 6}},
		Length: 6,
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move != "left" {
			t.Errorf("snake moved into its own body, %s", nextMove.Move)
		}
	}
}

func TestOwnBodyAvoidance4(t *testing.T) {
	// Arrange
	me := Battlesnake{

		//Snake is boxed into itself here, can only move up.
		//
		//       BB
		//		 BH
		//		 BB
		Head:   Coord{X: 6, Y: 5},
		Body:   []Coord{{X: 6, Y: 5}, {X: 6, Y: 4}, {X: 5, Y: 4}, {X: 5, Y: 5}, {X: 5, Y: 6}, {X: 6, Y: 6}},
		Length: 6,
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me},
		},
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move != "right" {
			t.Errorf("snake moved into its own body, %s", nextMove.Move)
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

func TestAvoidSnakes1(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 4, Y: 5}, {X: 3, Y: 5}},
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 6, Y: 6},
		Body:   []Coord{{X: 6, Y: 6}, {X: 6, Y: 5}, {X: 6, Y: 4}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
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
			t.Errorf("snake moved into an enemy, %s", nextMove.Move)
		}
	}
}

func TestAvoidSnakes2(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing left
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 6, Y: 5}, {X: 7, Y: 5}},
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 4},
		Body:   []Coord{{X: 4, Y: 4}, {X: 4, Y: 5}, {X: 4, Y: 6}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
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
			t.Errorf("snake moved into an enemy, %s", nextMove.Move)
		}
	}
}

func TestAvoidSnakes3(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing up
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 5, Y: 3}},
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 6},
		Body:   []Coord{{X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
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
			t.Errorf("snake moved into an enemy, %s", nextMove.Move)
		}
	}
}

func TestAvoidSnakes4(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}
	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 4},
		Body:   []Coord{{X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}},
		Length: 3,
	}
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
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
			t.Errorf("snake moved into an enemy, %s", nextMove.Move)
		}
	}
}

func TestFindCloseFood(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}
	/*
		enemy1 := Battlesnake{
			Head:   Coord{X: 4, Y: 4},
			Body:   []Coord{{X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}},
			Length: 3,
		}
	*/
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
			Food:   []Coord{{X: 5, Y: 4}},
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
		if nextMove.Move != "down" {
			t.Errorf("snake missed food, %s", nextMove.Move)
		}
	}
}

func TestCountEmptyCellsLeft(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 2, Y: 4},
		Body:   []Coord{{X: 2, Y: 4}, {X: 2, Y: 5}, {X: 2, Y: 6}},
		Length: 3,
	}

	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
			Food:   []Coord{{X: 5, Y: 4}},
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
		viewDistance := CountCellsUntilYouSeeASnake(state)
		//nextMove := move(state)
		// Assert never move right
		if viewDistance["left"] > viewDistance["right"] {
			t.Errorf("view distance, %v", viewDistance)
		}
		fmt.Println(viewDistance)
	}
}

// TODO: More GameState test cases!
