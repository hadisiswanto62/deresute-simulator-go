package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogic struct {
	name       string
	isViolated func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool
}

var unisonInCorrectSongType = &gameConfigLogic{
	name: "unisonInCorrectSongType",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {

		for _, ocards := range leadSkillActivableCards {
			for attr, lskill := range enum.LeadSkillUnisonMap {
				if lskill != ocards.LeadSkill.Name {
					continue
				}
				// if use unison:
				// violated if song is ALL (princess is objectively better)
				if song.Attribute == enum.AttrAll {
					return true
				}
				// violated if unison attr does not match song attr
				if attr != song.Attribute {
					return true
				}
			}
		}
		return false
	},
}

var allLeadSkillsActive = &gameConfigLogic{
	name: "allLeadSkillsActive",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		attrs := []enum.Attribute{guest.Card.Idol.Attribute}
		skills := []enum.SkillType{guest.Card.Skill.SkillType.Name}
		for _, ocard := range team.Ocards {
			attrs = append(attrs, ocard.Card.Idol.Attribute)
			skills = append(skills, ocard.Card.Skill.SkillType.Name)
		}
		for _, ocard := range leadSkillActivableCards {
			// violated if any lead skill is inactive
			if !ocard.LeadSkill.IsActive(attrs, skills) {
				return true
			}
		}
		return false
	},
}
