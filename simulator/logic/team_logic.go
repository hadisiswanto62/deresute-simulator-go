package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type teamLogic struct {
	name        string
	isSatisfied func(team *usermodel.Team, song *models.Song) bool
}

var leadSkillIsImplemented = &teamLogic{
	name: "leadSkillIsImplemented",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		// IsSatisfied when lead skill is implemented
		return helper.IsLeadSkillImplemented(team.Leader().Card.LeadSkill.Name)
	},
}

var attrStatUpLeadSkillOnUnicolorTeamOnly = &teamLogic{
	name: "attrStatUpLeadSkillOnUnicolorTeamOnly",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		for attr, statLeadSkillMap := range enum.LeadSkillAttrStatUpMap {
			for _, lskill := range statLeadSkillMap {
				if team.Leader().LeadSkill.Name == lskill {
					for _, ocard := range team.Ocards {
						if ocard.Card.Idol.Attribute != attr {
							return false
						}
					}
					return true
				}
			}
		}
		return true
	},
}

var twoCardSameLeadSkillThenUseLowerID = &teamLogic{
	name: "twoCardSameLeadSkillThenUseLowerID",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		// leadSkillMap maps lead skill to THE LOWEST ID of card with that lead skill in the team
		leadSkillMap := make(map[enum.LeadSkill]int, 0)
		for _, ocard := range team.Ocards {
			lskill := ocard.LeadSkill.Name
			_, ok := leadSkillMap[lskill]
			if !ok {
				leadSkillMap[lskill] = 9999999
			}
			if ocard.Card.ID < leadSkillMap[lskill] {
				leadSkillMap[lskill] = ocard.Card.ID
			}
		}
		activeLeadSkill := team.Leader().Card.LeadSkill.Name
		leadID := team.Leader().Card.ID
		for lskill, ID := range leadSkillMap {
			if lskill == activeLeadSkill && ID != leadID {
				return false
			}
		}
		return true
	},
}
