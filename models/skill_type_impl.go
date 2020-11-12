package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// SkillTypeBase is a base skill (for unimplemented skills)
var SkillTypeBase = SkillType{
	Name: enum.SkillTypeBase,
	IsActive: func(attr []enum.Attribute) bool {
		return false
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeScoreBonus = " Every k seconds: there is a l..m% chance
// that Great/Perfect notes will receive a n% score bonus for o..p seconds."
// Also covers PerfectScoreBonus
var SkillTypeScoreBonus = SkillType{
	Name: enum.SkillTypeScoreBonus,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
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
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeComboBonus = "Every k seconds: there is a l..m% chance
// that you will gain an extra n% combo bonus for o..p seconds. "
var SkillTypeComboBonus = SkillType{
	Name: enum.SkillTypeComboBonus,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
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
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeConcentration = "Every k seconds: there is a l..m% chance
// that Perfect notes will receive a n% score bonus, but become harder to hit for o..p seconds."
var SkillTypeConcentration = SkillType{
	Name: enum.SkillTypeConcentration,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
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
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeHealer = "Every k seconds: there is a l..m% chance
// that Perfect notes will restore n health for o..p seconds."
var SkillTypeHealer = SkillType{
	Name: enum.SkillTypeHealer,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		if judgement != enum.TapJudgementPerfect {
			return 0
		}
		switch rarity {
		case enum.RaritySSR:
			return 3
		case enum.RaritySR:
			return 3
		case enum.RarityN:
			return 2
		}
		return 0
	},
}

// SkillTypeAllRound = Every 9 seconds: there is a 35..52.5% chance
// that you will gain an extra 13% combo bonus, and Perfect notes
// will restore 1 health for 4..6 seconds.
var SkillTypeAllRound = SkillType{
	Name: enum.SkillTypeAllRound,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		switch rarity {
		case enum.RaritySSR:
			return 0.13
		case enum.RaritySR:
			return 0.10
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		if judgement == enum.TapJudgementPerfect {
			return 1
		}
		return 0
	},
}

// SkillTypeCoordinate = " Every 9 seconds: there is a 40..60% chance
// that Perfect notes will receive a 10% score bonus, and you will
// gain an extra 15% combo bonus for 4..6 seconds. "
var SkillTypeCoordinate = SkillType{
	Name: enum.SkillTypeCoordinate,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.15,
			enum.RaritySR:  0.12,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.1,
			enum.RaritySR:  0.08,
		}
		if judgement == enum.TapJudgementPerfect {
			if bonus, ok := rarityMap[rarity]; ok {
				return bonus
			}
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeOverload = " Every 9 seconds: there is a 35..52.5% chance
// that 15 life will be consumed, then: Perfect/Great notes receive a
// 18% score bonus, and Nice/Bad notes will not break combo for 5..7.5 seconds. "
var SkillTypeOverload = SkillType{
	Name: enum.SkillTypeOverload,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.18,
			enum.RaritySR:  0.16,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeTricolorSynergy = "Every 11 seconds: there is a 40..60% chance
// that with all three types of idols on the team, you will gain an extra 15%
// combo bonus, and Perfect notes will receive a 16% score bonus plus restore 1 HP, for 5..7.5 seconds.
var SkillTypeTricolorSynergy = SkillType{
	Name: enum.SkillTypeTricolorSynergy,
	IsActive: func(attr []enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attr {
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.15
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		if judgement == enum.TapJudgementPerfect {
			return 0.16
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		if judgement == enum.TapJudgementPerfect {
			return 1
		}
		return 0
	},
}

// SkillTypeTuning = " Every 11 seconds: there is a 35..52.5% chance
// that you will gain an extra 12% combo bonus, and Nice/Great notes
// will become Perfect notes for 6..9 seconds."
var SkillTypeTuning = SkillType{
	Name: enum.SkillTypeTuning,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.12
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypePerfectLock = "Every 9 seconds: there is a 40..60% chance
// that Bad/Nice/Great notes will become Perfect notes for 3..4.5 seconds. "
var SkillTypePerfectLock = SkillType{
	Name: enum.SkillTypePerfectLock,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeComboGuard = " Every 12 seconds: there is a 35..52.5% chance
// that Nice notes will not break combo for 5..7.5 seconds."
var SkillTypeComboGuard = SkillType{
	Name: enum.SkillTypeComboGuard,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeLifeSparkle = " Every 9 seconds: there is a 40..60% chance
// that you will gain an extra combo bonus based on your current health for 4..6 seconds. "
var SkillTypeLifeSparkle = SkillType{
	Name: enum.SkillTypeLifeSparkle,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		ssrTable := map[int]int{
			0: 9, 50: 10, 110: 11, 160: 12,
			200: 13, 250: 14, 300: 15, 330: 16, 380: 17,
			410: 18, 440: 19, 480: 20, 500: 21, 540: 22,
			580: 23, 610: 24, 810: 25, 860: 26, 910: 27,
			950: 28, 1000: 29, 1050: 30, 1100: 31, 1140: 32,
			1190: 33, 1240: 34, 1290: 35, 1330: 36, 1380: 37,
			1430: 38, 1480: 39, 1520: 40, 1570: 41,
		}
		srTable := map[int]int{
			0: 7, 260: 12, 320: 13, 330: 14, 380: 15,
			420: 16, 470: 17, 480: 18, 520: 19, 570: 20,
			610: 21, 830: 22, 890: 23, 940: 24, 990: 25,
			1040: 26, 1090: 27, 1140: 28, 1200: 29, 1250: 30,
			1300: 31, 1350: 32, 1400: 33, 1460: 34, 1510: 35, 1560: 36,
		}
		rarityMap := map[enum.Rarity]map[int]int{
			enum.RaritySSR: ssrTable,
			enum.RaritySR:  srTable,
		}
		usedTable, ok := rarityMap[rarity]
		if !ok {
			return 0.0
		}
		nearest, chosenBonus := 9999, 0
		// algorithm:
		// iterate over every key, if key is lower than currentHP, check if it is the closest
		// then return the value of closest key
		for life, bonus := range usedTable {
			lifeDiff := currentHp - life
			if lifeDiff < 0 {
				continue
			}
			if lifeDiff < nearest {
				nearest = lifeDiff
				chosenBonus = bonus
			}
		}
		return float64(chosenBonus) / 100.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeLifeGuard = "Every 11 seconds: there is a 35..52.5% chance
// that you will not lose health for 6..9 seconds."
var SkillTypeLifeGuard = SkillType{
	Name: enum.SkillTypeLifeGuard,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeSkillBoost = " Every 11 seconds: there is a 40..60% chance
// to boost the effects of currently active skills for 5..7.5 seconds. "
var SkillTypeSkillBoost = SkillType{
	Name: enum.SkillTypeSkillBoost,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
	ScoreComboBonusBonus: func(attr enum.Attribute) float64 {
		return 0.2
	},
	TapHealBonus: func() float64 {
		return 0.2
	},
}

// SkillTypeCuteFocus = " Every 11 seconds: there is a 40..60% chance
// that with only Cute idols on the team, Perfect notes will receive a
// 14% score bonus, and you will gain an extra 11% combo bonus for 5..7.5 seconds. "
var SkillTypeCuteFocus = SkillType{
	Name: enum.SkillTypeCuteFocus,
	IsActive: func(attr []enum.Attribute) bool {
		for _, attribute := range attr {
			if attribute == "" {
				continue
			}
			if attribute != enum.AttrCute {
				return false
			}
		}
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.14,
			enum.RaritySR:  0.11,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.16,
			enum.RaritySR:  0.14,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeCoolFocus = " Every 11 seconds: there is a 40..60% chance
// that with only Cool idols on the team, Perfect notes will receive a
// 14% score bonus, and you will gain an extra 11% combo bonus for 5..7.5 seconds. "
var SkillTypeCoolFocus = SkillType{
	Name: enum.SkillTypeCoolFocus,
	IsActive: func(attr []enum.Attribute) bool {
		for _, attribute := range attr {
			if attribute == "" {
				continue
			}
			if attribute != enum.AttrCool {
				return false
			}
		}
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.14,
			enum.RaritySR:  0.11,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.16,
			enum.RaritySR:  0.14,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypePassionFocus = " Every 11 seconds: there is a 40..60% chance
// that with only Passion idols on the team, Perfect notes will receive a
// 14% score bonus, and you will gain an extra 11% combo bonus for 5..7.5 seconds. "
var SkillTypePassionFocus = SkillType{
	Name: enum.SkillTypePassionFocus,
	IsActive: func(attr []enum.Attribute) bool {
		for _, attribute := range attr {
			if attribute == "" {
				continue
			}
			if attribute != enum.AttrPassion {
				return false
			}
		}
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.14,
			enum.RaritySR:  0.11,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		rarityMap := map[enum.Rarity]float64{
			enum.RaritySSR: 0.16,
			enum.RaritySR:  0.14,
		}
		if bonus, ok := rarityMap[rarity]; ok {
			return bonus
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeEncore = " Every 9 seconds: there is a 35..52.5% chance
// to activate the previous skill again for 3..4.5 seconds."
var SkillTypeEncore = SkillType{
	Name: enum.SkillTypeEncore,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeTricolorSymphony = " Every 9 seconds: there is a 40..60% chance
// that with all three types of idols on the team, to boost the score/combo
// bonus/health recovery of currently active skills for 4..6 seconds.
var SkillTypeTricolorSymphony = SkillType{
	Name: enum.SkillTypeTricolorSymphony,
	IsActive: func(attr []enum.Attribute) bool {
		attrMap := make(map[enum.Attribute]bool)
		for _, attribute := range attr {
			if attribute == "" {
				continue
			}
			attrMap[attribute] = true
		}
		return len(attrMap) == 3
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
	ScoreComboBonusBonus: func(attr enum.Attribute) float64 {
		return 0.5
	},
	TapHealBonus: func() float64 {
		return 0.2
	},
}

// SkillTypeAlternate = " Every 6 seconds: there is a 35..52.5% chance
// to reduce combo bonus by 20%, but also apply the highest score bonus
// gained so far with a boost of 50% for 3..4.5 seconds.
var SkillTypeAlternate = SkillType{
	Name: enum.SkillTypeAlternate,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return -0.2
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		// not handled here
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeVisualMotif = " Every 7 seconds: there is a 40..60% chance that
// Perfect notes will receive a score bonus determined by the team's Visual appeal for 3..4.5 seconds."
// Note: use BASE appeal of 5 members
var SkillTypeVisualMotif = SkillType{
	Name: enum.SkillTypeVisualMotif,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		bonusTable := map[int]int{
			10: 1, 3000: 2, 6000: 3, 9000: 4, 12000: 5,
			15000: 6, 17000: 7, 19000: 8, 21000: 9, 22000: 10,
			23000: 11, 24000: 12, 26000: 13, 27000: 15, 29000: 16,
			32000: 17, 36000: 18, 40000: 19, 42000: 20, 43000: 21,
			45000: 23, 47000: 25, 49000: 26, 52000: 27,
		}
		keys := []int{
			10, 3000, 6000, 9000, 12000, 15000, 17000, 19000, 21000, 22000,
			23000, 24000, 26000, 27000, 29000, 32000, 36000, 40000, 42000, 43000,
			45000, 47000, 49000, 52000,
		}
		if judgement == enum.TapJudgementPerfect {
			nearest, chosenBonus := 99999999, 0
			// algorithm:
			// iterate over every key, if key is lower than the team's appeal, check if it is the closest
			// then return the value of closest key
			for _, appeal := range keys {
				bonus, ok := bonusTable[appeal]
				if !ok {
					return -100.0
				}
				appealDiff := baseVisual - appeal
				if appealDiff < 0 {
					continue
				}
				if appealDiff < nearest {
					nearest = appealDiff
					chosenBonus = bonus
				} else {
					break
				}
			}
			return float64(chosenBonus) / 100.0
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeVocalMotif = " Every 7 seconds: there is a 40..60% chance that
// Perfect notes will receive a score bonus determined by the team's Vocal appeal for 3..4.5 seconds."
// Note: use BASE appeal of 5 members
var SkillTypeVocalMotif = SkillType{
	Name: enum.SkillTypeVocalMotif,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		bonusTable := map[int]int{
			10: 1, 3000: 2, 6000: 3, 9000: 4, 12000: 5,
			15000: 6, 17000: 7, 19000: 8, 21000: 9, 22000: 10,
			23000: 11, 24000: 12, 26000: 13, 27000: 15, 29000: 16,
			32000: 17, 36000: 18, 40000: 19, 42000: 20, 43000: 21,
			45000: 23, 47000: 25, 49000: 26, 52000: 27,
		}
		keys := []int{
			10, 3000, 6000, 9000, 12000, 15000, 17000, 19000, 21000, 22000,
			23000, 24000, 26000, 27000, 29000, 32000, 36000, 40000, 42000, 43000,
			45000, 47000, 49000, 52000,
		}
		if judgement == enum.TapJudgementPerfect {
			nearest, chosenBonus := 99999999, 0
			// algorithm:
			// iterate over every key, if key is lower than the team's appeal, check if it is the closest
			// then return the value of closest key
			for _, appeal := range keys {
				bonus, ok := bonusTable[appeal]
				if !ok {
					return -100.0
				}
				appealDiff := baseVocal - appeal
				if appealDiff < 0 {
					continue
				}
				if appealDiff < nearest {
					nearest = appealDiff
					chosenBonus = bonus
				} else {
					break
				}
			}
			return float64(chosenBonus) / 100.0
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeDanceMotif = " Every 7 seconds: there is a 40..60% chance that
// Perfect notes will receive a score bonus determined by the team's Vocal appeal for 3..4.5 seconds."
// Note: use BASE appeal of 5 members
var SkillTypeDanceMotif = SkillType{
	Name: enum.SkillTypeDanceMotif,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		bonusTable := map[int]int{
			10: 1, 3000: 2, 6000: 3, 9000: 4, 12000: 5,
			15000: 6, 17000: 7, 19000: 8, 21000: 9, 22000: 10,
			23000: 11, 24000: 12, 26000: 13, 27000: 15, 29000: 16,
			32000: 17, 36000: 18, 40000: 19, 42000: 20, 43000: 21,
			45000: 23, 47000: 25, 49000: 26, 52000: 27,
		}
		keys := []int{
			10, 3000, 6000, 9000, 12000, 15000, 17000, 19000, 21000, 22000,
			23000, 24000, 26000, 27000, 29000, 32000, 36000, 40000, 42000, 43000,
			45000, 47000, 49000, 52000,
		}
		if judgement == enum.TapJudgementPerfect {
			nearest, chosenBonus := 99999999, 0
			// algorithm:
			// iterate over every key, if key is lower than the team's appeal, check if it is the closest
			// then return the value of closest key
			for _, appeal := range keys {
				bonus, ok := bonusTable[appeal]
				if !ok {
					return -100.0
				}
				appealDiff := baseDance - appeal
				if appealDiff < 0 {
					continue
				}
				if appealDiff < nearest {
					nearest = appealDiff
					chosenBonus = bonus
				} else {
					break
				}
			}
			return float64(chosenBonus) / 100.0
		}
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeCuteEnsemble = "  Every 7 seconds: there is a 40..60% chance to
// boost the score/combo bonus of Cute idols' active skills for 3..4.5 seconds.""
var SkillTypeCuteEnsemble = SkillType{
	Name: enum.SkillTypeCuteEnsemble,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
	ScoreComboBonusBonus: func(attr enum.Attribute) float64 {
		required := enum.AttrCute
		if attr == required {
			return 0.5
		}
		return 0.0
	},
}

// SkillTypeCoolEnsemble = "  Every 7 seconds: there is a 40..60% chance to
// boost the score/combo bonus of Cool idols' active skills for 3..4.5 seconds.""
var SkillTypeCoolEnsemble = SkillType{
	Name: enum.SkillTypeCoolEnsemble,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
	ScoreComboBonusBonus: func(attr enum.Attribute) float64 {
		required := enum.AttrCool
		if attr == required {
			return 0.5
		}
		return 0.0
	},
}

// SkillTypePassionEnsemble = "  Every 7 seconds: there is a 40..60% chance to
// boost the score/combo bonus of Cute idols' active skills for 3..4.5 seconds.""
var SkillTypePassionEnsemble = SkillType{
	Name: enum.SkillTypePassionEnsemble,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
	ScoreComboBonusBonus: func(attr enum.Attribute) float64 {
		required := enum.AttrPassion
		if attr == required {
			return 0.5
		}
		return 0.0
	},
}

// SkillTypeFlickAct = "Every 11 seconds: there is a 40..60% chance that
// Perfect notes will receive a 8/10% score bonus, and flick notes a 24/30% score bonus for 6..9 seconds."
var SkillTypeFlickAct = SkillType{
	Name: enum.SkillTypeFlickAct,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		required := enum.NoteTypeFlick
		rarityMap := map[enum.Rarity][2]float64{
			enum.RaritySR:  [2]float64{0.08, 0.24},
			enum.RaritySSR: [2]float64{0.1, 0.3},
		}
		value, ok := rarityMap[rarity]
		if !ok {
			return -100
		}
		for _, noteType := range noteTypes {
			if noteType == required {
				return value[1]
			}
		}
		return value[0]
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeSlideAct = "Every 11 seconds: there is a 40..60% chance that
// Perfect notes will receive a 8/10% score bonus, and slide notes a 24/30% score bonus for 6..9 seconds."
var SkillTypeSlideAct = SkillType{
	Name: enum.SkillTypeSlideAct,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		required := enum.NoteTypeSlide
		rarityMap := map[enum.Rarity][2]float64{
			enum.RaritySR:  [2]float64{0.08, 0.24},
			enum.RaritySSR: [2]float64{0.1, 0.3},
		}
		value, ok := rarityMap[rarity]
		if !ok {
			return -100
		}
		for _, noteType := range noteTypes {
			if noteType == required {
				return value[1]
			}
		}
		return value[0]
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}

// SkillTypeHoldAct = "Every 11 seconds: there is a 40..60% chance that
// Perfect notes will receive a 8/10% score bonus, and long notes a 24/30% score bonus for 6..9 seconds."
var SkillTypeHoldAct = SkillType{
	Name: enum.SkillTypeHoldAct,
	IsActive: func(attr []enum.Attribute) bool {
		return true
	},
	ComboBonus: func(rarity enum.Rarity, currentHp int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		return 0.0
	},
	ScoreBonus: func(rarity enum.Rarity, baseVisual int, baseDance int, baseVocal int, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
		required := enum.NoteTypeHold
		rarityMap := map[enum.Rarity][2]float64{
			enum.RaritySR:  [2]float64{0.08, 0.24},
			enum.RaritySSR: [2]float64{0.1, 0.3},
		}
		value, ok := rarityMap[rarity]
		if !ok {
			return -100
		}
		for _, noteType := range noteTypes {
			if noteType == required {
				return value[1]
			}
		}
		return value[0]
	},
	TapHeal: func(rarity enum.Rarity, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
		return 0
	},
}
