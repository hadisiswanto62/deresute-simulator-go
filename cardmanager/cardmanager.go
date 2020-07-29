package cardmanager

import (
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/parser"
)

// CardManager manages cards
type CardManager struct {
	Cards []models.Card
}

// NewDefault returns default CardManager (card data from parser.AllCardsDir)
func NewDefault() CardManager {
	return CardManager{parser.Parse()}
}
