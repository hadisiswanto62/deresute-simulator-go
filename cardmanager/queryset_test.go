package cardmanager

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/stretchr/testify/assert"
)

var cm *CardManager

func init() {
	cm, _ = Default()
}

func TestAttribute(t *testing.T) {
	for _, attr := range enum.AllIdolAttributes {
		card := cm.Filter().Attribute(attr).First()
		if card.Idol.Attribute != attr {
			t.Errorf("Incorrect card attribute! %v != %v", card.Idol.Attribute, attr)
		}
	}
}

func TestRarity(t *testing.T) {
	assert := assert.New(t)
	for _, rarity := range enum.AllRarities {
		card := cm.Filter().Rarity(rarity).First()
		assert.Equal(card.Rarity.Rarity, rarity, "Incorrect rarity!")
	}
}

func TestEvolved(t *testing.T) {
	assert := assert.New(t)
	for _, evolveStatus := range [2]bool{true, false} {
		card := cm.Filter().IsEvolved(evolveStatus).First()
		assert.Equal(card.Rarity.IsEvolved, evolveStatus, "Incorrect evolve status!")
	}
}

func TestID(t *testing.T) {
	assert := assert.New(t)
	testcases := [5]int{
		100001,
		200609,
		300706,
		200195,
		100064,
	}
	for _, id := range testcases {
		card := cm.Filter().ID(id).First()
		assert.Equal(card.ID, id, "Incorrect ID")
	}
}
