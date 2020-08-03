package cardmanager

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func TestAttribute(t *testing.T) {
	cm := Default()
	for _, attr := range enum.AllIdolAttributes {
		card := cm.Filter().Attribute(attr).First()
		if card.Attribute != attr {
			t.Errorf("Incorrect card attribute! %v != %v", card.Attribute, attr)
		}
	}
}

func TestID(t *testing.T) {
	cm := Default()
	testcases := [5]int{
		100001,
		200609,
		300706,
		200195,
		100064,
	}
	for _, id := range testcases {
		card := cm.Filter().ID(id).First()
		if card.ID != id {
			t.Errorf("Incorrect card id! %v != %v", card.ID, id)
		}
	}
}
