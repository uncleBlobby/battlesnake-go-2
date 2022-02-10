package main

func AvoidNeck(state GameState, scoredMoves map[string]int) {

	myHead := state.You.Body[0]
	myNeck := state.You.Body[1]

	if myNeck.X < myHead.X {
		scoredMoves["left"] -= 100
	} else if myNeck.X > myHead.X {
		scoredMoves["right"] += -100
	} else if myNeck.Y < myHead.Y {
		scoredMoves["down"] += -100
	} else if myNeck.Y > myHead.Y {
		scoredMoves["up"] += -100
	}

}

func AvoidWalls(state GameState, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
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
}

func AvoidOwnBody(state GameState, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
	mybody := state.You.Body
	for i := 1; i < int(state.You.Length); i++ {
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
}

func AvoidOtherSnakes(state GameState, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
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
}

func FindCloseFood(state GameState, healthFactor int, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
	foods := state.Board.Food

	for i := 0; i < len(foods); i++ {
		if (myHead.X+1 == foods[i].X) && (myHead.Y == foods[i].Y) {
			scoredMoves["right"] += healthFactor
		}
		if (myHead.X-1 == foods[i].X) && (myHead.Y == foods[i].Y) {
			scoredMoves["left"] += healthFactor
		}
		if (myHead.X == foods[i].X) && (myHead.Y+1 == foods[i].Y) {
			scoredMoves["up"] += healthFactor
		}
		if (myHead.X == foods[i].X) && (myHead.Y-1 == foods[i].Y) {
			scoredMoves["down"] += healthFactor
		}
	}
}

func FindDistantFood(state GameState, healthFactor int, scoredMoves map[string]int) {
	myHead := state.You.Head
	closestFood := FindClosestFoodDistance(state)
	pathToClosestFood := FindAnyPath(state, myHead, Coord{closestFood.X, closestFood.Y})
	if len(pathToClosestFood) > 0 {
		if pathToClosestFood[0].X > myHead.X {
			scoredMoves["right"] += healthFactor / len(pathToClosestFood)
		}
		if pathToClosestFood[0].X < myHead.X {
			scoredMoves["left"] += healthFactor / len(pathToClosestFood)
		}
		if pathToClosestFood[0].Y > myHead.Y {
			scoredMoves["up"] += healthFactor / len(pathToClosestFood)
		}
		if pathToClosestFood[0].Y < myHead.Y {
			scoredMoves["up"] += healthFactor / len(pathToClosestFood)
		}
	}
}

func PreferNotSaucyMoves(state GameState, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
	hazards := state.Board.Hazards

	if len(hazards) > 0 {
		for i := 0; i < len(hazards); i++ {
			if (myHead.X-1 == hazards[i].X) && (myHead.Y == hazards[i].Y) {
				scoredMoves["left"] += -50
			}
			if (myHead.X+1 == hazards[i].X) && (myHead.Y == hazards[i].Y) {
				scoredMoves["right"] += -50
			}
			if (myHead.X == hazards[i].X) && (myHead.Y-1 == hazards[i].Y) {
				scoredMoves["down"] += -50
			}
			if (myHead.X == hazards[i].X) && (myHead.Y+1 == hazards[i].Y) {
				scoredMoves["up"] += -50
			}
		}
	}
}

func CheckIfCurrentlyInSauce(state GameState) bool {
	myHead := state.You.Body[0]
	hazards := state.Board.Hazards

	for i := 0; i < len(hazards); i++ {
		if (hazards[i].X == myHead.X) && (hazards[i].Y == myHead.Y) {
			return true
		}
	}

	return false
}

func MoveOutOfSauce(state GameState, scoredMoves map[string]int) {
	myHead := state.You.Body[0]
	hazards := state.Board.Hazards
	for i := 0; i < len(hazards); i++ {
		if FindNextCoord(state, "left", myHead) == hazards[i] {
			scoredMoves["left"] -= 75
		}
		if FindNextCoord(state, "right", myHead) == hazards[i] {
			scoredMoves["right"] -= 75
		}
		if FindNextCoord(state, "up", myHead) == hazards[i] {
			scoredMoves["up"] -= 75
		}
		if FindNextCoord(state, "down", myHead) == hazards[i] {
			scoredMoves["down"] -= 75
		}
	}

}
