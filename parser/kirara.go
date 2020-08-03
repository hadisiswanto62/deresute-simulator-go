package parser

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

const (
	rawAllCardsDir = "resources/all_cards.json"
	// AllCardsDir is directory where clean Cards data is stored
	AllCardsDir = "data/all_cards.json"
)

// SimplifyCardsData parses raw cards data into clean cards data
func SimplifyCardsData() []models.Card {
	filename := rawAllCardsDir
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cards []models.Card
	err = json.Unmarshal(text, &cards)
	if err != nil {
		panic(err)
	}

	for i := range cards {
		switch cards[i].TmpRarity.Value {
		case 1, 2:
			cards[i].Rarity = enum.RarityN
		case 3, 4:
			cards[i].Rarity = enum.RarityR
		case 5, 6:
			cards[i].Rarity = enum.RaritySR
		case 7, 8:
			cards[i].Rarity = enum.RaritySSR
		}
		cards[i].IsEvolved = cards[i].TmpRarity.Value%2 == 0
		cards[i].MaxLevel = cards[i].Rarity.BaseMaxLevel
		if cards[i].IsEvolved {
			cards[i].MaxLevel += 10
		}
	}
	return cards
}

// SaveSimplifiedCards save cards data to parser.AllCardsDir
func SaveSimplifiedCards(cards []models.Card) {
	data, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		panic(err)
	}

	text := []byte(data)
	err = ioutil.WriteFile(AllCardsDir, text, 0644)
	if err != nil {
		panic(err)
	}
}
