package helper

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Scale scales from min to max (max-min is divided into `steps`-1 parts evenly) (in uint16)
func Scale(min, max uint16, maxLevel, currentLevel uint8) uint16 {
	step := float64(max-min) / (float64(maxLevel) - 1)
	add := step * float64(currentLevel-1)
	x := math.Floor(float64(min) + add)
	return uint16(x)
}

// StatPotentialBonusLookupFor returns lookup table of stat bonus from potential for the specified rarity
func StatPotentialBonusLookupFor(rarity enum.Rarity) [11]uint16 {
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
	return [11]uint16{}
}

// LifePotentialBonusLookupFor returns lookup table of life bonus from potential for the specified rarity
func LifePotentialBonusLookupFor(rarity enum.Rarity) [11]uint16 {
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
	return [11]uint16{}
}

// SkillProbPotentialBonusLookup is lookup table of skill probability bonus from potential
var SkillProbPotentialBonusLookup = [11]uint16{0, 100, 200, 300, 400, 600, 800, 1000, 1300, 1600, 2000}

var ssrStatPotentialBonusLookup = [11]uint16{0, 40, 80, 120, 170, 220, 270, 320, 380, 440, 500}
var ssrLifePotentialBonusLookup = [11]uint16{0, 1, 2, 4, 6, 8, 10, 13, 16, 19, 22}

var srStatPotentialBonusLookup = [11]uint16{0, 60, 120, 180, 250, 320, 390, 460, 540, 620, 700}
var srLifePotentialBonusLookup = [11]uint16{0, 1, 2, 4, 6, 8, 10, 12, 14, 17, 20}

var rStatPotentialBonusLookup = [11]uint16{0, 60, 120, 180, 255, 330, 405, 480, 570, 660, 750}
var rLifePotentialBonusLookup = [11]uint16{0, 1, 2, 3, 4, 5, 6, 8, 10, 12, 14}

var nStatPotentialBonusLookup = [11]uint16{0, 80, 160, 250, 340, 440, 540, 650, 760, 880, 1000}
var nLifePotentialBonusLookup = [11]uint16{0, 1, 2, 3, 4, 5, 6, 7, 9, 11, 13}
