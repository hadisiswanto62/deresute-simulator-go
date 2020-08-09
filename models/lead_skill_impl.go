package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

// LeadSkillBase is a base lead skill (for unimplemented lead skills)
var LeadSkillBase = LeadSkill{
	Name: enum.LeadSkillBase,
	IsActive: func([6]enum.Attribute) bool {
		return false
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) int {
		return 0
	},
}

// LeadSkillCuteMakeup = "Raises the Visual appeal of all Cute members by 90%/60%/30%."
var LeadSkillCuteMakeup = LeadSkill{
	Name: enum.LeadSkillCuteMakeup,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		requiredStat := enum.StatVisual
		requiredAttr := enum.AttrCute
		bonus := 0.0
		switch rarity {
		case enum.RaritySSR:
			bonus = 0.9
		case enum.RaritySR:
			bonus = 0.6
		case enum.RarityR:
			bonus = 0.3
		}
		if cardAttr == requiredAttr && stat == requiredStat {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) int {
		return 0
	},
}
