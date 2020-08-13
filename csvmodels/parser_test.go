package csvmodels

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("../")
}

func TestParseOwnedCard(t *testing.T) {
	p := CSVDataParser{}
	ocds, err := p.ParseOwnedCardRawData("")
	if err != nil {
		t.Errorf("Test failed! %v", err)
	}
	for _, ocd := range ocds {
		assert.NotNilf(t, ocd, "Card not created!")
		assert.NotEqualf(t, ocd.CardID, 0, "Card data not correctly parsed!")
	}
}
