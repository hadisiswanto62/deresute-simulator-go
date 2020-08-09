package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	card := sampleCard()
	ocard := New(&card)
	if want, have := ocard.Card, &card; want != have {
		t.Errorf("Error on Card field! want = %v have = %v", want, have)
	}
	if want, have := ocard.Level(), 70; want != have {
		t.Errorf("Error on Level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.SkillLevel(), 1; want != have {
		t.Errorf("Error on skill level field! want = %v have = %v", want, have)
	}
	if want, have := ocard.StarRank, 1; want != have {
		t.Errorf("Error on Star rank field! want = %v have = %v", want, have)
	}
}

func TestRecalculate(t *testing.T) {
	card := sampleCard()
	ocard := New(&card)
	if want, have := ocard.Dance, 3517+143; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 2906+119; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 5501+224; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 39; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	// Recalculate from SetLevel
	ocard.SetLevel(10)
	if want, have := ocard.Dance, 2122+143; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 1755+119; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 3320+224; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 39; want != have {
		t.Errorf("Wrong Hp value! want = %d have = %d", want, have)
	}
	// Recalculate from SetPot*
	ocard.SetPotDance(10)
	if want, have := ocard.Dance, 2122+143+700; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 1755+119; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 3320+224; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 39; want != have {
		t.Errorf("Wrong Hp value! want = %d have = %d", want, have)
	}
	ocard.SetPotVisual(10)
	if want, have := ocard.Dance, 2122+143+700; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 1755+119+700; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 3320+224; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 39; want != have {
		t.Errorf("Wrong Hp value! want = %d have = %d", want, have)
	}
	ocard.SetPotVocal(5)
	if want, have := ocard.Dance, 2122+143+700; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 1755+119+700; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 3320+224+320; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 39; want != have {
		t.Errorf("Wrong Hp value! want = %d have = %d", want, have)
	}
	ocard.SetPotHp(100)
	if want, have := ocard.Dance, 2122+143+700; want != have {
		t.Errorf("Wrong Dance value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Visual, 1755+119+700; want != have {
		t.Errorf("Wrong Visual value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Vocal, 3320+224+320; want != have {
		t.Errorf("Wrong Vocal value! want = %d have = %d", want, have)
	}
	if want, have := ocard.Hp, 59; want != have {
		t.Errorf("Wrong Hp value! want = %d have = %d", want, have)
	}
}

func TestRecalculateSkill(t *testing.T) {
	assert := assert.New(t)
	card := sampleCard()
	ocard := New(&card)
	assert.Equal(ocard.SkillLevel(), 1, "Default skill level is not 1!")
	assert.Equal(ocard.SkillProcChance, 4000, "Default skill level is not 1!")
	assert.Equal(ocard.SkillEffectLength, 600, "Default skill level is not 1!")

	ocard.SetSkillLevel(99)
	assert.Equal(ocard.SkillLevel(), 10, "Wrong skill level value!")
	assert.Equal(ocard.SkillProcChance, 6000, "Wrong skill proc chance!")
	assert.Equal(ocard.SkillEffectLength, 900, "Wrong skill effect length!")

	ocard.SetPotSkill(99)
	assert.Equal(ocard.SkillLevel(), 10, "Wrong skill level value!")
	assert.Equal(ocard.SkillProcChance, 8000, "Wrong skill proc chance!")
	assert.Equal(ocard.SkillEffectLength, 900, "Wrong skill effect length!")

}

func BenchmarkRecalculate(b *testing.B) {
	card := sampleCard()
	ocard := New(&card)
	for i := 0; i < b.N; i++ {
		ocard.recalculate()
	}
}
