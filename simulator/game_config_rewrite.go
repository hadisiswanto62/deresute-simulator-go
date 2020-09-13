package simulator

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type GameConfig2 struct {
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

func (gc GameConfig2) getSkillActivableCards() []*usermodel.OwnedCard {
	return gc.ocards
}
func (gc GameConfig2) getLeadSkillActivableCards() []*usermodel.OwnedCard {
	return []*usermodel.OwnedCard{
		gc.ocards[gc.leaderIndex],
		gc.guest,
	}
}
func (gc GameConfig2) getSong() *models.Song {
	return gc.song
}
func (gc GameConfig2) getBaseVisual() int {
	return gc.baseVisual
}
func (gc GameConfig2) getBaseVocal() int {
	return gc.baseVocal
}
func (gc GameConfig2) getBaseDance() int {
	return gc.baseDance
}
func (gc GameConfig2) getAppeal() int {
	return gc.appeal
}
func (gc GameConfig2) getHp() int {
	return gc.hp
}
func (gc GameConfig2) getTeamAttributesv2() []enum.Attribute {
	return gc.teamAttributes
}
func (gc GameConfig2) getTeamSkillsv2() []enum.SkillType {
	return gc.teamSkills
}
func (gc GameConfig2) isResonantActive() bool {
	return gc.resonantOn
}
func (gc GameConfig2) getCards() []*usermodel.OwnedCard {
	return gc.ocards
}
func (gc GameConfig2) getLeaderIndex() int {
	return gc.leaderIndex
}
func (gc GameConfig2) getGuest() *usermodel.OwnedCard {
	return gc.guest
}

// recalculate appeal, hp, teamAttributes, teamSkills, resonantOn
func (gc *GameConfig2) recalculate() {
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

func NewGameConfig2(
	ocards []*usermodel.OwnedCard, leaderIndex int, supports []*usermodel.OwnedCard,
	guest *usermodel.OwnedCard, song *models.Song) *GameConfig2 {
	gc := GameConfig2{
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
