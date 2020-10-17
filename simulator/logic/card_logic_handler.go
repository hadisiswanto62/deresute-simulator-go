package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// cardLogicHandler handles logic for determining whether card/cards is good
type cardLogicHandler struct {
	cardLogics  []*cardLogic
	cardsLogics []*cardsLogic

	cardOkCache map[int]bool
}

var cardInstance *cardLogicHandler

// NewCardLogicHandler creates a new gameConfigLogicHandler instance
func NewCardLogicHandler() *cardLogicHandler {
	if cardInstance == nil {
		cardInstance = makeCardLogicHandler()
	}
	return cardInstance
}

func makeCardLogicHandler() *cardLogicHandler {
	cardLogics := []*cardLogic{
		cardCardIsSSR,
		cardSkillIsImplemented,
		// add cardCorrectColor if needed
	}
	cardsLogics := []*cardsLogic{
		motifWithHighCorrectStat,
	}

	if !helper.Features.UseConcentration() {
		cardLogics = append(cardLogics, cardSkillIsNotConcentration)
	}

	if helper.Features.DebugNoLogic() {
		cardLogics = []*cardLogic{}
		cardsLogics = []*cardsLogic{}
	}
	return &cardLogicHandler{
		cardLogics:  cardLogics,
		cardsLogics: cardsLogics,
		cardOkCache: make(map[int]bool),
	}
}

// IsOk checks whether cards and song is ok
func (clh cardLogicHandler) IsOk(cards [5]*usermodel.OwnedCard, song *models.Song) bool {
	// handle single card
	for _, card := range cards {
		valid, isCached := clh.cardOkCache[card.Card.ID]
		if !isCached {
			valid = true
			for _, logic := range clh.cardLogics {
				if !logic.isSatisfied(card, song) {
					valid = false
					break
				}
			}
			clh.cardOkCache[card.Card.ID] = valid
		}
		if !valid {
			return false
		}
	}

	// handle 5 cards as a whole
	for _, logic := range clh.cardsLogics {
		if !logic.isSatisfied(cards, song) {
			return false
		}
	}
	return true
}
