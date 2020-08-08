package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Rarity represents rarity of a card
type Rarity struct {
	ID        uint8
	Rarity    enum.Rarity `json:"rarity"`
	IsEvolved bool
	MaxLevel  uint8
}
