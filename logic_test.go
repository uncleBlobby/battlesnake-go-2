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

func TestCountEmptyCellsLeftWRAPPEDMODE(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 7, Y: 4},
		Body:   []Coord{{X: 7, Y: 4}, {X: 7, Y: 5}, {X: 7, Y: 6}},
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
				Name: "wrapped",
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

func TestCountEmptyCellsUpWRAPPEDMODE(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 6, Y: 3},
		Body:   []Coord{{X: 6, Y: 3}, {X: 5, Y: 3}, {X: 4, Y: 3}},
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
				Name: "wrapped",
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
func TestCountEmptyCellsRight(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 8, Y: 4},
		Body:   []Coord{{X: 8, Y: 4}, {X: 8, Y: 5}, {X: 8, Y: 6}},
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
		if viewDistance["right"] > viewDistance["left"] {
			t.Errorf("view distance, %v", viewDistance)
		}
		//fmt.Println(viewDistance)
	}
}

func TestCountEmptyCellsUp(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing up
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 5, Y: 3}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 8},
		Body:   []Coord{{X: 4, Y: 8}, {X: 5, Y: 8}, {X: 6, Y: 8}},
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
		if viewDistance["down"] < viewDistance["up"] {
			t.Errorf("view distance, %v", viewDistance)
		}
		//fmt.Println(viewDistance)
	}
}

func TestCountEmptyCellsDown(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 4, Y: 2},
		Body:   []Coord{{X: 4, Y: 2}, {X: 5, Y: 2}, {X: 6, Y: 2}},
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
		if viewDistance["down"] > viewDistance["up"] {
			t.Errorf("view distance, %v", viewDistance)
		}
		//fmt.Println(viewDistance)
	}
}

func TestFoodDistanceDebug(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}
	/*
		enemy1 := Battlesnake{
			Head:   Coord{X: 4, Y: 2},
			Body:   []Coord{{X: 4, Y: 2}, {X: 5, Y: 2}, {X: 6, Y: 2}},
			Length: 3,
		}
	*/
	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me},
			Food:   []Coord{{X: 10, Y: 2}},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 2; i++ {
		findPathToClosestFood(state, FindClosestFoodDistance(state))
		/*
			if viewDistance["down"] > viewDistance["up"] {
				t.Errorf("view distance, %v", viewDistance)
			}
		*/
		//fmt.Println(viewDistance)
	}
}

func TestFoodDistanceXPATHWithEnemyBlocking(t *testing.T) {
	// Arrange
	// Should be an enemy blocking the XPATH to food,
	// so the YPATH should be the chosen route
	me := Battlesnake{
		// Length 3, facing down
		Head: Coord{X: 5, Y: 5},
		Body: []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}},
	}

	enemy1 := Battlesnake{
		Head:   Coord{X: 7, Y: 5},
		Body:   []Coord{{X: 7, Y: 5}, {X: 8, Y: 5}, {X: 8, Y: 6}},
		Length: 3,
	}

	state := GameState{
		Board: Board{
			Width:  11,
			Height: 11,
			Snakes: []Battlesnake{me, enemy1},
			Food:   []Coord{{X: 10, Y: 2}},
		},
		You: me,
		Game: Game{
			Ruleset: Ruleset{
				Name: "standard",
			},
		},
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 2; i++ {
		findPathToClosestFood(state, FindClosestFoodDistance(state))
		/*
			if viewDistance["down"] > viewDistance["up"] {
				t.Errorf("view distance, %v", viewDistance)
			}
		*/
		//fmt.Println(viewDistance)
	}
}

// TODO: More GameState test cases!
