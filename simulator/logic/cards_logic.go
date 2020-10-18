package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type cardsLogic struct {
	name        string
	isSatisfied func(cards [5]*usermodel.OwnedCard, song *models.Song) bool
}

var motifWithHighCorrectStat = &cardsLogic{
	name: "motifWithHighCorrectStat",
	isSatisfied: func(ocards [5]*usermodel.OwnedCard, song *models.Song) bool {
		threshold := 0.4
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
					// for _, ocard := range ocards {
					// 	fmt.Printf("%d ", ocard.Card.ID)
					// }
					// fmt.Println(da, vo, vi, sum)
					if stat == enum.StatDance && float64(da)/sum > threshold {
						continue
					} else if stat == enum.StatVocal && float64(vo)/sum > threshold {
						continue
					} else if stat == enum.StatVisual && float64(vi)/sum > threshold {
						continue
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

var useUnevolvedWithoutEvolved = &cardsLogic{
	name: "useUnevolvedWithoutEvolved",
	isSatisfied: func(ocards [5]*usermodel.OwnedCard, song *models.Song) bool {
		for _, ocard := range ocards {
			if ocard.Card.Rarity.IsEvolved {
				continue
			}
			foundEvolved := false
			for _, ocard2 := range ocards {
				if ocard2.Card.ID == ocard.Card.ID+1 {
					foundEvolved = true
				}
			}
			if !foundEvolved {
				return false
			}
		}
		return true
	},
}
