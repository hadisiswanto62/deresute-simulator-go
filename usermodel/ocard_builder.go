package usermodel

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

type OwnedCardBuilder interface {
	Card(*models.Card) OwnedCardBuilder
	Level(int) OwnedCardBuilder
	SkillLevel(int) OwnedCardBuilder
	StarRank(int) OwnedCardBuilder
	PotVisual(int) OwnedCardBuilder
	PotDance(int) OwnedCardBuilder
	PotVocal(int) OwnedCardBuilder
	PotHp(int) OwnedCardBuilder
	PotSkill(int) OwnedCardBuilder
	Build() (*OwnedCard, error)
}

type ownedCardBuilder struct {
	card       *models.Card
	level      int // default is card's max
	skillLevel int // default is 10
	starRank   int // default is 1
	potVisual  int // default is 0
	potDance   int // default is 0
	potVocal   int // default is 0
	potHp      int // default is 0
	potSkill   int // default is 0
}

func (ocb *ownedCardBuilder) Card(card *models.Card) OwnedCardBuilder {
	ocb.card = card
	return ocb
}

func (ocb *ownedCardBuilder) Level(level int) OwnedCardBuilder {
	ocb.level = level
	return ocb
}
func (ocb *ownedCardBuilder) SkillLevel(skillLevel int) OwnedCardBuilder {
	ocb.skillLevel = skillLevel
	return ocb
}

func (ocb *ownedCardBuilder) StarRank(starRank int) OwnedCardBuilder {
	ocb.starRank = starRank
	return ocb
}

func (ocb *ownedCardBuilder) PotVisual(potVisual int) OwnedCardBuilder {
	ocb.potVisual = potVisual
	return ocb
}
func (ocb *ownedCardBuilder) PotDance(potDance int) OwnedCardBuilder {
	ocb.potDance = potDance
	return ocb
}
func (ocb *ownedCardBuilder) PotVocal(potVocal int) OwnedCardBuilder {
	ocb.potVocal = potVocal
	return ocb
}

func (ocb *ownedCardBuilder) PotHp(potHp int) OwnedCardBuilder {
	ocb.potHp = potHp
	return ocb
}
func (ocb *ownedCardBuilder) PotSkill(potSkill int) OwnedCardBuilder {
	ocb.potSkill = potSkill
	return ocb
}

func (ocb *ownedCardBuilder) Build() (*OwnedCard, error) {
	if ocb.card == nil {
		return nil, fmt.Errorf("Card is not defined")
	}
	if ocb.level == -1 {
		ocb.level = ocb.card.Rarity.MaxLevel
	}
	oc := OwnedCard{
		Card:       ocb.card,
		level:      ocb.level,
		skillLevel: ocb.skillLevel,
		StarRank:   ocb.starRank,
		potVisual:  ocb.potVisual,
		potDance:   ocb.potDance,
		potVocal:   ocb.potVocal,
		potHp:      ocb.potHp,
		potSkill:   ocb.potSkill,
	}
	oc.recalculate()
	oc.recalculateSkill()
	return &oc, nil
}

var _ ownedCardBuilder = ownedCardBuilder{}

// NewOwnedCardBuilder returns a new owned card builder
func NewOwnedCardBuilder() OwnedCardBuilder {
	ocb := ownedCardBuilder{}
	ocb.level = -1
	ocb.skillLevel = 10
	ocb.starRank = 1
	return &ocb
}
