package simulator

import (
	"fmt"
	"sort"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/logic"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// GetFilteredGameConfigs create game configs given the parameters, filtered by some logics
func GetFilteredGameConfigs(album *usermodel.Album2, guests []*usermodel.OwnedCard,
	song *models.Song) <-chan *simulatormodels.GameConfig {
	ch := make(chan *simulatormodels.GameConfig)

	cardsCount := 0
	teamCount := 0
	gcCount := 0
	final := 0
	go func() {
		cardLogicHandler := logic.NewCardLogicHandler()
		teamLogicHandler := logic.NewTeamLogicHandler()
		gameConfigLogicHandler := logic.NewGameConfigLogicHandler()

		// beneran := helper.Features.ReallySimulate()
		// resultChannel := make(chan SimulationSummary)
		for album.Next() {
			cards := album.GetCards()
			if !cardLogicHandler.IsOk(cards, song) {
				continue
			}
			cardsCount++

			for leaderIndex := range cards {
				team := &usermodel.Team{Ocards: cards, LeaderIndex: leaderIndex}
				if !teamLogicHandler.IsOk(team, song) {
					continue
				}
				teamCount++

				supports, err := album.FindSupportsFor(team, song.Attribute)
				if err != nil {
					panic(err)
				}
				tmpG := []*usermodel.OwnedCard{}
				for _, guest := range guests {
					if !gameConfigLogicHandler.IsOk(team, guest, song) {
						continue
					}
					gcCount++
					tmpG = append(tmpG, guest)
					gc := simulatormodels.NewGameConfig(team.Ocards[:], leaderIndex,
						supports[:], guest, song, 0, statcalculator.NormalStatCalculator)
					if helper.Features.LimitAppeals() {
						if gc.GetAppeal() < 310000 {
							if !gc.IsResonantActive() {
								continue
							}
						}
					}
					final++
					ch <- gc
				}
				// if len(tmpG) > 1 {
				// 	logger.Logf("%d guest for: ", len(tmpG))
				// 	str := ""
				// 	for _, ocard := range team.Ocards {
				// 		str = fmt.Sprintf("%s, %d (%s)(%s) ", str, ocard.Card.ID, ocard.Card.Idol.Attribute, ocard.Skill.SkillType.Name)
				// 	}
				// 	logger.Log(str)
				// 	logger.Logf("%d: %v", team.LeaderIndex, team.Leader().LeadSkill.Name)
				// 	str = ""
				// 	for _, ocard := range tmpG {
				// 		str = fmt.Sprintf("%s, %d (%s)", str, ocard.Card.ID, ocard.Card.LeadSkill.Name)
				// 	}
				// 	logger.Log(str)
				// }
			}
		}

		// debug stats
		m := album.MaxTeamID() + 1
		fmt.Printf("Without filter = %10d -> %10d -> %10d\n", m, m*5, m*5*len(guests))
		fmt.Printf("Cards filter   = %10d -> %10d -> %10d\n", cardsCount, cardsCount*5, cardsCount*5*len(guests))
		fmt.Printf("Team filter    = %10d -> %10d -> %10d\n", 0, teamCount, teamCount*len(guests))
		fmt.Printf("Gc filter      = %10d -> %10d -> %10d\n", 0, 0, gcCount)
		fmt.Printf("Appeal filter  = %10d -> %10d -> %10d\n", 0, 0, final)
		//
		close(ch)
	}()
	return ch
}

func Optimize(generator <-chan *simulatormodels.GameConfig, times int, filename string) error {
	beneran := helper.Features.ReallySimulate()
	// channel := make(chan SimulationSummary)
	channel2 := make(chan miniReport)
	count := 0
	for gc := range generator {
		if !beneran {
			continue
		}
		count++

		go func(gameConfig *simulatormodels.GameConfig, channel chan miniReport, times int) {
			result := Simulate(gameConfig, times)
			channel2 <- result.Minify()
		}(gc, channel2, times)
	}

	var results []miniReport
	i := 0
	for result := range channel2 {
		results = append(results, result)
		i++
		if i == count {
			close(channel2)
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].avg > results[j].avg
	})
	buffer := []string{}
	for _, result := range results {
		buffer = append(buffer, result.report)
	}
	saveBuffer(&buffer, filename)
	return nil

	// var summaries []SimulationSummary
	// i := 0
	// for summary := range channel {
	// 	summaries = append(summaries, summary)
	// 	i++
	// 	if i == count {
	// 		close(channel)
	// 	}
	// }

	// sort.SliceStable(summaries, func(i, j int) bool {
	// 	return summaries[i].Max > summaries[j].Max
	// })
	// buffer := []string{}
	// for _, summary := range summaries {
	// 	buffer = append(buffer, summary.ReportOneline())
	// }
	// saveBuffer(&buffer, filename)
	// return nil
}
