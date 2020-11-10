package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func sampleCard() Card {
	idol := Idol{
		ID:        181,
		Name:      "Mifune Miyu",
		Attribute: enum.AttrCool,
	}
	rarity := Rarity{
		ID:        6,
		Rarity:    enum.RaritySR,
		IsEvolved: true,
		MaxLevel:  70,
	}
	skillType, _ := GetSkillType("Perfect Score Bonus", -1)
	skill := &Skill{
		ID:           200803,
		Timer:        13,
		ProcChance:   [2]int{4000, 6000},
		EffectLength: [2]int{600, 900},
		SkillType:    skillType,
	}
	leadSkill, _ := GetLeadSkill("クールボイス")
	return Card{
		ID:        200804,
		SeriesID:  200803,
		Idol:      &idol,
		Rarity:    &rarity,
		LeadSkill: leadSkill,
		Skill:     skill,

		BonusDance:  143,
		BonusHp:     2,
		BonusVisual: 119,
		BonusVocal:  224,
		DanceMax:    3517,
		DanceMin:    1913,
		HpMax:       37,
		HpMin:       37,
		VisualMax:   2906,
		VisualMin:   1583,
		VocalMax:    5501,
		VocalMin:    2993,
	}
}
