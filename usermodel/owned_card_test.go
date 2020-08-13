package usermodel

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/stretchr/testify/assert"
)

func sampleCard() models.Card {
	idol := models.Idol{
		ID:        181,
		Name:      "Mifune Miyu",
		Attribute: enum.AttrCool,
	}
	rarity := models.Rarity{
		ID:        6,
		Rarity:    enum.RaritySR,
		IsEvolved: true,
		MaxLevel:  70,
	}
	skillType, _ := models.GetSkillType("Perfect Score Bonus")
	skill := &models.Skill{
		ID:           200803,
		Timer:        13,
		ProcChance:   [2]int{4000, 6000},
		EffectLength: [2]int{600, 900},
		SkillType:    skillType,
	}
	leadSkill, _ := models.GetLeadSkill("クールボイス")
	return models.Card{
		ID:        200804,
		SeriesID:  200803,
		Idol:      &idol,
		Rarity:    &rarity,
		LeadSkill: leadSkill,
		Skill:     skill,

		BonusDance:  143,
		BonusHp:     2,
		BonusVisual: 119,
		BonusVocal:  224,
		DanceMax:    3517,
		DanceMin:    1913,
		HpMax:       37,
		HpMin:       37,
		VisualMax:   2906,
		VisualMin:   1583,
		VocalMax:    5501,
		VocalMin:    2993,
	}
}

func sampleOwnedCard() *OwnedCard {
	card := sampleCard()
	return NewOwnedCard(&card)
}

func TestBatchCreate(t *testing.T) {
	var cards []*models.Card
	length := 10
	for i := 0; i < length; i++ {
		card := sampleCard()
		cards = append(cards, &card)
	}
	ocards := BatchNewOwnedCards(cards, 1, 1, 0, 0, 0, 0, 0)
	assert.Equal(t, len(ocards), length, "Failed to batch create!")
}

func TestCreate(t *testing.T) {
	card := sampleCard()
	ocard := NewOwnedCard(&card)
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
	ocard := sampleOwnedCard()
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
	ocard := sampleOwnedCard()
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

func TestInvalidValues(t *testing.T) {
	assertion := assert.New(t)
	ocard := sampleOwnedCard()
	card := ocard.Card
	// Invalid < value
	request := OwnedCardRequest{
		Card:       card,
		Level:      0,
		SkillLevel: 0,
		StarRank:   -1,
		PotVisual:  -1,
		PotDance:   -1,
		PotVocal:   -1,
		PotHp:      -1,
		PotSkill:   -1,
	}
	ocard = NewOwnedCard2(request)
	assertion.Equal(ocard.Rarity.MaxLevel, ocard.level, "Invalid level is accepted!")
	assertion.Equal(1, ocard.skillLevel, "Invalid skill level is accepted!")
	assertion.Equal(1, ocard.StarRank, "Invalid star rank is accepted!")
	assertion.Equal(0, ocard.potHp, "Invalid pothp is accepted!")
	assertion.Equal(0, ocard.potVisual, "Invalid potvisual is accepted!")
	assertion.Equal(0, ocard.potDance, "Invalid potdance is accepted!")
	assertion.Equal(0, ocard.potVocal, "Invalid potvocal is accepted!")
	assertion.Equal(0, ocard.potSkill, "Invalid potskill is accepted!")

	// valid < value
	request = OwnedCardRequest{
		Card:       card,
		Level:      1,
		SkillLevel: 1,
		StarRank:   1,
		PotVisual:  0,
		PotDance:   0,
		PotVocal:   0,
		PotHp:      0,
		PotSkill:   0,
	}
	ocard = NewOwnedCard2(request)
	assertion.Equal(1, ocard.level, "Invalid level is accepted!")
	assertion.Equal(1, ocard.skillLevel, "Invalid skill level is accepted!")
	assertion.Equal(1, ocard.StarRank, "Invalid star rank is accepted!")
	assertion.Equal(0, ocard.potHp, "Invalid pothp is accepted!")
	assertion.Equal(0, ocard.potVisual, "Invalid potvisual is accepted!")
	assertion.Equal(0, ocard.potDance, "Invalid potdance is accepted!")
	assertion.Equal(0, ocard.potVocal, "Invalid potvocal is accepted!")
	assertion.Equal(0, ocard.potSkill, "Invalid potskill is accepted!")

	// valid > value
	request = OwnedCardRequest{
		Card:       card,
		Level:      card.Rarity.MaxLevel,
		SkillLevel: 10,
		StarRank:   20,
		PotVisual:  10,
		PotDance:   10,
		PotVocal:   10,
		PotHp:      10,
		PotSkill:   10,
	}
	ocard = NewOwnedCard2(request)
	assertion.Equal(ocard.Card.Rarity.MaxLevel, ocard.level, "Invalid level is accepted!")
	assertion.Equal(10, ocard.skillLevel, "Invalid skill level is accepted!")
	assertion.Equal(20, ocard.StarRank, "Invalid star rank is accepted!")
	assertion.Equal(10, ocard.potHp, "Invalid pothp is accepted!")
	assertion.Equal(10, ocard.potVisual, "Invalid potvisual is accepted!")
	assertion.Equal(10, ocard.potDance, "Invalid potdance is accepted!")
	assertion.Equal(10, ocard.potVocal, "Invalid potvocal is accepted!")
	assertion.Equal(10, ocard.potSkill, "Invalid potskill is accepted!")

	// Invalid > value
	request = OwnedCardRequest{
		Card:       card,
		Level:      card.Rarity.MaxLevel + 1,
		SkillLevel: 11,
		StarRank:   21,
		PotVisual:  11,
		PotDance:   11,
		PotVocal:   11,
		PotHp:      11,
		PotSkill:   11,
	}
	ocard = NewOwnedCard2(request)
	assertion.Equal(ocard.Card.Rarity.MaxLevel, ocard.level, "Invalid level is accepted!")
	assertion.Equal(10, ocard.skillLevel, "Invalid skill level is accepted!")
	assertion.Equal(20, ocard.StarRank, "Invalid star rank is accepted!")
	assertion.Equal(10, ocard.potHp, "Invalid pothp is accepted!")
	assertion.Equal(10, ocard.potVisual, "Invalid potvisual is accepted!")
	assertion.Equal(10, ocard.potDance, "Invalid potdance is accepted!")
	assertion.Equal(10, ocard.potVocal, "Invalid potvocal is accepted!")
	assertion.Equal(10, ocard.potSkill, "Invalid potskill is accepted!")

	// setLevel(0) works
	ocard.SetLevel(0)
	assertion.Equal(1, ocard.level, "Invlaid level is accepted!")
}

func BenchmarkRecalculate(b *testing.B) {
	ocard := sampleOwnedCard()
	for i := 0; i < b.N; i++ {
		ocard.recalculate()
	}
}
