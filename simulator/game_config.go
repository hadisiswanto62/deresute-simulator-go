package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// GameConfig is a config for normal game (with cards, supports, guest, and song).
// USE NewGameConfig() FOR CREATING
type GameConfig struct {
	ocards      []*usermodel.OwnedCard
	leaderIndex int
	supports    []*usermodel.OwnedCard
	guest       *usermodel.OwnedCard
	song        *models.Song

	leadSkillActivableCards []*usermodel.OwnedCard
	baseVisual              int
	baseVocal               int
	baseDance               int
	appeal                  int
	hp                      int
	teamAttributes          []enum.Attribute
	teamSkills              []enum.SkillType
	resonantOn              bool
	statcalculator          statcalculator.IStatCalculator
	bonusAppeal             int
}

func (gc GameConfig) GetOcards() []*usermodel.OwnedCard {
	return gc.ocards
}

func (gc GameConfig) Appeal() int {
	return gc.getAppeal()
}

func (gc GameConfig) getSkillActivableCards() []*usermodel.OwnedCard {
	return gc.ocards
}
func (gc GameConfig) getLeadSkillActivableCards() []*usermodel.OwnedCard {
	return []*usermodel.OwnedCard{
		gc.ocards[gc.leaderIndex],
		gc.guest,
	}
}
func (gc GameConfig) getSong() *models.Song {
	return gc.song
}
func (gc GameConfig) getBaseVisual() int {
	return gc.baseVisual
}
func (gc GameConfig) getBaseVocal() int {
	return gc.baseVocal
}
func (gc GameConfig) getBaseDance() int {
	return gc.baseDance
}
func (gc GameConfig) getAppeal() int {
	return gc.appeal
}
func (gc GameConfig) getHp() int {
	return gc.hp
}
func (gc GameConfig) getTeamAttributesv2() []enum.Attribute {
	return gc.teamAttributes
}
func (gc GameConfig) getTeamSkillsv2() []enum.SkillType {
	return gc.teamSkills
}
func (gc GameConfig) isResonantActive() bool {
	return gc.resonantOn
}
func (gc GameConfig) getCards() []*usermodel.OwnedCard {
	return gc.ocards
}
func (gc GameConfig) getLeaderIndex() int {
	return gc.leaderIndex
}
func (gc GameConfig) getGuest() *usermodel.OwnedCard {
	return gc.guest
}

// recalculate appeal, hp, teamAttributes, teamSkills, resonantOn
func (gc *GameConfig) recalculate() {
	stats, err := gc.statcalculator.Calculate(gc.bonusAppeal)
	if err != nil {
		panic(err)
	}
	gc.appeal = stats.Appeal
	gc.hp = stats.Hp
	gc.teamAttributes = stats.TeamAttributes
	gc.teamSkills = stats.TeamSkills
	gc.resonantOn = stats.IsResonantOn()
}

// NewGameConfig creates, initializes, and returns GameConfig
func NewGameConfig(
	ocards []*usermodel.OwnedCard, leaderIndex int, supports []*usermodel.OwnedCard,
	guest *usermodel.OwnedCard, song *models.Song, bonusAppeal int, calcType statcalculator.StatCalculatorType) *GameConfig {
	statCalc := statcalculator.CalculatorDispatcher(calcType)
	statCalc.SetCards(ocards)
	statCalc.SetLeaderIndex(leaderIndex)
	statCalc.SetSupports(supports)
	statCalc.SetGuest(guest)
	statCalc.SetSong(song)
	gc := GameConfig{
		ocards:         ocards,
		leaderIndex:    leaderIndex,
		supports:       supports,
		guest:          guest,
		song:           song,
		statcalculator: statCalc,
	}
	for _, ocard := range gc.ocards {
		for statType, statValue := range ocard.Stats() {
			switch statType {
			case enum.StatVisual:
				gc.baseVisual += statValue
			case enum.StatVocal:
				gc.baseVocal += statValue
			case enum.StatDance:
				gc.baseDance += statValue
			}
		}
	}
	gc.bonusAppeal = bonusAppeal
	gc.recalculate()
	return &gc
}
