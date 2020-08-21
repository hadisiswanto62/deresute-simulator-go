package simulator

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// GameConfig represents full game configuration (team, supports, guest, and song)
type GameConfig struct {
	team     *usermodel.Team
	supports [10]*usermodel.OwnedCard
	guest    *usermodel.OwnedCard
	song     *models.Song

	BaseVisual int
	BaseVocal  int
	BaseDance  int
	Appeal     int
	Hp         int
}

// SetGuest set guest and recalculate Appeal
func (gc *GameConfig) SetGuest(guest *usermodel.OwnedCard) {
	gc.guest = guest
	gc.recalculate()
}

// SetSong set song and recalculate Appeal
func (gc *GameConfig) SetSong(song *models.Song) {
	gc.song = song
	gc.recalculate()
}

// NewGameConfig returns pointer to a new GameConfig with calculated Appeal
func NewGameConfig(
	team *usermodel.Team, supports [10]*usermodel.OwnedCard,
	guest *usermodel.OwnedCard, song *models.Song) *GameConfig {
	gc := GameConfig{
		team:     team,
		supports: supports,
		guest:    guest,
		song:     song,
	}
	for _, ocard := range team.Ocards {
		for statType, statValue := range ocard.Stats() {
			switch statType {
			case enum.StatVisual:
				gc.BaseVisual += statValue
			case enum.StatVocal:
				gc.BaseVocal += statValue
			case enum.StatDance:
				gc.BaseDance += statValue
			}
		}
	}
	gc.recalculate()
	return &gc
}

func (gc *GameConfig) getTeamAttributes() [6]enum.Attribute {
	var ret [6]enum.Attribute
	ocards := append(gc.team.Ocards[:], gc.guest)
	for i, ocard := range ocards {
		ret[i] = ocard.Card.Idol.Attribute
	}
	return ret
}

func (gc *GameConfig) getTeamSkills() [6]enum.SkillType {
	var ret [6]enum.SkillType
	ocards := append(gc.team.Ocards[:], gc.guest)
	for i, ocard := range ocards {
		ret[i] = ocard.Card.Skill.SkillType.Name
	}
	return ret
}

// from: https://hpt.moe/deresute/Appeal_Score_Calculations
func (gc *GameConfig) recalculate() {
	appeal := 0
	hp := 0

	// Stat Appeal (team/guest) = ceiling(Base * (1 + C + (G or B) + R + T))
	ocards := append(gc.team.Ocards[:], gc.guest)
	leader := gc.team.Leader()
	teamAttributes := gc.getTeamAttributes()
	skills := gc.getTeamSkills()

	leadSkillActive := leader.Card.LeadSkill.IsActive(teamAttributes, skills)
	guestLeadSkillActive := gc.guest.Card.LeadSkill.IsActive(teamAttributes, skills)

	isResonant := gc.resonantOn()
	var resonantStat enum.Stat
	if isResonant {
		resonantStat = gc.resonantStat()
	}
	for _, ocard := range ocards {
		for statType, statValue := range ocard.Stats() {
			multiplier := 1.0
			if isResonant {
				if statType != resonantStat {
					multiplier = 0.0
				}
			}
			if leadSkillActive {
				multiplier += leader.Card.LeadSkill.StatBonus(
					leader.Card.Rarity.Rarity,
					ocard.Card.Idol.Attribute,
					statType,
					gc.song.Attribute,
				)
			}
			if guestLeadSkillActive {
				multiplier += gc.guest.Card.LeadSkill.StatBonus(
					leader.Card.Rarity.Rarity,
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
		if leadSkillActive {
			multiplier += leader.Card.LeadSkill.HpBonus(leader.Card.Rarity.Rarity, ocard.Card.Idol.Attribute)
		}
		// TODO: confirm rounding in-game (floor/round/ceil)
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
	gc.Appeal = appeal
	gc.Hp = hp

}

var resonantMap = map[enum.LeadSkill]enum.Stat{
	enum.LeadSkillResonantMakeup: enum.StatVisual,
	enum.LeadSkillResonantStep:   enum.StatDance,
	enum.LeadSkillResonantVoice:  enum.StatVocal,
}

func (gc *GameConfig) resonantStat() enum.Stat {
	for lskill, stat := range resonantMap {
		if gc.team.Leader().LeadSkill.Name == lskill {
			return stat
		}
	}
	panic("is not resonant leader")
}

func (gc *GameConfig) resonantOn() bool {
	found := false
	for lskill := range resonantMap {
		if gc.team.Leader().LeadSkill.Name == lskill {
			found = true
		}
	}
	if found == false {
		return false
	}
	ocards := append(gc.team.Ocards[:], gc.guest)
	attrs := [6]enum.Attribute{}
	skills := [6]enum.SkillType{}
	for i, ocard := range ocards {
		attrs[i] = ocard.Card.Idol.Attribute
		skills[i] = ocard.Card.Skill.SkillType.Name
	}
	return gc.team.Leader().LeadSkill.IsActive(attrs, skills)
}
