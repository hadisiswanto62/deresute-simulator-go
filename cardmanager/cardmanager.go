package cardmanager

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/jsonmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

// CardManager manages cards
type CardManager struct {
	Cards []models.Card
}

type dataParser interface {
	Parse() ([]models.Card, error)
}

// Filter returns pointer to QuerySet (used for filtering cards)
func (cm CardManager) Filter() *QuerySet {
	return &QuerySet{cm.Cards}
}

// Default returns pointer to default CardManager (card data from parser.AllCardsDir).
// It uses jsonmodels.JSONDataParser by default
func Default() (*CardManager, error) {
	var dp dataParser
	dp = jsonmodels.JSONDataParser{}
	if instance == nil {
		cards, err := dp.Parse()
		if err != nil {
			return nil, fmt.Errorf("cannot parse cards: %v", err)
		}
		instance = &CardManager{cards}
	}
	return instance, nil
}

var instance *CardManager
