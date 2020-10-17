package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// TeamLogicHandler handles logic for
type teamLogicHandler struct {
	teamLogics []*teamLogic
}

var teamInstance *teamLogicHandler

func NewTeamLogicHandler() *teamLogicHandler {
	if teamInstance == nil {
		teamInstance = makeTeamLogicHandler()
	}
	return teamInstance
}

func makeTeamLogicHandler() *teamLogicHandler {
	teamLogics := []*teamLogic{
		leadSkillIsImplemented,
		attrStatUpLeadSkillOnUnicolorTeamOnly,
		twoCardSameLeadSkillThenUseLowerID,
	}

	if helper.Features.DebugNoLogic() {
		teamLogics = []*teamLogic{}
	}
	return &teamLogicHandler{
		teamLogics: teamLogics,
	}
}

func (tlh teamLogicHandler) IsOk(team *usermodel.Team, song *models.Song) bool {
	for _, logic := range tlh.teamLogics {
		if !logic.isSatisfied(team, song) {
			return false
		}
	}
	return true
}
