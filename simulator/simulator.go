package simulator

// SimulationSummary is summary of the simulation
type SimulationSummary struct {
	GameConfig Playable
	Min        int
	Max        int
	Average    float64
	SimCount   int
	Results    []int
}

// Simulate simulates the game `times` times and return the summary in SimulationSummary
func Simulate(gc Playable, times int) SimulationSummary {
	// defer helper.MeasureTime(time.Now(), "Simulate")
	game := NewGame(gc)
	// game := NewGame(gc)
	resultChannel := make(chan int, times)
	for i := 0; i < times; i++ {
		go func(game *Game, i int) {
			// randSeed := (time.Now().UnixNano() * int64(i+1)) % math.MaxInt64
			state := game.Play()
			resultChannel <- state.Score
		}(game, i)
	}
	i := 0
	result := SimulationSummary{GameConfig: gc, Min: 999999999, Results: make([]int, 0, times)}
	sum := 0
	for score := range resultChannel {
		result.Results = append(result.Results, score)
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
