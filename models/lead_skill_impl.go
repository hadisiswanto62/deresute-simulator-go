package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

func attrStatUp(cardStat enum.Stat, cardAttr enum.Attribute, rarity enum.Rarity,
	requiredStat enum.Stat, requiredAttr enum.Attribute) float64 {
	bonus := 0.0
	if cardStat == requiredStat && cardAttr == requiredAttr {
		switch rarity {
		case enum.RaritySSR:
			bonus = 0.9
		case enum.RaritySR:
			bonus = 0.6
		case enum.RarityR:
			bonus = 0.3
		}
	}
	return bonus
}

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
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteMakeup = "Raises the Visual appeal of all Cute members by 90%/60%/30%."
var LeadSkillCuteMakeup = LeadSkill{
	Name: enum.LeadSkillCuteMakeup,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVisual, enum.AttrCute)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionVoice = "Raises the Visual appeal of all Cute members by 90%/60%/30%."
var LeadSkillPassionVoice = LeadSkill{
	Name: enum.LeadSkillPassionVoice,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVocal, enum.AttrPassion)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}
