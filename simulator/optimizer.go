package simulator

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

func IsValid(team *usermodel.Team) bool {
	for _, ocard := range team.Ocards {
		if ocard.Card.Skill.SkillType.Name == enum.SkillTypeConcentration {
			return false
		}
		if !helper.IsLeadSkillImplemented(ocard.Card.LeadSkill.Name) {
			return false
		}
		if !helper.IsSkillImplemented(ocard.Card.Skill.SkillType.Name) {
			return false
		}
		if ocard.Card.Rarity.Rarity != enum.RaritySSR {
			return false
		}
	}
	return true
}

func FindOptimal(album *usermodel.Album, guests []*usermodel.OwnedCard, song *models.Song) error {
	defer helper.MeasureTime(time.Now(), "FindOptimal")
	resultChannel := make(chan SimulationSummary)
	i := 0
	expectedNumberOfResults := (album.MaxTeamID() + 1) * len(guests)
	fmt.Printf("Running %d samples:", expectedNumberOfResults)

	wg := sync.WaitGroup{}
	waitGroupInitiateJob := sync.WaitGroup{}
	waitGroupInitiateJob.Add(1)
	for album.Next() {
		i++
		team := album.GetTeam()
		if !IsValid(team) {
			continue
		}
		supports, err := album.FindSupportsFor(team, song.Attribute)
		if err != nil {
			return fmt.Errorf("could not find optimal: %v", err)
		}
		for _, guest := range guests {
			gameConfig := NewGameConfig(team, supports, guest, song)
			wg.Add(1)
			go func(gameConfig *GameConfig) {
				defer wg.Done()
				resultChannel <- Simulate(gameConfig, 100)
			}(gameConfig)
		}
	}
	waitGroupInitiateJob.Done()
	go func() {
		waitGroupInitiateJob.Wait()
		wg.Wait()
		close(resultChannel)
	}()

	maxAvg := 0.0
	i = 0
	var summaries []SimulationSummary
	var maxAvgSummary SimulationSummary
	for summary := range resultChannel {
		summaries = append(summaries, summary)
		if summary.Average > maxAvg {
			maxAvg = summary.Average
			maxAvgSummary = summary
		}
		i++
	}
	fmt.Printf("%d summaries received\n", len(summaries))
	fmt.Printf("Game config with highest average score:")
	maxAvgSummary.Report()
	sort.SliceStable(summaries, func(i int, j int) bool {
		return summaries[i].Average > summaries[j].Average
	})
	for _, summary := range summaries[:10] {
		fmt.Println("----------------------")
		fmt.Println(summary.GameConfig.team)
		fmt.Println(summary.GameConfig.guest)
		fmt.Printf("%.2f\n", summary.Average)
		fmt.Println("----------------------")
	}
	return nil
}
