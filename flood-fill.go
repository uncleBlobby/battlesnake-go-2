package main

func CountOpenCells(state GameState, start Coord, direction string) int {
	var openCells int
	myHead := state.You.Head
	var depth int = 0

	if direction == "left" {
		for !CheckAnyCoordLEFTForEnemies(state, start, depth) {
			CheckAnyCoordLEFTForEnemies(state, start, depth)
		}
	}
	if direction == "right" {

	}

	return openCells
}
