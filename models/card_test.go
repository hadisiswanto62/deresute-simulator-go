package models

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func sampleCard() Card {
	idol := Idol{
		ID:        233,
		Name:      "Takafuji Kako",
		Attribute: enum.AttrCool,
	}
	rarity := Rarity{
		ID:        8,
		Rarity:    enum.RaritySSR,
		IsEvolved: true,
		MaxLevel:  90,
	}
	return Card{
		ID:          200698,
		SeriesID:    200697,
		Idol:        &idol,
		Rarity:      &rarity,
		IsEvolved:   true,
		MaxLevel:    90,
		BonusDance:  184,
		BonusHp:     2,
		BonusVisual: 152,
		BonusVocal:  284,
		DanceMax:    4471,
		DanceMin:    2317,
		HpMax:       42,
		HpMin:       42,
		VisualMax:   3685,
		VisualMin:   1910,
		VocalMax:    6898,
		VocalMin:    3575,
	}
}

func BenchmarkRarityMember(b *testing.B) {
	card := sampleCard()
	for i := 0; i < b.N; i++ {
		a := card.Rarity
		_ = a
	}
}
