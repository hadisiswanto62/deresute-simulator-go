package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogicHandler struct {
	gameConfigLogics []*gameConfigLogic
}

var gameConfigInstance *gameConfigLogicHandler

// NewGameConfigLogicHandler creates a new gameConfigLogicHandler instance
func NewGameConfigLogicHandler() *gameConfigLogicHandler {
	if gameConfigInstance == nil {
		gameConfigInstance = makeGameConfigLogicHandler()
	}
	return gameConfigInstance
}

func makeGameConfigLogicHandler() *gameConfigLogicHandler {
	gameConfigLogics := []*gameConfigLogic{
		unisonInCorrectSongType,
		allLeadSkillsActive,
		princessWhenShouldBeUnison,
		skillsAreActive,          // TODO: need test
		guestTriColorCorrectStat, // TODO: need test
		guestPrincessUnisonCorrectStat,
		tricolorCorrectColor, // TODO: investigate why this one cause some team to have 0 guest
	}

	if helper.Features.DebugNoLogic() {
		gameConfigLogics = []*gameConfigLogic{}
	}
	return &gameConfigLogicHandler{
		gameConfigLogics: gameConfigLogics,
	}
}

func (gclh gameConfigLogicHandler) IsOk(team *usermodel.Team, guest *usermodel.OwnedCard, song *models.Song) bool {
	leadSkillActivable := []*usermodel.OwnedCard{team.Leader(), guest}
	for _, logic := range gclh.gameConfigLogics {
		if logic.isViolated(team, leadSkillActivable, guest, song) {
			// fmt.Println(logic.name)
			return false
		}
	}
	return true
}
