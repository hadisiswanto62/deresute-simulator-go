package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
)

// OwnedCard represents a custom card
type OwnedCard struct {
	*Card
	level      int8
	skillLevel int8
	StarRank   int8
	potVisual  int8
	potDance   int8
	potVocal   int8
	potHp      int8
	potSkill   int8
	Visual     uint16
	Vocal      uint16
	Dance      uint16
	Appeal     uint16
	Hp         uint16
}

// Level get the level of the card
func (oc *OwnedCard) Level() int8 {
	return oc.level
}

// SetLevel sets level of the card (and recalculate Visual, Vocal, Dance)
func (oc *OwnedCard) SetLevel(level int8) {
	if level > oc.MaxLevel {
		level = oc.MaxLevel
	}
	oc.level = level
	oc.recalculate()
}

// SkillLevel get the skill level of the card
func (oc *OwnedCard) SkillLevel() int8 {
	return oc.skillLevel
}

// SetSkillLevel sets skill level of the card (and recalculate skill prob/duration NOT IMPLEMENTED)
func (oc *OwnedCard) SetSkillLevel(skillLevel int8) {
	if skillLevel > 10 {
		skillLevel = 10
	}
	oc.skillLevel = skillLevel
}

// PotVisual gets the potential visual of the card
func (oc *OwnedCard) PotVisual() int8 {
	return oc.potVisual
}

// SetPotVisual sets potential visual of the card (and recalculate)
func (oc *OwnedCard) SetPotVisual(value int8) {
	if value > 10 {
		value = 10
	}
	oc.potVisual = value
	oc.recalculate()
}

// PotDance gets the potential dance of the card
func (oc *OwnedCard) PotDance() int8 {
	return oc.potDance
}

// SetPotDance sets potential dance of the card (and recalculate)
func (oc *OwnedCard) SetPotDance(value int8) {
	if value > 10 {
		value = 10
	}
	oc.potDance = value
	oc.recalculate()
}

// PotVocal gets the potential vocal of the card
func (oc *OwnedCard) PotVocal() int8 {
	return oc.potVocal
}

// SetPotVocal sets potential vocal of the card (and recalculate)
func (oc *OwnedCard) SetPotVocal(value int8) {
	if value > 10 {
		value = 10
	}
	oc.potVocal = value
	oc.recalculate()
}

// PotHp gets the potential hp of the card
func (oc *OwnedCard) PotHp() int8 {
	return oc.potHp
}

// SetPotHp sets potential hp of the card (and recalculate)
func (oc *OwnedCard) SetPotHp(value int8) {
	if value > 10 {
		value = 10
	}
	oc.potHp = value
	oc.recalculate()
}

// PotSkill gets the potential skill of the card
func (oc *OwnedCard) PotSkill() int8 {
	return oc.potSkill
}

// SetPotSkill sets potential skill of the card (and recalculate skill prob NOT IMPLEMENTED)
func (oc *OwnedCard) SetPotSkill(value int8) {
	if value > 10 {
		value = 10
	}
	oc.potSkill = value
}

func (oc *OwnedCard) recalculate() {
	statLookup := helper.StatPotentialBonusLookupFor(oc.Rarity)
	lifeLookup := helper.LifePotentialBonusLookupFor(oc.Rarity)

	oc.Dance = helper.Scale(oc.Card.DanceMin, oc.Card.DanceMax, oc.Card.MaxLevel, oc.level) + oc.Card.BonusDance + statLookup[oc.potDance]
	oc.Visual = helper.Scale(oc.Card.VisualMin, oc.Card.VisualMax, oc.Card.MaxLevel, oc.level) + oc.Card.BonusVisual + statLookup[oc.potVisual]
	oc.Vocal = helper.Scale(oc.Card.VocalMin, oc.Card.VocalMax, oc.Card.MaxLevel, oc.level) + oc.Card.BonusVocal + statLookup[oc.potVocal]
	oc.Hp = helper.Scale(oc.Card.HpMin, oc.Card.HpMax, oc.Card.MaxLevel, oc.level) + oc.Card.BonusHp + lifeLookup[oc.potHp]
	oc.Appeal = oc.Dance + oc.Visual + oc.Vocal
}

// New creates a new OwnedCard object with max level and skillLevel&starRank=1
func New(card *Card) OwnedCard {
	oc := OwnedCard{
		Card:       card,
		level:      card.MaxLevel,
		skillLevel: 1,
		StarRank:   1,
	}
	oc.recalculate()
	return oc
}
