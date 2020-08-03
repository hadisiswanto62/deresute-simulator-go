package cardmanager

import (
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/parser"
)

// CardManager manages cards
type CardManager struct {
	Cards []models.Card
}

// Filter returns pointer to QuerySet (used for filtering cards)
func (cm CardManager) Filter() *QuerySet {
	return &QuerySet{cm.Cards}
}

// Default returns pointer to default CardManager (card data from parser.AllCardsDir)
func Default() *CardManager {
	if instance == nil {
		instance = &CardManager{parser.Parse()}
	}
	return instance
}

var instance *CardManager
