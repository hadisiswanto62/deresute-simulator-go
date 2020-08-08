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
	if want, have := ocard.Level(), uint8(90); want != have {
		t.Errorf("Error on Level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.SkillLevel(), uint8(1); want != have {
		t.Errorf("Error on skill level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.StarRank, uint8(1); want != have {
		t.Errorf("Error on Star rank field! want = %v have = %v", want, have)
	}
}

func TestRecalculate(t *testing.T) {
	card := sampleCard()
	ocard := New(&card)
	if want, have := ocard.Dance, uint16(4655); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(3837); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(7182); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(44); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	// Recalculate from SetLevel
	ocard.SetLevel(10)
	if want, have := ocard.Dance, uint16(2718); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(2241); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(4195); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(44); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	// Recalculate from SetPot*
	ocard.SetPotDance(10)
	if want, have := ocard.Dance, uint16(3218); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(2241); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(4195); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(44); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	ocard.SetPotVisual(10)
	if want, have := ocard.Dance, uint16(3218); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(2741); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(4195); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(44); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	ocard.SetPotVocal(5)
	if want, have := ocard.Dance, uint16(3218); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(2741); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(4415); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(44); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	ocard.SetPotHp(100)
	if want, have := ocard.Dance, uint16(3218); want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, uint16(2741); want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, uint16(4415); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, uint16(66); want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}

}

func BenchmarkRecalculate(b *testing.B) {
	card := sampleCard()
	ocard := New(&card)
	for i := 0; i < b.N; i++ {
		ocard.recalculate()
	}
}
