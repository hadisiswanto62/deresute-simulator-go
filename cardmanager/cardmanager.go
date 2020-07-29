package cardmanager

import (
	"example.com/deresute/testo/models"
	"example.com/deresute/testo/parser"
)

// CardManager manages cards
type CardManager struct {
	Cards []models.Card
}

// NewDefault returns default CardManager (card data from parser.AllCardsDir)
func NewDefault() CardManager {
	return CardManager{parser.Parse()}
}
