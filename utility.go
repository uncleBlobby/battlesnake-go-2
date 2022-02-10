package main

import (
	"fmt"
	"math"
)

func FindClosestFoodDistance(state GameState) Food {
	food := state.Board.Food
	myHead := state.You.Body[0]

	foodList := []Food{}
	for i := 0; i < len(food); i++ {
		var newFood Food
		newFood.X = food[i].X
		newFood.Y = food[i].Y
		newFood.Distance.X = int(math.Abs(float64(food[i].X - myHead.X)))
		newFood.Distance.Y = int(math.Abs(float64(food[i].Y - myHead.Y)))
		newFood.Distance.Total = newFood.Distance.X + newFood.Distance.Y
		foodList = append(foodList, newFood)
	}

	// sort food list so first member is closest
	closestFood := foodList[0]
	for i := 1; i < len(foodList); i++ {
		if foodList[i].Distance.Total < closestFood.Distance.Total {
			closestFood = foodList[i]
		}
	}

	// can also return foodList (foodlist with distances)
	// removed for the time being

	return closestFood
}

func findPathToClosestFood(state GameState, closestFood Food) /*[]Coord*/ {
	// find X distance toward food
	// find Y distance toward food

	myHead := state.You.Head
	var closestFoodCoord Coord
	closestFoodCoord.X = closestFood.X
	closestFoodCoord.Y = closestFood.Y

	fmt.Printf("Closest Food: %v\n", closestFood)
	fmt.Printf("Closest Food Distance X: %v\n", closestFood.Distance.X)
	fmt.Printf("Closest Food Distance Y: %v\n", closestFood.Distance.Y)
	fmt.Printf("Closest Food Total Dist. : %v\n", closestFood.Distance.Total)

	fmt.Printf("Path to Food: %v", FindAnyPath(state, myHead, closestFoodCoord))
}

func returnManhattanDistance(start Coord, end Coord) int {
	xDistance := int(math.Abs(float64(start.X - end.X)))
	yDistance := int(math.Abs(float64(start.Y - end.Y)))
	return xDistance + yDistance
}

func returnRelativeXDistance(start Coord, end Coord) int {
	return start.X - end.X
}

func returnRelativeYDistance(start Coord, end Coord) int {
	return start.Y - end.Y
}

func FindAnyPath(state GameState, start Coord, end Coord) []Coord {
	shortestPathLength := returnManhattanDistance(start, end)
	fmt.Printf("Manhattan distance to food: %v\n", shortestPathLength)

	pathXCoords := []int{}
	pathYCoords := []int{}
	// shortest path will be a slice of length = manhattan distance

	if start.X > end.X {
		fmt.Println("Path must move left.")
		// path must move left
		for i := start.X - 1; i >= end.X; i-- {
			pathXCoords = append(pathXCoords, i)
		}
	}
	if start.X < end.X {
		// path must move right
		for i := start.X + 1; i <= end.X; i++ {
			pathXCoords = append(pathXCoords, i)
		}
	}
	if start.X == end.X {
		// path moves neither left nor right
		for i := 0; i < shortestPathLength; i++ {
			pathXCoords = append(pathXCoords, start.X)
		}
	}

	if start.Y > end.Y {
		// path must move down
		for i := start.Y - 1; i >= end.Y; i-- {
			pathYCoords = append(pathYCoords, i)
		}
	}

	if start.Y < end.Y {
		// path must move up
		for i := start.Y + 1; i <= end.Y; i++ {
			pathYCoords = append(pathYCoords, i)
		}
	}

	if start.Y == end.Y {
		// path moves neither up nor down
		for i := 0; i < shortestPathLength; i++ {
			pathYCoords = append(pathYCoords, start.Y)
		}
	}

	// take the two slices of ints for each coord parameter and mash
	// ASSEMBLE PATH
	// DO ALL X MOVES FIRST< THEN DO ALL Y MOVES (or vice versa)
	// CANNOT DO BOTH AT THE SAME TIME, DIAGONALS ARE NOT ALLOWED

	XFIRSTpathCoords := []Coord{}
	YFIRSTpathCoords := []Coord{}

	fmt.Println("Shortest path length: ", shortestPathLength)
	fmt.Println("X Path Coords: ", pathXCoords)
	fmt.Println("Y Path Coords: ", pathYCoords)

	// X FIRST
	// Do all the X moves

	for i := 0; i < len(pathXCoords); i++ {
		var CoordToAdd Coord
		CoordToAdd.X = pathXCoords[i]
		CoordToAdd.Y = start.Y
		XFIRSTpathCoords = append(XFIRSTpathCoords, CoordToAdd)
	}

	// Then all the Y moves
	for i := 0; i < len(pathYCoords); i++ {
		var CoordToAdd Coord
		CoordToAdd.X = end.X
		CoordToAdd.Y = pathYCoords[i]
		XFIRSTpathCoords = append(XFIRSTpathCoords, CoordToAdd)
	}

	// Y FIRST
	// Do all the Y moves

	for i := 0; i < len(pathYCoords); i++ {
		var CoordToAdd Coord
		CoordToAdd.X = start.X
		CoordToAdd.Y = pathYCoords[i]
		YFIRSTpathCoords = append(YFIRSTpathCoords, CoordToAdd)
	}

	for i := 0; i < len(pathXCoords); i++ {
		var CoordToAdd Coord
		CoordToAdd.X = pathXCoords[i]
		CoordToAdd.Y = end.Y
		YFIRSTpathCoords = append(YFIRSTpathCoords, CoordToAdd)
	}

	// Decide Which set of path coords is safer

	fmt.Println("X First Path: ", XFIRSTpathCoords)
	fmt.Println("Y First Path: ", YFIRSTpathCoords)

	xPathFirstIsBlocked := checkPathForEnemies(state, XFIRSTpathCoords)
	if xPathFirstIsBlocked {
		fmt.Println("Moving X first is blocked by a snake.")
	}
	if !xPathFirstIsBlocked {
		fmt.Println("Moving X first is clear of any snake.")
		pathCoords := XFIRSTpathCoords
		return pathCoords
	}
	yPathFirstIsBlocked := checkPathForEnemies(state, YFIRSTpathCoords)
	if yPathFirstIsBlocked {
		fmt.Println("Moving Y first is blocked by a snake.")
	}
	if !yPathFirstIsBlocked {
		fmt.Println("Moving Y first is clear of any snakes.")
		pathCoords := YFIRSTpathCoords
		return pathCoords
	}
	var pathCoords []Coord
	pathCoords = append(pathCoords, Coord{X: start.X, Y: start.Y})
	fmt.Println("Should not have got this far.")
	return pathCoords
}

func checkPathForEnemies(state GameState, path []Coord) bool {
	snakes := state.Board.Snakes
	for i := 0; i < len(path); i++ {
		for j := 0; j < len(snakes); j++ {
			for k := 0; k < int(snakes[j].Length); k++ {
				if path[i].X == snakes[j].Body[k].X && path[i].Y == snakes[j].Body[k].Y {
					return true
				}
			}
		}
	}
	return false
}
