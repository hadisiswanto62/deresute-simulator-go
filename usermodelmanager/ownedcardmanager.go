package usermodelmanager

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
)

type dataParser interface {
	ParseOwnedCardRawData(path string) ([]*usermodel.OwnedCardRawData, error)
}

func ParseOwnedCard(dp dataParser, path string) ([]*usermodel.OwnedCard, error) {
	cm, _ := cardmanager.Default()

	var ocards []*usermodel.OwnedCard
	ocds, err := dp.ParseOwnedCardRawData(path)
	if err != nil {
		return nil, fmt.Errorf("cannot parse owned cards: %v", err)
	}
	for _, ocd := range ocds {
		card := cm.Filter().ID(ocd.CardID).First()
		request := usermodel.OwnedCardRequest{
			Card:       card,
			SkillLevel: ocd.SkillLevel,
			StarRank:   ocd.StarRank,
			PotVisual:  ocd.PotVisual,
			PotDance:   ocd.PotDance,
			PotVocal:   ocd.PotVocal,
			PotHp:      ocd.PotHp,
			PotSkill:   ocd.PotSkill,
		}
		ocard := usermodel.NewOwnedCard2(request)
		ocards = append(ocards, ocard)
	}
	return ocards, nil
}
