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

// LeadSkillIrrelevant is for irrelevant lead skills
var LeadSkillIrrelevant = LeadSkill{
	Name: enum.LeadSkillIrrelevant,
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

// LeadSkillCuteMakeup = "Raises the Visual appeal of all Cute members by x%"
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

// LeadSkillCuteStep = "Raises the Dance appeal of all Cute members by x%"
var LeadSkillCuteStep = LeadSkill{
	Name: enum.LeadSkillCuteStep,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatDance, enum.AttrCute)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteVoice = "Raises the Vocal appeal of all Cute members by x%"
var LeadSkillCuteVoice = LeadSkill{
	Name: enum.LeadSkillCuteVoice,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVocal, enum.AttrCute)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionMakeup = "Raises the Visual appeal of all Passion members by x%"
var LeadSkillPassionMakeup = LeadSkill{
	Name: enum.LeadSkillPassionMakeup,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVisual, enum.AttrPassion)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionStep = "Raises the Dance appeal of all Passion members by x%"
var LeadSkillPassionStep = LeadSkill{
	Name: enum.LeadSkillPassionStep,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatDance, enum.AttrPassion)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionVoice = "Raises the Vocal appeal of all Passion members by x%"
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

// LeadSkillCoolMakeup = "Raises the Visual appeal of all Cool members by x%"
var LeadSkillCoolMakeup = LeadSkill{
	Name: enum.LeadSkillCoolMakeup,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVisual, enum.AttrCool)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolStep = "Raises the Dance appeal of all Cool members by x%"
var LeadSkillCoolStep = LeadSkill{
	Name: enum.LeadSkillCoolStep,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatDance, enum.AttrCool)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolVoice = "Raises the Vocal appeal of all Cool members by x%"
var LeadSkillCoolVoice = LeadSkill{
	Name: enum.LeadSkillCoolVoice,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return attrStatUp(stat, cardAttr, rarity, enum.StatVocal, enum.AttrCool)
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolAbility = "Raises the skill probability of all <Attr> members by x%"
var LeadSkillCoolAbility = LeadSkill{
	Name: enum.LeadSkillCoolAbility,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrCool {
			rarityMap := map[enum.Rarity]float64{
				enum.RaritySSR: 0.4,
				enum.RaritySR:  0.3,
				enum.RarityR:   0.15,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteAbility = "Raises the skill probability of all <Attr> members by x%"
var LeadSkillCuteAbility = LeadSkill{
	Name: enum.LeadSkillCuteAbility,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrCute {
			rarityMap := map[enum.Rarity]float64{
				enum.RaritySSR: 0.4,
				enum.RaritySR:  0.3,
				enum.RarityR:   0.15,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionAbility = "Raises the skill probability of all <Attr> members by x%"
var LeadSkillPassionAbility = LeadSkill{
	Name: enum.LeadSkillPassionAbility,
	IsActive: func([6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrPassion {
			rarityMap := map[enum.Rarity]float64{
				enum.RaritySSR: 0.4,
				enum.RaritySR:  0.3,
				enum.RarityR:   0.15,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteCheer = " If there are only Cute idols on the team:
// Raises the life of all Cute members by 40%."
var LeadSkillCuteCheer = LeadSkill{
	Name: enum.LeadSkillCuteCheer,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrCute {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.4
	},
}

// LeadSkillCoolCheer = " If there are only Cool idols on the team:
// Raises the life of all Cool members by 40%."
var LeadSkillCoolCheer = LeadSkill{
	Name: enum.LeadSkillCoolCheer,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrCool {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.4
	},
}

// LeadSkillPassionCheer = " If there are only Passion idols on the team:
// Raises the life of all Passion members by 40%."
var LeadSkillPassionCheer = LeadSkill{
	Name: enum.LeadSkillPassionCheer,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrPassion {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.4
	},
}

// LeadSkillCutePrincess =  "If there are only Cute idols on the team:
// Raises all appeals of all Cute members by 50%.
var LeadSkillCutePrincess = LeadSkill{
	Name: enum.LeadSkillCutePrincess,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrCute {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.5,
			enum.RaritySR:  0.35,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolPrincess =  "If there are only Cool idols on the team:
// Raises all appeals of all Cool members by 50%.
var LeadSkillCoolPrincess = LeadSkill{
	Name: enum.LeadSkillCoolPrincess,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrCool {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.5,
			enum.RaritySR:  0.35,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionPrincess =  "If there are only Passion idols on the team:
// Raises all appeals of all Passion members by 50%.
var LeadSkillPassionPrincess = LeadSkill{
	Name: enum.LeadSkillPassionPrincess,
	IsActive: func(attrs [6]enum.Attribute) bool {
		for _, attr := range attrs {
			if attr != enum.AttrPassion {
				return false
			}
		}
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.5,
			enum.RaritySR:  0.35,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteUnison = "Raises all appeals of all Cute members by 30% (55% when playing a Cute-type song)."
var LeadSkillCuteUnison = LeadSkill{
	Name: enum.LeadSkillCuteUnison,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrCute {
			if songAttr == enum.AttrCute {
				return 0.55
			}
			return 0.3
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolUnison = "Raises all appeals of all Cool members by 30% (55% when playing a Cool-type song)."
var LeadSkillCoolUnison = LeadSkill{
	Name: enum.LeadSkillCoolUnison,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrCool {
			if songAttr == enum.AttrCool {
				return 0.55
			}
			return 0.3
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionUnison = "Raises all appeals of all Passion members by 30% (55% when playing a Passion-type song)."
var LeadSkillPassionUnison = LeadSkill{
	Name: enum.LeadSkillPassionUnison,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if cardAttr == enum.AttrPassion {
			if songAttr == enum.AttrPassion {
				return 0.55
			}
			return 0.3
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteBrilliance = " Raises all appeals of all Cute members by 40%."
var LeadSkillCuteBrilliance = LeadSkill{
	Name: enum.LeadSkillCuteBrilliance,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.4,
			enum.RaritySR:  0.3,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrCute {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolBrilliance = " Raises all appeals of all Cool members by 40%."
var LeadSkillCoolBrilliance = LeadSkill{
	Name: enum.LeadSkillCoolBrilliance,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.4,
			enum.RaritySR:  0.3,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrCool {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionBrilliance = " Raises all appeals of all Passion members by 40%."
var LeadSkillPassionBrilliance = LeadSkill{
	Name: enum.LeadSkillPassionBrilliance,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.4,
			enum.RaritySR:  0.3,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrPassion {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteEnergy = "  Raises the life of all Cute members by 30%. "
var LeadSkillCuteEnergy = LeadSkill{
	Name: enum.LeadSkillCuteEnergy,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrCute {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
}

// LeadSkillCoolEnergy = "  Raises the life of all Cool members by 30%. "
var LeadSkillCoolEnergy = LeadSkill{
	Name: enum.LeadSkillCoolEnergy,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrCool {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
}

// LeadSkillPassionEnergy = "  Raises the life of all Passion members by 30%. "
var LeadSkillPassionEnergy = LeadSkill{
	Name: enum.LeadSkillPassionEnergy,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
			enum.RarityR:   0.1,
		}
		if cardAttr == enum.AttrPassion {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
}

// LeadSkillTricolorVoice = " If there are Cute, Cool, and Passion idols on the team:
// Raises the Vocal appeal of all members by 100%. "
var LeadSkillTricolorVoice = LeadSkill{
	Name: enum.LeadSkillTricolorVoice,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatVocal {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 1.0,
				enum.RaritySR:  0.8,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillTricolorStep = " If there are Cute, Cool, and Passion idols on the team:
// Raises the Dance appeal of all members by 100%. "
var LeadSkillTricolorStep = LeadSkill{
	Name: enum.LeadSkillTricolorStep,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatDance {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 1.0,
				enum.RaritySR:  0.8,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillTricolorMakeup = " If there are Cute, Cool, and Passion idols on the team:
// Raises the Visual appeal of all members by 100%. "
var LeadSkillTricolorMakeup = LeadSkill{
	Name: enum.LeadSkillTricolorMakeup,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatVisual {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 1.0,
				enum.RaritySR:  0.8,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillTricolorAbility =  If there are Cute, Cool, and Passion idols on the team:
// Raises the skill probability of all members by 50%.
var LeadSkillTricolorAbility = LeadSkill{
	Name: enum.LeadSkillTricolorAbility,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.5,
			enum.RaritySR:  0.4,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillShinyVoice : " Raises the Vocal appeal of all members by 80%. "
var LeadSkillShinyVoice = LeadSkill{
	Name: enum.LeadSkillShinyVoice,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatVocal {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 0.8,
				enum.RaritySR:  0.48,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillShinyMakeup : " Raises the Visual appeal of all members by 80%. "
var LeadSkillShinyMakeup = LeadSkill{
	Name: enum.LeadSkillShinyMakeup,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatVisual {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 0.8,
				enum.RaritySR:  0.48,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillShinyStep : " Raises the Dance appeal of all members by 80%. "
var LeadSkillShinyStep = LeadSkill{
	Name: enum.LeadSkillShinyStep,
	IsActive: func(attrs [6]enum.Attribute) bool {
		return true
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		if stat == enum.StatDance {
			var rarityMap = map[enum.Rarity]float64{
				enum.RaritySSR: 0.8,
				enum.RaritySR:  0.48,
			}
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteCrossCool = " If there are Cute and Cool idols on the team:
// Raises all appeals of all members by 30%, and of all members by 10%. "
var LeadSkillCuteCrossCool = LeadSkill{
	Name: enum.LeadSkillCuteCrossCool,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		CuteExist, _ := attrMap[enum.AttrCute]
		CoolExist, _ := attrMap[enum.AttrCool]
		return CuteExist && CoolExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCuteCrossPassion = " If there are Cute and Passion idols on the team:
// Raises all appeals of all members by 30%, and of all members by 10%. "
var LeadSkillCuteCrossPassion = LeadSkill{
	Name: enum.LeadSkillCuteCrossPassion,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		CuteExist, _ := attrMap[enum.AttrCute]
		PassionExist, _ := attrMap[enum.AttrPassion]
		return CuteExist && PassionExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolCrossCute =  If there are Cute and Cool idols on the team:
// Raises all appeals of all members by 20%, and the skill probability of all members by 25%.
var LeadSkillCoolCrossCute = LeadSkill{
	Name: enum.LeadSkillCoolCrossCute,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		CuteExist, _ := attrMap[enum.AttrCute]
		CoolExist, _ := attrMap[enum.AttrCool]
		return CuteExist && CoolExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.35,
			enum.RaritySR:  0.25,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillCoolCrossPassion =  If there are Cool and Passion idols on the team:
// Raises all appeals of all members by 20%, and the skill probability of all members by 25%.
var LeadSkillCoolCrossPassion = LeadSkill{
	Name: enum.LeadSkillCoolCrossPassion,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		PassionExist, _ := attrMap[enum.AttrPassion]
		CoolExist, _ := attrMap[enum.AttrCool]
		return PassionExist && CoolExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.35,
			enum.RaritySR:  0.25,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
}

// LeadSkillPassionCrossCool =  " If there are Cool and Passion idols on the team:
// Raises all appeals of all members by 30%, and the life of all members by 25%. "
var LeadSkillPassionCrossCool = LeadSkill{
	Name: enum.LeadSkillPassionCrossCool,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		PassionExist, _ := attrMap[enum.AttrPassion]
		CoolExist, _ := attrMap[enum.AttrCool]
		return PassionExist && CoolExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.25,
			enum.RaritySR:  0.15,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
}

// LeadSkillPassionCrossCute =  " If there are Cute and Passion idols on the team:
// Raises all appeals of all members by 30%, and the life of all members by 25%. "
var LeadSkillPassionCrossCute = LeadSkill{
	Name: enum.LeadSkillPassionCrossCute,
	IsActive: func(attrs [6]enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attrs {
			attrMap[attribute] = true
		}
		PassionExist, _ := attrMap[enum.AttrPassion]
		CuteExist, _ := attrMap[enum.AttrCute]
		return PassionExist && CuteExist
	},
	StatBonus: func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.3,
			enum.RaritySR:  0.2,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	SkillProbBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		return 0.0
	},
	HpBonus: func(rarity enum.Rarity, cardAttr enum.Attribute) float64 {
		var rarityMap = map[enum.Rarity]float64{
			enum.RaritySSR: 0.25,
			enum.RaritySR:  0.15,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
}

// LeadSkillResonantMakeup : " If the team has at least 5 different skill types:
// Allows active skill effects to stack, but only the Visual appeal of the team applies during the live. "
var LeadSkillResonantMakeup = LeadSkill{
	Name: enum.LeadSkillResonantMakeup,
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

// LeadSkillResonantVoice : " If the team has at least 5 different skill types:
// Allows active skill effects to stack, but only the Vocal appeal of the team applies during the live. "
var LeadSkillResonantVoice = LeadSkill{
	Name: enum.LeadSkillResonantVoice,
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

// LeadSkillResonantStep : " If the team has at least 5 different skill types:
// Allows active skill effects to stack, but only the Dance appeal of the team applies during the live. "
var LeadSkillResonantStep = LeadSkill{
	Name: enum.LeadSkillResonantStep,
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
