package parser

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

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
