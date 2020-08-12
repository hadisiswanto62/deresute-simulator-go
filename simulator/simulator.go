package simulator

import (
	"fmt"
	"math"
	"time"

	"github.com/aybabtme/uniplot/histogram"
)

// SimulationSummary is summary of the simulation
type SimulationSummary struct {
	*GameConfig
	Min      int
	Max      int
	Average  float64
	SimCount int
	Results  []int
}

// Report reports the simulation summary
func (ss SimulationSummary) Report() {
	var data []float64
	for _, score := range ss.Results {
		data = append(data, float64(score/1000))
	}
	fmt.Println("--- Summary ---")
	fmt.Printf("Team:\n")
	for i, ocard := range ss.GameConfig.team.Ocards {
		fmt.Printf("     %d. %s", i, ocard)
		if i == ss.GameConfig.team.LeaderIndex {
			fmt.Printf(" (leader)")
		}
		fmt.Println()
	}
	fmt.Printf("Support:\n")
	for i, ocard := range ss.GameConfig.supports {
		fmt.Printf("     %d. %s\n", i, ocard)
	}
	fmt.Printf("Guest: %s\n", ss.GameConfig.guest)
	fmt.Printf("Song: %s\n", ss.GameConfig.song)
	fmt.Printf("Appeal: %d\n", ss.GameConfig.Appeal)
	fmt.Printf("Played %d times:\n", ss.SimCount)
	hist := histogram.Hist(10, data)
	maxCount, max, min := 0, 0.0, 0.0
	for _, bucket := range hist.Buckets {
		if bucket.Count > maxCount {
			maxCount = bucket.Count
			max = bucket.Max
			min = bucket.Min
		}
	}
	fmt.Printf("Expected value = between %.2f-%.2f\n", min, max)
	fmt.Printf("Min: %d\n", ss.Min)
	fmt.Printf("Max: %d\n", ss.Max)
	fmt.Printf("Avg: %f\n", ss.Average)
	fmt.Printf("------------------\n")
}

// Simulate simulates the game `times` times and return the summary in SimulationSummary
func Simulate(gc *GameConfig, times int) SimulationSummary {
	game := NewGame(gc)
	// var results []int
	resultChannel := make(chan int)
	for i := 0; i < times; i++ {
		go func(game *Game, i int) {
			randSeed := (time.Now().UnixNano() * int64(i+1)) % math.MaxInt64
			state := game.Play(randSeed)
			resultChannel <- state.Score
		}(game, i)
	}
	i := 0
	result := SimulationSummary{GameConfig: gc, Min: 999999999}
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
