package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
)

// SimulationSummary is summary of the simulation
type SimulationSummary struct {
	GameConfig Playable
	Min        int
	Max        int
	Average    float64
	SimCount   int
	Results    []int
}

type GameLike interface {
	Play(bool) *GameState
}

// Simulate simulates the game `times` times and return the summary in SimulationSummary
func Simulate(gc Playable, times int) SimulationSummary {
	game := NewGame(gc)
	// game := NewGameFast(gc)
	maxScore := game.Play(true).Score
	if helper.Features.LimitScore() {
		if !gc.isResonantActive() {
			if maxScore < helper.Features.GetScoreLimitForAttr(gc.getSong().Attribute) {
				return SimulationSummary{
					GameConfig: gc,
					Min:        maxScore,
					Max:        maxScore,
					Average:    float64(maxScore),
					SimCount:   -1,
					Results:    []int{maxScore},
				}
			}
		}
	}
	// game := NewGame(gc)
	resultChannel := make(chan int, times)
	goodRolls := helper.Features.AlwaysGoodRolls()
	for i := 0; i < times; i++ {
		go func(game GameLike, i int) {
			// randSeed := (time.Now().UnixNano() * int64(i+1)) % math.MaxInt64
			state := game.Play(goodRolls)
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
