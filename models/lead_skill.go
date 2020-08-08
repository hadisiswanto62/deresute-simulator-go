package models

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// LeadSkill represents a leader skill. (Do not create new instance. Use GetLeadSkill(name) instead)
type LeadSkill struct {
	Name           string
	IsActive       func(attr [6]enum.Attribute) bool
	StatBonus      func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64
	SkillProbBonus func(rarity enum.Rarity, cardAttr enum.Attribute) float64
	HpBonus        func(rarity enum.Rarity, cardAttr enum.Attribute) uint8
}

// GetLeadSkill returns pointer to lead skill with the requested name
func GetLeadSkill(name string) (*LeadSkill, error) {
	switch name {
	case "キュートメイク":
		return &LeadSkillCuteMakeup, nil
	}

	returnError := false
	if returnError {
		err := fmt.Errorf("invalid skill name: %s", name)
		return &LeadSkillBase, err
	} else {
		return &LeadSkillBase, nil
	}
}
