package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Rarity represents rarity of a card
type Rarity struct {
	ID        int
	Rarity    enum.Rarity `json:"rarity"`
	IsEvolved bool
	MaxLevel  int
}
