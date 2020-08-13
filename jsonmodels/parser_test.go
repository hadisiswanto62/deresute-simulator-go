package jsonmodels

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestParseOwnedCardRawData(t *testing.T) {
	assertion := assert.New(t)
	dp := JSONDataParser{}
	ocds, err := dp.ParseOwnedCardRawData("")
	assertion.Nilf(err, "Test failed: %v", err)
	for _, ocd := range ocds {
		fmt.Println(ocd)
	}
}
