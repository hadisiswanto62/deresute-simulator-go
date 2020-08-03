package parser

import "testing"

func TestSimplify(t *testing.T) {
	cards := SimplifyCardsData()
	if len(cards) == 0 {
		t.Errorf("Parsing error!")
	}
}
