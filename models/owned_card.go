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
}

// Level get the level of the card
func (oc *OwnedCard) Level() int8 {
	return oc.level
}

// SetLevel sets level of the card (and recalculate Visual, Vocal, Dance)
func (oc *OwnedCard) SetLevel(level int8) {
	oc.level = level
	oc.recalculate()
}

// SkillLevel get the skill level of the card
func (oc *OwnedCard) SkillLevel() int8 {
	return oc.skillLevel
}

// SetSkillLevel sets skill level of the card (and recalculate skill prob/duration NOT IMPLEMENTED)
func (oc *OwnedCard) SetSkillLevel(skillLevel int8) {
	oc.skillLevel = skillLevel
}

func (oc *OwnedCard) recalculate() {
	oc.Dance = helper.Scale(oc.Card.DanceMin, oc.Card.DanceMax, oc.Card.MaxLevel, oc.level)
	oc.Visual = helper.Scale(oc.Card.VisualMin, oc.Card.VisualMax, oc.Card.MaxLevel, oc.level)
	oc.Vocal = helper.Scale(oc.Card.VocalMin, oc.Card.VocalMax, oc.Card.MaxLevel, oc.level)
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
