package cardmanager

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/jsonmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

// CardManager manages cards
type CardManager struct {
	Cards    []*models.Card
	querySet *QuerySet
}

type dataParser interface {
	Parse() ([]*models.Card, error)
}

// Filter returns pointer to QuerySet (used for filtering cards)
func (cm *CardManager) Filter() *QuerySet {
	cm.querySet.cards = cm.Cards
	return cm.querySet
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
		instance = &CardManager{cards, &QuerySet{cards}}
	}
	return instance, nil
}

var instance *CardManager
