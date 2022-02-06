package main

//var left = Coord{X: myHead.X - 1, Y: myHead.Y}

func FindNextCoord(state GameState, direction string, start Coord) Coord {

	switch direction {
	case "left":
		return Coord{X: start.X - 1, Y: start.Y}
	case "right":
		return Coord{X: start.X + 1, Y: start.Y}
	case "up":
		return Coord{X: start.X, Y: start.Y + 1}
	case "down":
		return Coord{X: start.X, Y: start.Y - 1}
	}

	return Coord{X: 0, Y: 0}
}

/*
func FindAreaOfEmptyCells(state GameState, direction string) int {
	myHead := state.You.Body[0]

	distanceToLeftWall = myHead.X

	// left

}
*/

// iterate through all the coords looking a given direction, and check if any
// are occupied by any snakes body

func CountCellsUntilYouSeeASnake(state GameState) map[string]int {
	myHead := state.You.Body[0]
	freeCells := map[string]int{
		"up":    0,
		"down":  0,
		"left":  0,
		"right": 0,
	}

	// left
	for i := myHead.X; i >= 1; i-- {
		if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", Coord{X: i, Y: myHead.Y})) {
			freeCells["left"] += 1
		} else {
			break
		}
	}

	for i := myHead.X; i < state.Board.Width-1; i++ {
		if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", Coord{X: i, Y: myHead.Y})) {
			freeCells["right"] += 1
		} else {
			break
		}
	}

	for i := myHead.Y; i >= 1; i-- {
		if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", Coord{X: myHead.X, Y: i})) {
			freeCells["down"] += 1
		} else {
			break
		}
	}

	for i := myHead.Y; i < state.Board.Height-1; i++ {
		if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", Coord{X: myHead.X, Y: i})) {
			freeCells["up"] += 1
		} else {
			break
		}
	}

	return freeCells
}

func CheckCoordForAnySnake(state GameState, target Coord) bool {
	snakes := state.Board.Snakes
	for i := 0; i < len(snakes); i++ {
		for j := 0; j < int(snakes[i].Length); j++ {
			if target.X == snakes[i].Body[j].X && target.Y == snakes[i].Body[j].Y {
				return true
			}
		}
	}
	return false
}
