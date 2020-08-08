package enum

// Rarity represents card's rarity
type Rarity string

// All rarities
var (
	RaritySSR Rarity = "SSR"
	RaritySR  Rarity = "SR"
	RarityR   Rarity = "R"
	RarityN   Rarity = "N"
)

// AllRarities is all valid rarities
var AllRarities = [4]Rarity{RaritySSR, RaritySR, RarityR, RarityN}
