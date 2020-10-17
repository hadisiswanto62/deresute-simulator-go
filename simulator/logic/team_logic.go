package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type teamLogic struct {
	name        string
	isSatisfied func(team *usermodel.Team, song *models.Song) bool
}

var leadSkillIsImplemented = teamLogic{
	name: "leadSkillIsImplemented",
	isSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		// IsSatisfied when lead skill is implemented
		return helper.IsLeadSkillImplemented(team.Leader().Card.LeadSkill.Name)
	},
}
