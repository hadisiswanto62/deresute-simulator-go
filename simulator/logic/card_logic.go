package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type cardLogic struct {
	name        string
	isSatisfied func(ocard *usermodel.OwnedCard, song *models.Song) bool
}

type cardsLogic struct {
	name        string
	isSatisfied func(cards [5]*usermodel.OwnedCard, song *models.Song) bool
}

var cardCardIsSSR = &cardLogic{
	name: "cardCardIsSSR",
	isSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		return ocard.Card.Rarity.Rarity == enum.RaritySSR
	},
}

var cardSkillIsNotConcentration = &cardLogic{
	name: "cardSkillIsNotConcentration",
	isSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		return ocard.Card.Skill.SkillType.Name != enum.SkillTypeConcentration
	},
}

var cardSkillIsImplemented = &cardLogic{
	name: "cardSkillIsImplemented",
	isSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		return helper.IsSkillImplemented(ocard.Card.Skill.SkillType.Name)
	},
}
