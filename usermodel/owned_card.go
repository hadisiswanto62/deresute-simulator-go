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
	return fmt.Sprintf("<OwnedCard (%s); (%d); %d,%d,%d; %d,%d,%d,%d,%d>",
		oc.Card, oc.Appeal,
		oc.level, oc.skillLevel, oc.StarRank,
		oc.potVisual, oc.potDance, oc.potVocal, oc.potHp, oc.potSkill,
	)
}

func (oc OwnedCard) GetStat(stat enum.Stat) int {
	switch stat {
	case enum.StatDance:
		return oc.Dance
	case enum.StatVisual:
		return oc.Visual
	case enum.StatVocal:
		return oc.Vocal
	}
	return -1
}

func (oc OwnedCard) MaxStat() enum.Stat {
	if oc.Visual > oc.Dance && oc.Visual > oc.Vocal {
		return enum.StatVisual
	} else if oc.Vocal > oc.Dance && oc.Vocal > oc.Visual {
		return enum.StatVocal
	} else {
		return enum.StatDance
	}
}

// Level get the level of the card
func (oc *OwnedCard) Level() int {
	return oc.level
}

// SetLevel sets level of the card (and recalculate Visual, Vocal, Dance)
func (oc *OwnedCard) SetLevel(level int) {
	oc.level = level
	oc.recalculate()
}

// SkillLevel get the skill level of the card
func (oc *OwnedCard) SkillLevel() int {
	return oc.skillLevel
}

// SetSkillLevel sets skill level of the card (and recalculate skill prob/duration)
func (oc *OwnedCard) SetSkillLevel(skillLevel int) {
	oc.skillLevel = skillLevel
	oc.recalculateSkill()
}

// PotVisual gets the potential visual of the card
func (oc *OwnedCard) PotVisual() int {
	return oc.potVisual
}

// SetPotVisual sets potential visual of the card (and recalculate)
func (oc *OwnedCard) SetPotVisual(value int) {
	oc.potVisual = value
	oc.recalculate()
}

// PotDance gets the potential dance of the card
func (oc *OwnedCard) PotDance() int {
	return oc.potDance
}

// SetPotDance sets potential dance of the card (and recalculate)
func (oc *OwnedCard) SetPotDance(value int) {
	oc.potDance = value
	oc.recalculate()
}

// PotVocal gets the potential vocal of the card
func (oc *OwnedCard) PotVocal() int {
	return oc.potVocal
}

// SetPotVocal sets potential vocal of the card (and recalculate)
func (oc *OwnedCard) SetPotVocal(value int) {
	oc.potVocal = value
	oc.recalculate()
}

// PotHp gets the potential hp of the card
func (oc *OwnedCard) PotHp() int {
	return oc.potHp
}

// SetPotHp sets potential hp of the card (and recalculate)
func (oc *OwnedCard) SetPotHp(value int) {
	oc.potHp = value
	oc.recalculate()
}

// PotSkill gets the potential skill of the card
func (oc *OwnedCard) PotSkill() int {
	return oc.potSkill
}

// SetPotSkill sets potential skill of the card (and recalculate skill prob)
func (oc *OwnedCard) SetPotSkill(value int) {
	oc.potSkill = value
	oc.recalculateSkill()
}

func (oc *OwnedCard) recalculate() {
	// recalculate is responsible for checking validity of:
	// level, potVisual, potDance, potVocal, potHp, starrank
	if oc.level < 1 {
		oc.level = 1
	} else if oc.level > oc.Rarity.MaxLevel {
		oc.level = oc.Rarity.MaxLevel
	}
	if oc.potVisual < 0 {
		oc.potVisual = 0
	} else if oc.potVisual > 10 {
		oc.potVisual = 10
	}
	if oc.potDance < 0 {
		oc.potDance = 0
	} else if oc.potDance > 10 {
		oc.potDance = 10
	}
	if oc.potVocal < 0 {
		oc.potVocal = 0
	} else if oc.potVocal > 10 {
		oc.potVocal = 10
	}
	if oc.potHp < 0 {
		oc.potHp = 0
	} else if oc.potHp > 10 {
		oc.potHp = 10
	}
	// TODO: Star rank max is not always 20.
	if oc.StarRank < 1 {
		oc.StarRank = 1
	} else if oc.StarRank > 20 {
		oc.StarRank = 20
	}

	statLookup := helper.StatPotentialBonusLookupFor(oc.Rarity.Rarity)
	lifeLookup := helper.LifePotentialBonusLookupFor(oc.Rarity.Rarity)

	oc.Dance = helper.Scale(oc.Card.DanceMin, oc.Card.DanceMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusDance + statLookup[oc.potDance]
	oc.Visual = helper.Scale(oc.Card.VisualMin, oc.Card.VisualMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusVisual + statLookup[oc.potVisual]
	oc.Vocal = helper.Scale(oc.Card.VocalMin, oc.Card.VocalMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusVocal + statLookup[oc.potVocal]
	oc.Hp = helper.Scale(oc.Card.HpMin, oc.Card.HpMax, oc.Card.Rarity.MaxLevel, oc.level) + oc.Card.BonusHp + lifeLookup[oc.potHp]
	oc.Appeal = oc.Dance + oc.Visual + oc.Vocal
}

func (oc *OwnedCard) recalculateSkill() {
	// recalculateSkill is responseible for checking validity of:
	// skillLevel, potSkill
	if oc.potSkill < 0 {
		oc.potSkill = 0
	} else if oc.potSkill > 10 {
		oc.potSkill = 10
	}
	if oc.skillLevel < 1 {
		oc.skillLevel = 1
	} else if oc.skillLevel > 10 {
		oc.skillLevel = 10
	}
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
//
// Deprecated: use NewOwnedCard2 instead
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
//
// Deprecated: use NewOwnedCard2 instead
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

// OwnedCardRequest is struct that can be converted directly to OwnedCard (to prettify NewOwnedCard2 only)
type OwnedCardRequest struct {
	Card       *models.Card
	Level      int
	SkillLevel int
	StarRank   int
	PotVisual  int
	PotDance   int
	PotVocal   int
	PotHp      int
	PotSkill   int
}

// NewOwnedCard2 converts object that implements OwnedCardRequester interface to OwnedCard
func NewOwnedCard2(request OwnedCardRequest) *OwnedCard {
	level := request.Level
	skillLevel := request.SkillLevel
	starRank := request.StarRank
	if level == 0 {
		level = request.Card.Rarity.MaxLevel
	}
	if request.SkillLevel == 0 {
		skillLevel = 1
	}
	if request.StarRank == 0 {
		starRank = 1
	}
	oc := OwnedCard{
		Card:       request.Card,
		level:      level,
		skillLevel: skillLevel,
		StarRank:   starRank,
		potVisual:  request.PotVisual,
		potDance:   request.PotDance,
		potVocal:   request.PotVocal,
		potHp:      request.PotHp,
		potSkill:   request.PotSkill,
	}
	oc.recalculate()
	oc.recalculateSkill()
	return &oc
}

// OwnedCardRawData represents bare minimum of data that can uniquely define OwnedCard.
// Cannot be converted to OwnedCard directly since it does not contain the actual Card
type OwnedCardRawData struct {
	CardID     int
	SkillLevel int
	StarRank   int
	PotVisual  int
	PotDance   int
	PotVocal   int
	PotHp      int
	PotSkill   int
}
