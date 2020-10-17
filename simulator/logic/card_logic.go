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

var motifWithHighCorrectStat = &cardsLogic{
	name: "motifWithHighCorrectStat",
	isSatisfied: func(ocards [5]*usermodel.OwnedCard, song *models.Song) bool {
		threshold := 0.5
		da, vo, vi := 0, 0, 0
		for _, ocard := range ocards {
			da += ocard.Dance
			vo += ocard.Vocal
			vi += ocard.Visual
		}
		sum := float64(da + vo + vi)
		for stat, skill := range enum.MotifStatMap {
			for _, ocard := range ocards {
				if ocard.Card.Skill.SkillType.Name == skill {

					// IsSatisfied when the stat required by the motif is > 0.4*totalAppeal
					if stat == enum.StatDance && float64(da)/sum > threshold {
						return true
					} else if stat == enum.StatVocal && float64(vo)/sum > threshold {
						return true
					} else if stat == enum.StatVisual && float64(vi)/sum > threshold {
						return true
					}

					// for _, ocard := range ocards {
					// 	fmt.Printf("%d ", ocard.Card.ID)
					// }
					// fmt.Println(da, vo, vi, sum)
					return false
				}
			}
		}
		return true
	},
}
