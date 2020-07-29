package parser

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

const (
	rawAllCardsDir = "resources/all_cards.json"
	// AllCardsDir is directory where clean Cards data is stored
	AllCardsDir = "data/all_cards.json"
)

// SimplifyCardsData parses raw cards data into clean cards data in AllCardsDir
func SimplifyCardsData() {
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

	// for i := range cards {
	// 	cards[i].Field = func(a, b int) int {
	// 		return a + b
	// 	}
	// }

	data, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		panic(err)
	}

	text = []byte(data)
	err = ioutil.WriteFile(AllCardsDir, text, 0644)
	if err != nil {
		panic(err)
	}
}

// Parse parses clean cards data into slice of Cards
func Parse() []models.Card {
	filename, err := filepath.Abs(AllCardsDir)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cards []models.Card
	if err = json.Unmarshal(text, &cards); err != nil {
		panic(err)
	}
	return cards
}
