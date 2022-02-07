package main

func FindMissingHealth(state GameState) int {
	var missingHealth = 100 - state.You.Health

	return int(missingHealth)
}
