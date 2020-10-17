package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogic struct {
	name       string
	isViolated func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, song *models.Song) bool
}

var unisonInCorrectSongType = gameConfigLogic{
	name: "unisonInCorrectSongType",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, song *models.Song) bool {
		// violated if song is ALL (princess is objectively better)
		if song.Attribute == enum.AttrAll {
			return true
		}

		for _, ocards := range leadSkillActivableCards {
			for attr, lskill := range enum.UnisonMap {
				if lskill != ocards.LeadSkill.Name {
					continue
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
