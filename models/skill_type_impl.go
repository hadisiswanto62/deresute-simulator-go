package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

// SkillTypeBase is a base skill (for unimplemented skills)
var SkillTypeBase = SkillType{
	Name: "BASE SKILL",
	IsActive: func(attr [6]enum.Attribute) bool {
		return false
	},
	ComboBonus: func(rarity enum.Rarity, currentHp float64) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity) int {
		return 0
	},
}

// SkillTypeScoreBonus = " Every k seconds: there is a l..m% chance
// that Perfect notes will receive a n% score bonus for o..p seconds."
var SkillTypeScoreBonus = SkillType{
	Name: "Score Bonus",
	IsActive: func(attr [6]enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp float64) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity) float64 {
		bonusMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.17,
			enum.RaritySR:  0.15,
			enum.RarityR:   0.1,
			enum.RarityN:   0.0,
		}
		return bonusMap[rarity]
	},
	TapHeal: func(rarity enum.Rarity) int {
		return 0
	},
}
