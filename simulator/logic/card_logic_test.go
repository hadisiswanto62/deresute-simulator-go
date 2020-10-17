package logic

import (
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

var cm *cardmanager.CardManager

func init() {
	os.Chdir("../../")
	var err error
	cm, err = cardmanager.Default()
	if err != nil {
		panic(err)
	}
}

func TestCardLogic_CardIsSSR(t *testing.T) {
	testcases := []struct {
		rarity   enum.Rarity
		expected bool
	}{
		{rarity: enum.RaritySSR, expected: true},
		{rarity: enum.RaritySR, expected: false},
		{rarity: enum.RarityR, expected: false},
		{rarity: enum.RarityN, expected: false},
	}
	logic := cardCardIsSSR
	for _, tc := range testcases {
		card := cm.Filter().Rarity(tc.rarity).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		actual := logic.isSatisfied(ocard, nil)
		assert.Equalf(t, tc.expected, actual, "Wrong result for rarity %s", tc.rarity)
	}
}

func TestCardLogic_SkillIsNotConcentration(t *testing.T) {
	testcases := []struct {
		nameID   string
		expected bool
	}{
		{nameID: "yoshino4", expected: true},
		{nameID: "uzuki4", expected: true},
		{nameID: "megumi1", expected: false},
	}
	logic := cardSkillIsNotConcentration
	for _, tc := range testcases {
		card := cm.Filter().SsrNameID(tc.nameID).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		actual := logic.isSatisfied(ocard, nil)
		assert.Equalf(t, tc.expected, actual, "Wrong result for nameID %s", tc.nameID)
	}
}

func TestCardLogic_SkillIsImplemented(t *testing.T) {
	testcases := []struct {
		nameID   string
		expected bool
	}{
		{nameID: "yoshino4", expected: true},
		{nameID: "uzuki4", expected: true},
		{nameID: "megumi1", expected: true},
		{nameID: "nagi2", expected: false},
	}
	logic := cardSkillIsImplemented
	for _, tc := range testcases {
		card := cm.Filter().SsrNameID(tc.nameID).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		// fmt.Println(ocard)
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		actual := logic.isSatisfied(ocard, nil)
		assert.Equalf(t, tc.expected, actual, "Wrong result for nameID %s", tc.nameID)
	}
}

func TestCardLogic_MotifWithHighCorrectStat(t *testing.T) {
	testcases := []struct {
		name     string
		cardIDs  []int
		expected bool
	}{
		{
			name:     "vocal motif on vocal team",
			cardIDs:  []int{300711, 300709, 300717, 300743, 300685},
			expected: true,
		},
		{
			name:     "visual motif on vocal team",
			cardIDs:  []int{300783, 300709, 300717, 300743, 300685},
			expected: false,
		},
		{
			name:     "visual and vocal motif on vocal team",
			cardIDs:  []int{300783, 300711, 300717, 300743, 300685},
			expected: false,
		},
		{
			name:     "vocal motif on average team",
			cardIDs:  []int{300711, 300783, 300077, 300125, 300811},
			expected: false,
		},
	}
	logic := motifWithHighCorrectStat
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIDs)
		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)
		actual := logic.isSatisfied(ocards, nil)
		assert.Equalf(t, tc.expected, actual, "Wrong result on tc %s", tc.name)
	}
}
