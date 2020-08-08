package jsonmodels

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

func makeLeadSkill(tmp TmpLeadSkill) (*models.LeadSkill, error) {
	lskill, err := models.GetLeadSkill(tmp.Name)
	if err != nil {
		return lskill, fmt.Errorf("could not make lead skill: %v", err)
	}
	return lskill, nil
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
