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

// TODO: Build a recursive depth version of the above function
/*
func FindNextCoordAtDepth(state GameState, direction string, start Coord, depth int) {}

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

	if state.Game.Ruleset.Name != "wrapped" {
		for i := myHead.X; i >= 1; i-- {
			currentCoord := Coord{X: i, Y: myHead.Y}
			if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", currentCoord)) {
				freeCells["left"] += 1
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", FindNextCoord(state, "up", currentCoord))) {
					freeCells["left"] += 1
				}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", FindNextCoord(state, "down", currentCoord))) {
					freeCells["left"] += 1
				}
			} else {
				break
			}
		}

		for i := myHead.X; i < state.Board.Width-1; i++ {
			currentCoord := Coord{X: i, Y: myHead.Y}
			if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", currentCoord)) {
				freeCells["right"] += 1
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", FindNextCoord(state, "up", currentCoord))) {
					freeCells["right"] += 1
				}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", FindNextCoord(state, "down", currentCoord))) {
					freeCells["right"] += 1
				}
			} else {
				break
			}
		}

		for i := myHead.Y; i >= 1; i-- {
			currentCoord := Coord{X: myHead.X, Y: i}
			if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", currentCoord)) {
				freeCells["down"] += 1
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", FindNextCoord(state, "left", currentCoord))) {
					freeCells["down"] += 1
				}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", FindNextCoord(state, "right", currentCoord))) {
					freeCells["down"] += 1
				}
			} else {
				break
			}
		}

		for i := myHead.Y; i < state.Board.Height-1; i++ {
			currentCoord := Coord{X: myHead.X, Y: i}
			if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", currentCoord)) {
				freeCells["up"] += 1
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", FindNextCoord(state, "left", currentCoord))) {
					freeCells["up"] += 1
				}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", FindNextCoord(state, "right", currentCoord))) {
					freeCells["up"] += 1
				}
			} else {
				break
			}
		}

		if state.Game.Ruleset.Name == "wrapped" {
			for i := myHead.X; i >= 1; i-- {
				currentCoord := Coord{X: i, Y: myHead.Y}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", currentCoord)) {
					freeCells["left"] += 1
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", FindNextCoord(state, "up", currentCoord))) {
						freeCells["left"] += 1
					}
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "left", FindNextCoord(state, "down", currentCoord))) {
						freeCells["left"] += 1
					}
				} else {
					break
				}
			}

			for i := myHead.X; i < 25; i++ {
				currentCoord := Coord{X: i, Y: myHead.Y}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", currentCoord)) {
					freeCells["right"] += 1
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", FindNextCoord(state, "up", currentCoord))) {
						freeCells["right"] += 1
					}
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "right", FindNextCoord(state, "down", currentCoord))) {
						freeCells["right"] += 1
					}
				} else {
					break
				}
			}

			for i := myHead.Y; i >= 1; i-- {
				currentCoord := Coord{X: myHead.X, Y: i}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", currentCoord)) {
					freeCells["down"] += 1
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", FindNextCoord(state, "left", currentCoord))) {
						freeCells["down"] += 1
					}
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "down", FindNextCoord(state, "right", currentCoord))) {
						freeCells["down"] += 1
					}
				} else {
					break
				}
			}

			for i := myHead.Y; i < 25; i++ {
				currentCoord := Coord{X: myHead.X, Y: i}
				if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", currentCoord)) {
					freeCells["up"] += 1
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", FindNextCoord(state, "left", currentCoord))) {
						freeCells["up"] += 1
					}
					if !CheckCoordForAnySnake(state, FindNextCoord(state, "up", FindNextCoord(state, "right", currentCoord))) {
						freeCells["up"] += 1
					}
				} else {
					break
				}
			}
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
