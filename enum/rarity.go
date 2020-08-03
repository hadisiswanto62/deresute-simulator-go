package enum

// Rarity represents card's rarity
type Rarity struct {
	Value        string
	BaseMaxLevel int8
}

// All rarities
var (
	RaritySSR = Rarity{"ssr", 80}
	RaritySR  = Rarity{"sr", 60}
	RarityR   = Rarity{"r", 40}
	RarityN   = Rarity{"n", 20}
)

// AllRarities is all valid rarities
var AllRarities = [4]Rarity{RaritySSR, RaritySR, RarityR, RarityN}
