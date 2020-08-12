package simulator

// SimulationSummary is summary of the simulation
type SimulationSummary struct {
	Min      int
	Max      int
	Average  float64
	SimCount int
}

// Simulate simulates the game `times` times and return the summary in SimulationSummary
func Simulate(gc *GameConfig, times int) SimulationSummary {
	game := NewGame(gc)
	// var results []int
	resultChannel := make(chan int)
	for i := 0; i < times; i++ {
		go func(game *Game, isLast bool) {
			state := game.Play()
			resultChannel <- state.Score
		}(game, i == times-1)
	}
	i := 0
	result := SimulationSummary{Min: 999999999}
	sum := 0
	for score := range resultChannel {
		if score > result.Max {
			result.Max = score
		}
		if score < result.Min {
			result.Min = score
		}
		sum += score
		result.SimCount++

		i++
		if i == times {
			close(resultChannel)
		}
	}
	result.Average = float64(sum) / float64(result.SimCount)
	return result
}
