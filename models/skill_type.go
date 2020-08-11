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
	ScoreBonus func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) float64
	TapHeal    func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int
}

// GetSkillType returns pointer to skill type with the requested name
func GetSkillType(name string) (*SkillType, error) {
	switch name {
	case string(enum.SkillTypeScoreBonus), "Perfect Score Bonus":
		return &SkillTypeScoreBonus, nil
	case string(enum.SkillTypeComboBonus):
		return &SkillTypeComboBonus, nil
	case string(enum.SkillTypeConcentration):
		return &SkillTypeConcentration, nil
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
