package models

import (
	"testing"
)

func TestCreate(t *testing.T) {
	card := sampleCard()
	ocard := New(&card)
	if want, have := ocard.Card, &card; want != have {
		t.Errorf("Error on Card field! want = %v have = %v", want, have)
	}
	if want, have := ocard.Level(), int8(90); want != have {
		t.Errorf("Error on Level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.SkillLevel(), int8(1); want != have {
		t.Errorf("Error on skill level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.StarRank, int8(1); want != have {
		t.Errorf("Error on Star rank field! want = %v have = %v", want, have)
	}
}
