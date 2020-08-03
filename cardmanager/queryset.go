package cardmanager

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

// QuerySet for filtering
type QuerySet struct {
	cards []models.Card
}

// Attribute filters current cards by attribute
func (q *QuerySet) Attribute(attr enum.Attribute) *QuerySet {
	result := []models.Card{}
	for i := range q.cards {
		if q.cards[i].Attribute == attr {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// ID filters current cards by ID
func (q *QuerySet) ID(id int) *QuerySet {
	result := []models.Card{}
	for i := range q.cards {
		if q.cards[i].ID == id {
			result = append(result, q.cards[i])
			break
		}
	}
	q.cards = result
	return q
}

// Rarity filters current cards by rarity
// func (q *QuerySet) Rarity(rarity enum.Rarity) *QuerySet {
// 	result := []models.Card{}
// 	for i := range q.cards {
// 		if q.cards[i].Rarity == rarity {
// 			result = append(result, q.cards[i])
// 		}
// 	}
// 	q.cards = result
// 	return q
// }

// Get gets all cards that matches current filter
func (q *QuerySet) Get() []models.Card {
	return q.cards
}

// First gets the first card that matches current filter
func (q *QuerySet) First() models.Card {
	return q.cards[0]
}
