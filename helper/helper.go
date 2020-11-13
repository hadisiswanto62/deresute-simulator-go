package helper

import (
	"math"
	"math/rand"

	"github.com/hadisiswanto62/deresute-simulator-go/config"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Scale scales from min to max (max-min is divided into `steps`-1 parts evenly)
func Scale(min, max int, maxLevel, currentLevel int) int {
	step := float64(max-min) / (float64(maxLevel) - 1)
	add := step * float64(currentLevel-1)
	x := math.Floor(float64(min) + add)
	return int(x)
}

// StatPotentialBonusLookupFor returns lookup table of stat bonus from potential for the specified rarity
func StatPotentialBonusLookupFor(rarity enum.Rarity) [11]int {
	switch rarity {
	case enum.RaritySSR:
		return ssrStatPotentialBonusLookup
	case enum.RaritySR:
		return srStatPotentialBonusLookup
	case enum.RarityR:
		return rStatPotentialBonusLookup
	case enum.RarityN:
		return nStatPotentialBonusLookup
	}
	return [11]int{}
}

// LifePotentialBonusLookupFor returns lookup table of life bonus from potential for the specified rarity
func LifePotentialBonusLookupFor(rarity enum.Rarity) [11]int {
	switch rarity {
	case enum.RaritySSR:
		return ssrLifePotentialBonusLookup
	case enum.RaritySR:
		return srLifePotentialBonusLookup
	case enum.RarityR:
		return rLifePotentialBonusLookup
	case enum.RarityN:
		return nLifePotentialBonusLookup
	}
	return [11]int{}
}

// GetRoomItemBonus gets room item bonus for specified attributes from config
func GetRoomItemBonus(attr enum.Attribute) float64 {
	switch attr {
	case enum.AttrCool:
		return config.CoolRoomItemBonus
	case enum.AttrCute:
		return config.CuteRoomItemBonus
	case enum.AttrPassion:
		return config.PassionRoomItemBonus
	}
	return 0.0
}

// Roll returns true prob*100% of the time
func Roll(prob float64) bool {
	roll := rand.Float64()
	return roll < prob
}

// RollSafe returns true prob*100% of the time.
// Safe because the caller need to provides their own random generator
func RollSafe(prob float64, generator *rand.Rand) bool {
	roll := generator.Float64()
	return roll < prob
	// roll := fastrand.Uint32n(100000)
	// return float64(roll)/100000.0 < prob
}

func RollFast(prob float64, seed int) bool {
	// roll := fastrand.Uint32n(10000)
	roll := Uint32n(10000, seed)
	return float64(roll)/10000.0 < prob
}

// WithInt is helper function to use Math methods with int (does not check anything! make sure it is correct)
func WithInt(f func(float64) float64, val int) int {
	return int(f(float64(val)))
}

// SkillProbPotentialBonusLookup is lookup table of skill probability bonus from potential
var SkillProbPotentialBonusLookup = [11]int{0, 100, 200, 300, 400, 600, 800, 1000, 1300, 1600, 2000}

var ssrStatPotentialBonusLookup = [11]int{0, 40, 80, 120, 170, 220, 270, 320, 380, 440, 500}
var ssrLifePotentialBonusLookup = [11]int{0, 1, 2, 4, 6, 8, 10, 13, 16, 19, 22}

var srStatPotentialBonusLookup = [11]int{0, 60, 120, 180, 250, 320, 390, 460, 540, 620, 700}
var srLifePotentialBonusLookup = [11]int{0, 1, 2, 4, 6, 8, 10, 12, 14, 17, 20}

var rStatPotentialBonusLookup = [11]int{0, 60, 120, 180, 255, 330, 405, 480, 570, 660, 750}
var rLifePotentialBonusLookup = [11]int{0, 1, 2, 3, 4, 5, 6, 8, 10, 12, 14}

var nStatPotentialBonusLookup = [11]int{0, 80, 160, 250, 340, 440, 540, 650, 760, 880, 1000}
var nLifePotentialBonusLookup = [11]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 11, 13}

var unimplementedLeadSkills = []enum.LeadSkill{
	enum.LeadSkillBase,
	enum.LeadSkillIrrelevant,
	// enum.LeadSkillResonantMakeup, enum.LeadSkillResonantStep, enum.LeadSkillResonantVoice,
}

var unimplementedSkills = []enum.SkillType{
	enum.SkillTypeAlternate,
	enum.SkillTypeBase,
	// enum.SkillTypeCoolEnsemble, enum.SkillTypeCuteEnsemble, enum.SkillTypePassionEnsemble,
	enum.SkillTypeEncore,
	// acts also not implemented
}

func IsLeadSkillImplemented(ls enum.LeadSkill) bool {
	for _, lskill := range unimplementedLeadSkills {
		if lskill == ls {
			return false
		}
	}
	return true
}

func IsSkillImplemented(s enum.SkillType) bool {
	for _, skill := range unimplementedSkills {
		if skill == s {
			return false
		}
	}
	return true
}

var songDifficultyMultiplierMap = map[int]float64{
	5: 1.0, 6: 1.025, 7: 1.05, 8: 1.075, 9: 1.1,
	10: 1.2, 11: 1.225, 12: 1.25, 13: 1.275, 14: 1.3,
	15: 1.4, 16: 1.425, 17: 1.45, 18: 1.475, 19: 1.5,
	20: 1.6, 21: 1.65, 22: 1.7, 23: 1.75, 24: 1.8,
	25: 1.85, 26: 1.9, 27: 1.95, 28: 2.0, 29: 2.1,
	30: 2.2, 31: 2.3, 32: 2.4,
}

func GetSongDifficultyMultiplier(songLevel int) float64 {
	ret, ok := songDifficultyMultiplierMap[songLevel]
	if !ok {
		ret = 0
	}
	return ret
}
