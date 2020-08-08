package parser

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/jsonmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

func MakeRarity(tmpRarity jsonmodels.TmpRarity) models.Rarity {
	var rarity enum.Rarity
	switch (tmpRarity.Rarity - 1) / 2 {
	case 0:
		rarity = enum.RarityN
	case 1:
		rarity = enum.RarityR
	case 2:
		rarity = enum.RaritySR
	case 3:
		rarity = enum.RaritySSR
	}
	isEvolved := tmpRarity.Rarity%2 == 0
	maxLevel := tmpRarity.BaseMaxLevel
	return models.Rarity{
		ID:        tmpRarity.Rarity,
		Rarity:    rarity,
		IsEvolved: isEvolved,
		MaxLevel:  maxLevel,
	}
}

func MakeIdol(tmpIdol jsonmodels.TmpIdol) models.Idol {
	return models.Idol(tmpIdol)
}
