package main

import "math"

func findFoodDistances(state GameState) Food {
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
