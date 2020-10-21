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

var tricolorMin2Color = &teamLogic{
	name: "tricolorMin2Color",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		mapp := enum.LeadSkillTricolorMap
		mapp["tmp"] = enum.LeadSkillTricolorAbility
		for _, lskill := range mapp {
			if team.Leader().LeadSkill.Name == lskill {
				attributes := make(map[enum.Attribute]bool)
				for _, ocard := range team.Ocards {
					attributes[ocard.Card.Idol.Attribute] = true
				}
				return len(attributes) >= 2
			}
		}
		return true
	},
}

var tricolorMin3Color = &teamLogic{
	name: "tricolorMin3Color",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		mapp := enum.LeadSkillTricolorMap
		mapp["tmp"] = enum.LeadSkillTricolorAbility
		for _, lskill := range mapp {
			if team.Leader().LeadSkill.Name == lskill {
				attributes := make(map[enum.Attribute]bool)
				for _, ocard := range team.Ocards {
					attributes[ocard.Card.Idol.Attribute] = true
				}
				return len(attributes) == 3
			}
		}
		return true
	},
}
var cardsResoOn3UniqueSkills = &teamLogic{
	name: "cardsResoOn3UniqueSkills",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		skills := make(map[enum.SkillType]bool)
		for _, ocard := range team.Ocards {
			skills[ocard.Card.Skill.SkillType.Name] = true
		}
		if len(skills) > 3 {
			return true
		}
		for _, lskill := range enum.LeadSkillResonantMap {
			if team.Leader().Card.LeadSkill.Name == lskill {
				return false
			}
		}
		return true
	},
}
