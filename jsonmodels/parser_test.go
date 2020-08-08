package jsonmodels

import (
	"os"
	"testing"
)

func init() {
	os.Chdir("../")
}

func TestParse(t *testing.T) {
	dp := JSONDataParser{}
	cards, err := dp.Parse()
	if err != nil {
		t.Errorf("Test failed! %v", err)
	}
	if cards == nil {
		t.Errorf("Card parsing failed!")
	}
	// for _, card := range cards {
	// 	if card.ID < 100100 {
	// 		fmt.Printf("%d: %p %v\n", card.ID, &card.Rarity.Rarity, *(card.Rarity))
	// 	}
	// }
}
