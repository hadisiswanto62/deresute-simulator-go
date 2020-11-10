package jsonmodels

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

func makeLeadSkill(tmp TmpLeadSkill) (*models.LeadSkill, error) {
	lskill, err := models.GetLeadSkill(tmp.Name)
	if err != nil {
		return lskill, fmt.Errorf("could not make lead skill: %v", err)
	}
	return lskill, nil
}

func makeSkill(tmp TmpSkill) (*models.Skill, error) {
	skillType, err := models.GetSkillType(tmp.SkillType, tmp.SkillTypeID)
	if err != nil {
		return nil, fmt.Errorf("could not make skill: %v", err)
	}
	activationCost := 0
	if skillType.Name == enum.SkillTypeOverload {
		activationCost = tmp.SkillTriggerValue
	}
	skill := models.Skill{
		ID:             tmp.ID,
		ProcChance:     tmp.ProcChance,
		EffectLength:   tmp.EffectLength,
		SkillType:      skillType,
		Timer:          tmp.Condition,
		ActivationCost: activationCost,
	}
	return &skill, nil
}

func makeRarity(tmp TmpRarity) *models.Rarity {
	return &models.Rarity{
		ID:        tmp.Rarity,
		Rarity:    tmp.ToEnum(),
		IsEvolved: tmp.IsEvolved(),
		MaxLevel:  tmp.BaseMaxLevel,
	}
}

func makeIdol(tmp TmpIdol) *models.Idol {
	return &models.Idol{
		ID:        tmp.ID,
		Name:      tmp.Name,
		Attribute: tmp.Attribute,
	}
}
