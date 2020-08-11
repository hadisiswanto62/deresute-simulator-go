package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

// SkillTypeBase is a base skill (for unimplemented skills)
var SkillTypeBase = SkillType{
	Name: enum.SkillTypeBase,
	IsActive: func(attr [6]enum.Attribute) bool {
		return false
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int {
		return 0
	},
}

// SkillTypeScoreBonus = " Every k seconds: there is a l..m% chance
// that Great/Perfect notes will receive a n% score bonus for o..p seconds."
// Also covers PerfectScoreBonus
var SkillTypeScoreBonus = SkillType{
	Name: enum.SkillTypeScoreBonus,
	IsActive: func(attr [6]enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		if rarity == enum.RaritySSR {
			if judgement == enum.TapJudgementPerfect || judgement == enum.TapJudgementGreat {
				return 0.17
			}
		} else {
			if judgement == enum.TapJudgementPerfect {
				switch rarity {
				case enum.RaritySR:
					return 0.15
				case enum.RarityR:
					return 0.1
				}
			}
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int {
		return 0
	},
}

// SkillTypeComboBonus = "Every k seconds: there is a l..m% chance
// that you will gain an extra n% combo bonus for o..p seconds. "
var SkillTypeComboBonus = SkillType{
	Name: enum.SkillTypeComboBonus,
	IsActive: func(attr [6]enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		switch rarity {
		case enum.RaritySSR:
			return 0.18
		case enum.RaritySR:
			return 0.14
		case enum.RarityR:
			return 0.08
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int {
		return 0
	},
}

// SkillTypeConcentration = "Every k seconds: there is a l..m% chance
// that Perfect notes will receive a n% score bonus, but become harder to hit for o..p seconds."
var SkillTypeConcentration = SkillType{
	Name: enum.SkillTypeConcentration,
	IsActive: func(attr [6]enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
		if judgement == enum.TapJudgementPerfect {
			switch rarity {
			case enum.RaritySSR:
				return 0.22
			case enum.RaritySR:
				return 0.19
			}
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteType enum.NoteType) int {
		return 0
	},
}
