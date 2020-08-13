package models

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/config"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// SkillType represents a card's skill
type SkillType struct {
	Name       enum.SkillType
	IsActive   func(attr [6]enum.Attribute) bool
	ComboBonus func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteType enum.NoteType) float64
	ScoreBonus func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteType enum.NoteType) float64
	TapHeal    func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int
}

var skillMap = map[string]*SkillType{
	string(enum.SkillTypeScoreBonus):       &SkillTypeScoreBonus,
	string(enum.SkillTypeComboBonus):       &SkillTypeComboBonus,
	string(enum.SkillTypeConcentration):    &SkillTypeConcentration,
	string(enum.SkillTypeHealer):           &SkillTypeHealer,
	string(enum.SkillTypeAllRound):         &SkillTypeAllRound,
	string(enum.SkillTypeCoordinate):       &SkillTypeCoordinate,
	string(enum.SkillTypeOverload):         &SkillTypeOverload,
	string(enum.SkillTypeTricolorSynergy):  &SkillTypeTricolorSynergy,
	string(enum.SkillTypeTricolorSymphony): &SkillTypeTricolorSymphony,
	string(enum.SkillTypeTuning):           &SkillTypeTuning,
	string(enum.SkillTypePerfectLock):      &SkillTypePerfectLock,
	string(enum.SkillTypeComboGuard):       &SkillTypeComboGuard,
	string(enum.SkillTypeLifeSparkle):      &SkillTypeLifeSparkle,
	string(enum.SkillTypeLifeGuard):        &SkillTypeLifeGuard,
	string(enum.SkillTypeSkillBoost):       &SkillTypeSkillBoost,
	string(enum.SkillTypeCuteFocus):        &SkillTypeCuteFocus,
	string(enum.SkillTypeCoolFocus):        &SkillTypeCoolFocus,
	string(enum.SkillTypePassionFocus):     &SkillTypePassionFocus,
	string(enum.SkillTypeEncore):           &SkillTypeEncore,
	string(enum.SkillTypeAlternate):        &SkillTypeAlternate,
	string(enum.SkillTypeVisualMotif):      &SkillTypeVisualMotif,
	string(enum.SkillTypeDanceMotif):       &SkillTypeDanceMotif,
	string(enum.SkillTypeVocalMotif):       &SkillTypeVocalMotif,
	string(enum.SkillTypeCuteEnsemble):     &SkillTypeCuteEnsemble,
	string(enum.SkillTypeCoolEnsemble):     &SkillTypeCoolEnsemble,
	string(enum.SkillTypePassionEnsemble):  &SkillTypePassionEnsemble,
}

// GetSkillType returns pointer to skill type with the requested name
func GetSkillType(name string) (*SkillType, error) {
	skill, ok := skillMap[name]
	if ok {
		return skill, nil
	}
	// special case
	if name == "Perfect Score Bonus" {
		return &SkillTypeScoreBonus, nil
	}
	if (name == "Lesser Perfect Lock") || (name == "Greater Perfect Lock") || (name == "Extreme Perfect Lock") {
		return &SkillTypePerfectLock, nil
	}

	if config.DebugMode {
		return &SkillTypeBase, nil
	}
	err := fmt.Errorf("invalid skill name: %s", name)
	return &SkillTypeBase, err
}

// GetSkillTypeFromEnum returns pointer to skill type with the requested name
// (but in enum.SkillType)
func GetSkillTypeFromEnum(name enum.SkillType) (*SkillType, error) {
	return GetSkillType(string(name))
}
