package enum

// Rarity represents card's rarity
type Rarity string

// All rarities
const (
	RaritySSR Rarity = "SSR"
	RaritySR  Rarity = "SR"
	RarityR   Rarity = "R"
	RarityN   Rarity = "N"
)

// AllRarities is all valid rarities
var AllRarities = [4]Rarity{RaritySSR, RaritySR, RarityR, RarityN}

func GetRarity(s string) Rarity {
	for _, rarity := range AllRarities {
		if string(rarity) == s {
			return rarity
		}
	}
	return RaritySSR
}
