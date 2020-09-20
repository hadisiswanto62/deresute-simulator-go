package simulator

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
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
}

func (gc GameConfig) GetOcards() []*usermodel.OwnedCard {
	return gc.ocards
}

func (gc GameConfig) Recalculate() {
	gc.recalculate()
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
	appeal := 0
	hp := 0

	// Stat Appeal (team/guest) = ceiling(Base * (1 + C + (G or B) + R + T))
	ocards := append(gc.ocards, gc.guest)
	leader := gc.ocards[gc.leaderIndex]

	// teamAttributes and teamSkills
	gc.teamAttributes = make([]enum.Attribute, 0, len(ocards))
	gc.teamSkills = make([]enum.SkillType, 0, len(ocards))
	for _, ocard := range ocards {
		gc.teamAttributes = append(gc.teamAttributes, ocard.Card.Idol.Attribute)
		gc.teamSkills = append(gc.teamSkills, ocard.Card.Skill.SkillType.Name)
	}

	// resonantOn
	gc.resonantOn = false
	var resonantStat enum.Stat
	leadSkillsActive := map[*usermodel.OwnedCard]bool{
		leader:   leader.Card.LeadSkill.IsActive(gc.teamAttributes, gc.teamSkills),
		gc.guest: gc.guest.Card.LeadSkill.IsActive(gc.teamAttributes, gc.teamSkills),
	}
	for ocard, active := range leadSkillsActive {
		if !active {
			continue
		}
		leadSkill := ocard.Card.LeadSkill
		for stat, reso := range enum.ResonantMap {
			if leadSkill.Name == reso {
				gc.resonantOn = true
				resonantStat = stat
				break
			}
		}
	}

	// appeal and hp
	for _, ocard := range ocards {
		for statType, statValue := range ocard.Stats() {
			multiplier := 1.0
			if gc.resonantOn && statType != resonantStat {
				multiplier = 0.0
			}
			for leadOcard, active := range leadSkillsActive {
				if !active {
					continue
				}
				multiplier += leadOcard.Card.LeadSkill.StatBonus(
					leadOcard.Card.Rarity.Rarity,
					ocard.Card.Idol.Attribute,
					statType,
					gc.song.Attribute,
				)
			}
			multiplier += helper.GetRoomItemBonus(ocard.Card.Idol.Attribute)
			if (ocard.Card.Idol.Attribute == gc.song.Attribute) || (gc.song.Attribute == enum.AttrAll) {
				multiplier += 0.3
			}
			appeal += int(math.Ceil(multiplier * float64(statValue)))
		}
		multiplier := 1.0
		for leadOcard, active := range leadSkillsActive {
			if !active {
				continue
			}
			multiplier += leadOcard.Card.LeadSkill.HpBonus(leadOcard.Card.Rarity.Rarity, ocard.Card.Idol.Attribute)
		}
		hp += int(multiplier * float64(ocard.Hp))
	}
	for _, ocard := range gc.supports {
		for _, statValue := range ocard.Stats() {
			multiplier := 1.0
			if (ocard.Card.Idol.Attribute == gc.song.Attribute) || (gc.song.Attribute == enum.AttrAll) {
				multiplier += 0.3
			}
			appeal += int(math.Ceil(multiplier * float64(statValue) * 0.5))
		}
	}
	gc.appeal = appeal
	gc.hp = hp
}

// NewGameConfig creates, initializes, and returns GameConfig
func NewGameConfig(
	ocards []*usermodel.OwnedCard, leaderIndex int, supports []*usermodel.OwnedCard,
	guest *usermodel.OwnedCard, song *models.Song) *GameConfig {
	gc := GameConfig{
		ocards:      ocards,
		leaderIndex: leaderIndex,
		supports:    supports,
		guest:       guest,
		song:        song,
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
	gc.recalculate()
	return &gc
}
