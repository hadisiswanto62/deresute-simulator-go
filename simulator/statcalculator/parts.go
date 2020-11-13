package statcalculator

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

func populateTeamAttributesSkills(data *GameConfigStats, ocards []*usermodel.OwnedCard) {
	attrs := make([]enum.Attribute, 0, len(ocards))
	skills := make([]enum.SkillType, 0, len(ocards))
	for _, ocard := range ocards {
		if ocard == nil {
			continue
		}
		attrs = append(attrs, ocard.Card.Idol.Attribute)
		skills = append(skills, ocard.Card.Skill.SkillType.Name)
	}
	data.TeamAttributes = attrs
	data.TeamSkills = skills
}

func populateResonant(data *GameConfigStats, leadSkillActivableCards []*usermodel.OwnedCard) {
	for _, ocard := range leadSkillActivableCards {
		if ocard == nil {
			continue
		}
		leadSkill := ocard.Card.LeadSkill
		if !leadSkill.IsActive(data.TeamAttributes, data.TeamSkills) {
			continue
		}
		for stat, reso := range enum.LeadSkillResonantMap {
			if leadSkill.Name == reso {
				data.resonantStat = stat
				return
			}
		}
	}
}

func populateAppealHp(data *GameConfigStats, ocards, leadSkillActivableCards []*usermodel.OwnedCard, song *models.Song) {
	appeal := 0
	hp := 0
	for _, ocard := range ocards {
		if ocard == nil {
			continue
		}
		for statType, statValue := range ocard.Stats() {
			multiplier := 1.0
			if data.IsResonantOn() && statType != data.resonantStat {
				multiplier = 0.0
			}
			for _, leadOcard := range leadSkillActivableCards {
				if leadOcard == nil {
					continue
				}
				leadSkill := leadOcard.Card.LeadSkill
				if !leadSkill.IsActive(data.TeamAttributes, data.TeamSkills) {
					continue
				}
				multiplier += leadSkill.StatBonus(
					leadOcard.Card.Rarity.Rarity,
					ocard.Card.Idol.Attribute,
					statType,
					song.Attribute,
				)
			}
			multiplier += helper.GetRoomItemBonus(ocard.Card.Idol.Attribute)
			if (ocard.Card.Idol.Attribute == song.Attribute) || (song.Attribute == enum.AttrAll) {
				multiplier += 0.3
			}
			cardAppeal := int(math.Ceil(multiplier*float64(statValue) - 0.000001))
			appeal += cardAppeal
		}
		multiplier := 1.0
		for _, leadOcard := range leadSkillActivableCards {
			if leadOcard == nil {
				continue
			}
			leadSkill := leadOcard.Card.LeadSkill
			if !leadSkill.IsActive(data.TeamAttributes, data.TeamSkills) {
				continue
			}
			multiplier += leadOcard.Card.LeadSkill.HpBonus(leadOcard.Card.Rarity.Rarity, ocard.Card.Idol.Attribute)
		}
		hp += int(multiplier * float64(ocard.Hp))
	}
	data.Appeal += appeal
	data.Hp += hp
}

func calcSupportAppeals(data *GameConfigStats, supports []*usermodel.OwnedCard) int {
	appeal := 0
	for _, ocard := range supports {
		if ocard == nil {
			continue
		}
		for _, statValue := range ocard.Stats() {
			// Supports attribute always matches song
			multiplier := 1.3
			appeal += int(math.Ceil(multiplier * float64(statValue) * 0.5))
		}
	}
	return appeal
}
