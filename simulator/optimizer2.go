package simulator

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/logic"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// GetFilteredGameConfigs create game configs given the parameters, filtered by some logics
func GetFilteredGameConfigs(album *usermodel.Album2, guests []*usermodel.OwnedCard,
	song *models.Song, times int, filename string) <-chan *simulatormodels.GameConfig {
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
