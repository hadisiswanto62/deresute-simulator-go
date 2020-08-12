package usermodel

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

// OwnedCard represents a custom card
type OwnedCard struct {
	*models.Card
	level             int
	skillLevel        int
	StarRank          int
	potVisual         int
	potDance          int
	potVocal          int
	potHp             int
	potSkill          int
	Visual            int
	Vocal             int
	Dance             int
	Appeal            int
	Hp                int
	SkillEffectLength int
	SkillProcChance   int
}

func (oc OwnedCard) String() string {
	return fmt.Sprintf("<OwnedCard (%s); %d,%d,%d; %d,%d,%d,%d,%d",
		oc.Card, oc.level, oc.skillLevel, oc.StarRank,
		oc.potVisual, oc.potDance, oc.potVocal, oc.potHp, oc.potSkill,
	)
}

// Level get the level of the card
func (oc *OwnedCard) Level() int {
	return oc.level
}

// SetLevel sets level of the card (and recalculate Visual, Vocal, Dance)
func (oc *OwnedCard) SetLevel(level int) {
	if level > oc.Rarity.MaxLevel {
		level = oc.Rarity.MaxLevel
	}
	oc.level = level
	oc.recalculate()
}

// SkillLevel get the skill level of the card
func (oc *OwnedCard) SkillLevel() int {
	return oc.skillLevel
}

// SetSkillLevel sets skill level of the card (and recalculate skill prob/duration)
func (oc *OwnedCard) SetSkillLevel(skillLevel int) {
	if skillLevel > 10 {
		skillLevel = 10
	}
	oc.skillLevel = skillLevel
	oc.recalculateSkill()
}

// PotVisual gets the potential visual of the card
func (oc *OwnedCard) PotVisual() int {
	return oc.potVisual
}

// SetPotVisual sets potential visual of the card (and recalculate)
func (oc *OwnedCard) SetPotVisual(value int) {
	if value > 10 {
		value = 10
	}
	oc.potVisual = value
	oc.recalculate()
}

// PotDance gets the potential dance of the card
func (oc *OwnedCard) PotDance() int {
	return oc.potDance
}

// SetPotDance sets potential dance of the card (and recalculate)
func (oc *OwnedCard) SetPotDance(value int) {
	if value > 10 {
		value = 10
	}
	oc.potDance = value
	oc.recalculate()
}

// PotVocal gets the potential vocal of the card
func (oc *OwnedCard) PotVocal() int {
	return oc.potVocal
}

// SetPotVocal sets potential vocal of the card (and recalculate)
func (oc *OwnedCard) SetPotVocal(value int) {
	if value > 10 {
		value = 10
	}
	oc.potVocal = value
	oc.recalculate()
}

// PotHp gets the potential hp of the card
func (oc *OwnedCard) PotHp() int {
	return oc.potHp
}

// SetPotHp sets potential hp of the card (and recalculate)
func (oc *OwnedCard) SetPotHp(value int) {
	if value > 10 {
		value = 10
	}
	oc.potHp = value
	oc.recalculate()
}

// PotSkill gets the potential skill of the card
func (oc *OwnedCard) PotSkill() int {
	return oc.potSkill
}

// SetPotSkill sets potential skill of the card (and recalculate skill prob)
func (oc *OwnedCard) SetPotSkill(value int) {
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

// Stats returns the ocard's stats in the form of map
func (oc OwnedCard) Stats() map[enum.Stat]int {
	return map[enum.Stat]int{
		enum.StatVisual: oc.Visual,
		enum.StatVocal:  oc.Vocal,
		enum.StatDance:  oc.Dance,
	}
}

// NewCustomOwnedCard creates a new OwnedCard object with max level and specified skillLevel and starRank
func NewCustomOwnedCard(card *models.Card, skillLevel, starRank,
	potVisual, potDance, potVocal, potHp, potSkill int) *OwnedCard {
	oc := OwnedCard{
		Card:       card,
		level:      card.Rarity.MaxLevel,
		skillLevel: skillLevel,
		StarRank:   starRank,
		potHp:      potHp,
		potVisual:  potVisual,
		potDance:   potDance,
		potVocal:   potVocal,
		potSkill:   potSkill,
	}
	oc.recalculate()
	oc.recalculateSkill()
	return &oc
}

// NewOwnedCard creates a new OwnedCard object with max level and skillLevel&starRank=1
func NewOwnedCard(card *models.Card) *OwnedCard {
	return NewCustomOwnedCard(
		card, 1, 1, 0, 0, 0, 0, 0,
	)
}

// BatchNewOwnedCards creates slice of OwnedCards object (from NewCustomOwnedCard)
func BatchNewOwnedCards(cards []*models.Card, skillLevel, starRank,
	potVisual, potDance, potVocal, potHp, potSkill int) []*OwnedCard {
	var ret []*OwnedCard
	for _, card := range cards {
		ret = append(ret, NewCustomOwnedCard(
			card, skillLevel, starRank,
			potVisual, potDance, potVocal, potHp, potSkill))
	}
	return ret
}
