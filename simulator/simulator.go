package simulator

import (
	"fmt"
	"strings"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"gonum.org/v1/gonum/stat"
)

type miniReport struct {
	avg    float64
	report string
}

// SimulationSummary is summary of the simulation
type SimulationSummary struct {
	GameConfig   simulatormodels.Playable
	Min          int
	Max          int
	Average      float64
	SimCount     int
	Results      []int
	resultsFloat []float64
	stddev       float64
}

func (ss *SimulationSummary) calc() {
	ss.stddev = stat.StdDev(ss.resultsFloat, nil)
}

func (ss SimulationSummary) Minify() miniReport {
	return miniReport{ss.Average, ss.ReportOneline()}
}

func (ss SimulationSummary) ReportOneline() string {
	// Score: 100123 (Score Bonus 80.00%), ........, Leader Index = 4, 300096 <LEad Skill>
	ret := []string{
		fmt.Sprintf("Score: %.2f", ss.Average),
	}
	cards := []string{}
	for _, ocard := range ss.GameConfig.GetCards() {
		cards = append(cards, fmt.Sprintf("%d (%s %.2f%%)", ocard.Card.ID,
			ocard.Card.Skill.SkillType.Name, float64(ocard.SkillProcChance)/100.0))
	}
	ret = append(ret, strings.Join(cards, ","))
	ret = append(ret, fmt.Sprintf("Leader index: %d", ss.GameConfig.GetLeaderIndex()))
	for _, ocard := range ss.GameConfig.GetLeadSkillActivableCards() {
		ret = append(ret, fmt.Sprintf("%d %s", ocard.Card.ID, ocard.Card.LeadSkill.Name))
	}
	ret = append(ret, fmt.Sprintf("%d times", ss.SimCount))
	return strings.Join(ret, "|")
}

func (ss SimulationSummary) Report() string {
	ret := []string{
		"-----Simulation Summary-----",
		fmt.Sprintf("Sim done  = %d times", ss.SimCount),
		fmt.Sprintf("Min score = %d", ss.Min),
		fmt.Sprintf("Max score = %d", ss.Max),
		fmt.Sprintf("Average   = %.2f", ss.Average),
		fmt.Sprintf("Std Dev   = %.2f", ss.stddev),
	}
	cards := []string{}
	for _, ocard := range ss.GameConfig.GetCards() {
		cards = append(cards, fmt.Sprintf("%d (%s %.2f%%)", ocard.Card.ID,
			ocard.Card.Skill.SkillType.Name, float64(ocard.SkillProcChance)/100.0))
	}
	ret = append(ret, strings.Join(cards, ","))
	ret = append(ret, fmt.Sprintf("Leader index = %d", ss.GameConfig.GetLeaderIndex()))
	for _, ocard := range ss.GameConfig.GetLeadSkillActivableCards() {
		ret = append(ret, fmt.Sprintf("Lead skill activable: %d %s", ocard.Card.ID, ocard.Card.LeadSkill.Name))
	}
	ret = append(ret, "---------------------")
	return strings.Join(ret, "\n")
}

type GameLike interface {
	Play(bool) *GameState
}

var _ GameLike = GameFast{}

// Simulate simulates the game `times` times and return the summary in SimulationSummary
func Simulate(gc simulatormodels.Playable, times int) SimulationSummary {
	var game GameLike
	if helper.Features.UseFastGame() {
		game = NewGameFast(gc)
	} else {
		game = NewGame(gc)
	}
	maxScore := game.Play(true).Score
	if helper.Features.LimitScore() {
		if !gc.IsResonantActive() {
			threshold := helper.Features.GetScoreLimitForAttr(gc.GetSong().Attribute, gc.GetSong().Level)
			// fmt.Println(threshold)
			if maxScore < threshold {
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
	allSkills100Percent := true
	for _, ocard := range gc.GetSkillActivableCards() {
		if ocard.Skill.SkillType.IsActive(gc.GetTeamAttributesv2()) {
			probMultiplier := 1.0
			if ocard.Card.Idol.Attribute == gc.GetSong().Attribute || gc.GetSong().Attribute == enum.AttrAll {
				probMultiplier += 0.3
			}
			for _, leader := range gc.GetLeadSkillActivableCards() {
				probMultiplier += leader.LeadSkill.SkillProbBonus(
					leader.Card.Rarity.Rarity,
					ocard.Card.Idol.Attribute,
				)
			}
			if float64(ocard.SkillProcChance)/10000.0*probMultiplier < 100.0 {
				allSkills100Percent = false
			}
		}
	}
	if allSkills100Percent {
		times = 1
	}
	// game := NewGame(gc)
	resultChannel := make(chan int, times)
	goodRolls := helper.Features.AlwaysGoodRolls()
	for i := 0; i < times; i++ {
		go func(game GameLike, i int) {
			// randSeed := (time.Now().UnixNano() * int64(i+1)) % math.MaxInt64
			// game.Play
			state := game.Play(goodRolls)
			fmt.Printf("outside = %d\n", state.Score)
			resultChannel <- state.Score
		}(game, i)
	}
	i := 0
	sum := 0
	result := SimulationSummary{GameConfig: gc, Min: 999999999, Results: make([]int, 0, times)}
	// fmt.Println("-----")
	for score := range resultChannel {
		// fmt.Println(score)
		result.Results = append(result.Results, score)
		result.resultsFloat = append(result.resultsFloat, float64(score))
		if score > result.Max {
			result.Max = score
		}
		if score < result.Min {
			result.Min = score
		}
		result.SimCount++
		sum += score
		i++
		if i == times {
			close(resultChannel)
		}
	}
	result.Average = float64(sum) / float64(result.SimCount)
	result.calc()
	return result
}
