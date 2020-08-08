package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
)

// OwnedCard represents a custom card
type OwnedCard struct {
	*Card
	level             uint8
	skillLevel        uint8
	StarRank          uint8
	potVisual         uint8
	potDance          uint8
	potVocal          uint8
	potHp             uint8
	potSkill          uint8
	Visual            uint16
	Vocal             uint16
	Dance             uint16
	Appeal            uint16
	Hp                uint16
	SkillEffectLength uint16
	SkillProcChance   uint16
}

// Level get the level of the card
func (oc *OwnedCard) Level() uint8 {
	return oc.level
}

// SetLevel sets level of the card (and recalculate Visual, Vocal, Dance)
func (oc *OwnedCard) SetLevel(level uint8) {
	if level > oc.Rarity.MaxLevel {
		level = oc.Rarity.MaxLevel
	}
	oc.level = level
	oc.recalculate()
}

// SkillLevel get the skill level of the card
func (oc *OwnedCard) SkillLevel() uint8 {
	return oc.skillLevel
}

// SetSkillLevel sets skill level of the card (and recalculate skill prob/duration)
func (oc *OwnedCard) SetSkillLevel(skillLevel uint8) {
	if skillLevel > 10 {
		skillLevel = 10
	}
	oc.skillLevel = skillLevel
	oc.recalculateSkill()
}

// PotVisual gets the potential visual of the card
func (oc *OwnedCard) PotVisual() uint8 {
	return oc.potVisual
}

// SetPotVisual sets potential visual of the card (and recalculate)
func (oc *OwnedCard) SetPotVisual(value uint8) {
	if value > 10 {
		value = 10
	}
	oc.potVisual = value
	oc.recalculate()
}

// PotDance gets the potential dance of the card
func (oc *OwnedCard) PotDance() uint8 {
	return oc.potDance
}

// SetPotDance sets potential dance of the card (and recalculate)
func (oc *OwnedCard) SetPotDance(value uint8) {
	if value > 10 {
		value = 10
	}
	oc.potDance = value
	oc.recalculate()
}

// PotVocal gets the potential vocal of the card
func (oc *OwnedCard) PotVocal() uint8 {
	return oc.potVocal
}

// SetPotVocal sets potential vocal of the card (and recalculate)
func (oc *OwnedCard) SetPotVocal(value uint8) {
	if value > 10 {
		value = 10
	}
	oc.potVocal = value
	oc.recalculate()
}

// PotHp gets the potential hp of the card
func (oc *OwnedCard) PotHp() uint8 {
	return oc.potHp
}

// SetPotHp sets potential hp of the card (and recalculate)
func (oc *OwnedCard) SetPotHp(value uint8) {
	if value > 10 {
		value = 10
	}
	oc.potHp = value
	oc.recalculate()
}

// PotSkill gets the potential skill of the card
func (oc *OwnedCard) PotSkill() uint8 {
	return oc.potSkill
}

// SetPotSkill sets potential skill of the card (and recalculate skill prob)
func (oc *OwnedCard) SetPotSkill(value uint8) {
	if value > 10 {
		value = 10
	}
	oc.potSkill = value
	oc.recalculateSkill()
}

func (oc *OwnedCard) recalculate() {
	statLookup := helper.StatPotentialBonusLookupFor(oc.Rarity.Rarity)
	lifeLookup := helper.LifePotentialBonusLookupFor(oc.Rarity.Rarity)

	oc.Dance = helper.Scale(oc.Card.DanceMin, oc.Card.DanceMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusDance + statLookup[oc.potDance]
	oc.Visual = helper.Scale(oc.Card.VisualMin, oc.Card.VisualMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusVisual + statLookup[oc.potVisual]
	oc.Vocal = helper.Scale(oc.Card.VocalMin, oc.Card.VocalMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusVocal + statLookup[oc.potVocal]
	oc.Hp = helper.Scale(oc.Card.HpMin, oc.Card.HpMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusHp + lifeLookup[oc.potHp]
	oc.Appeal = oc.Dance + oc.Visual + oc.Vocal
}

func (oc *OwnedCard) recalculateSkill() {
	procChanceLookup := helper.SkillProbPotentialBonusLookup

	oc.SkillEffectLength = helper.Scale(oc.Skill.EffectLength[0], oc.Skill.EffectLength[1], 10, oc.skillLevel)
	oc.SkillProcChance = helper.Scale(oc.Skill.ProcChance[0], oc.Skill.ProcChance[1], 10, oc.skillLevel) + procChanceLookup[oc.potSkill]
}

// New creates a new OwnedCard object with max level and skillLevel&starRank=1
func New(card *Card) OwnedCard {
	oc := OwnedCard{
		Card:       card,
		level:      card.Rarity.MaxLevel,
		skillLevel: 1,
		StarRank:   1,
	}
	oc.recalculate()
	oc.recalculateSkill()
	return oc
}
