package simulator

import (
	"fmt"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

func FindOptimal(album *usermodel.Album, guests []*usermodel.OwnedCard, song *models.Song) error {
	defer helper.MeasureTime(time.Now(), "FindOptimal")
	resultChannel := make(chan SimulationSummary)
	i := 0
	for album.Next() {
		i++
		team := album.GetTeam()
		supports, err := album.FindSupportsFor(team, song.Attribute)
		if err != nil {
			return fmt.Errorf("could not find optimal: %v", err)
		}
		for _, guest := range guests {
			gameConfig := NewGameConfig(team, supports, guest, song)
			go func(gameConfig *GameConfig) {
				resultChannel <- Simulate(gameConfig, 10)
			}(gameConfig)
		}
	}

	maxAvg := 0.0
	i = 0
	expectedNumberOfResults := (album.MaxTeamID() + 1) * len(guests)
	var maxAvgSummary SimulationSummary
	for summary := range resultChannel {
		if summary.Average > maxAvg {
			maxAvg = summary.Average
			maxAvgSummary = summary
		}
		i++
		if i == expectedNumberOfResults {
			close(resultChannel)
		}
		if i%100 == 0 {
			fmt.Println(i)
		}
	}
	maxAvgSummary.Report()
	return nil
}
