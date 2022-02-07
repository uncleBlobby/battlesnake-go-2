package main

func TurnTowardSmallerSnakes(state GameState, scoredMoves map[string]int) {
	value := 50
	myHead := state.You.Body[0]
	snakes := state.Board.Snakes

	leftCoord := FindNextCoord(state, "left", myHead)
	rightCoord := FindNextCoord(state, "right", myHead)
	upCoord := FindNextCoord(state, "up", myHead)
	downCoord := FindNextCoord(state, "down", myHead)

	for i := 0; i < len(snakes); i++ {
		if (FindNextCoord(state, "left", leftCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "up", leftCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "down", leftCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "right", rightCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "up", rightCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "down", rightCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "right", upCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "up", upCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "left", upCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "right", downCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["down"] += value
		}
		if (FindNextCoord(state, "down", downCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["down"] += value
		}
		if (FindNextCoord(state, "left", downCoord) == snakes[i].Head) && (state.You.Length > snakes[i].Length) {
			scoredMoves["down"] += value
		}
	}

}

func TurnAwayFromLargerSnakes(state GameState, scoredMoves map[string]int) {
	value := -75
	myHead := state.You.Body[0]
	snakes := state.Board.Snakes

	leftCoord := FindNextCoord(state, "left", myHead)
	rightCoord := FindNextCoord(state, "right", myHead)
	upCoord := FindNextCoord(state, "up", myHead)
	downCoord := FindNextCoord(state, "down", myHead)

	for i := 0; i < len(snakes); i++ {
		if (FindNextCoord(state, "left", leftCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "up", leftCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "down", leftCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["left"] += value
		}
		if (FindNextCoord(state, "right", rightCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "up", rightCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "down", rightCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["right"] += value
		}
		if (FindNextCoord(state, "right", upCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "up", upCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "left", upCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["up"] += value
		}
		if (FindNextCoord(state, "right", downCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["down"] += value
		}
		if (FindNextCoord(state, "down", downCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["down"] += value
		}
		if (FindNextCoord(state, "left", downCoord) == snakes[i].Head) && (state.You.Length <= snakes[i].Length) {
			scoredMoves["down"] += value
		}
	}

}
