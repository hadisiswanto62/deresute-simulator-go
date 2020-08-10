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
	ComboBonus func(rarity enum.Rarity, currentHp float64) float64
	ScoreBonus func(rarity enum.Rarity) float64
	TapHeal    func(rarity enum.Rarity) int
}

// GetSkillType returns pointer to skill with the requested name
func GetSkillType(name string) (*SkillType, error) {
	switch name {
	case "Score Bonus", "Perfect Score Bonus":
		return &SkillTypeScoreBonus, nil
	}
	if config.DebugMode {
		return &SkillTypeBase, nil
	}
	err := fmt.Errorf("invalid skill name: %s", name)
	return &SkillTypeBase, err
}
