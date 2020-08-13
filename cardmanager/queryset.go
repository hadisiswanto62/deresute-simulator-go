package cardmanager

import (
	"strings"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

// QuerySet for filtering
type QuerySet struct {
	cards []*models.Card
}

// Attribute filters current cards by attribute
func (q *QuerySet) Attribute(attr enum.Attribute) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].Idol.Attribute == attr {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// ID filters current cards by ID
func (q *QuerySet) ID(id int) *QuerySet {
	result := []*models.Card{}
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
func (q *QuerySet) Rarity(rarity enum.Rarity) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].Rarity.Rarity == rarity {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// IsEvolved filters current cards by evolved status
func (q *QuerySet) IsEvolved(evolveStatus bool) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].Rarity.IsEvolved == evolveStatus {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// LeadSkill filters current cards by lead skill
func (q *QuerySet) LeadSkill(leadSkill enum.LeadSkill) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].LeadSkill.Name == leadSkill {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// SkillType filters current cards by skill type
func (q *QuerySet) SkillType(skillType enum.SkillType) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].Skill.SkillType.Name == skillType {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// Name filters current cards by the idol's Name
func (q *QuerySet) Name(name string) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if q.cards[i].Idol.Name == name {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// NameLike filters current cards by the idol's Name (substring allowed, case insensitive).
// It is very slow so use QuerySet.Name() if possible.
func (q *QuerySet) NameLike(name string) *QuerySet {
	result := []*models.Card{}
	for i := range q.cards {
		if strings.Contains(
			strings.ToLower(q.cards[i].Idol.Name),
			strings.ToLower(name),
		) {
			result = append(result, q.cards[i])
		}
	}
	q.cards = result
	return q
}

// Get gets all cards that matches current filter
func (q *QuerySet) Get() []*models.Card {
	return q.cards
}

// First gets the first card that matches current filter
func (q *QuerySet) First() *models.Card {
	return q.cards[0]
}
