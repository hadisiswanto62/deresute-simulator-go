package logic

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"
)

type cardsLogicTestcase struct {
	name     string
	cardIDs  []int
	songAttr enum.Attribute
	expected bool
}

func testCardsLogic(t *testing.T, testcases []cardsLogicTestcase, logic *cardsLogic) bool {
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIDs)
		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)

		var song *models.Song
		if tc.songAttr != "" {
			tmp := models.NewDefaultSong("", 26, tc.songAttr, 100, 100)
			song = &tmp
		}

		actual := logic.isSatisfied(ocards, song)
		result := assert.Equalf(t, tc.expected, actual, "Wrong result on tc %s", tc.name)
		if !result {
			return result
		}
	}
	return true
}

func TestCardLogicHandler_IsOK(t *testing.T) {
	testcases := []cardsLogicTestcase{
		{
			name:     "Current trico team",
			cardIDs:  []int{200726, 300830, 300236, 300572, 200314},
			songAttr: enum.AttrAll,
			expected: true,
		},
		{
			name:     "Current cool team",
			cardIDs:  []int{200726, 200378, 200624, 200584, 200314},
			songAttr: enum.AttrCool,
			expected: true,
		},
		{
			name:     "One invalid card",
			cardIDs:  []int{200726, 100002, 200624, 200584, 200314},
			songAttr: enum.AttrCool,
			expected: false,
		},
	}
	handler := NewCardLogicHandler()
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIDs)
		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)

		var song *models.Song
		if tc.songAttr != "" {
			tmp := models.NewDefaultSong("", 26, tc.songAttr, 100, 100)
			song = &tmp
		}

		actual := handler.IsOk(ocards, song)

		assert.Equalf(t, tc.expected, actual, "Wrong result for tc: %s", tc.name)
	}
}

func TestCardLogic_MotifWithHighCorrectStat(t *testing.T) {
	testcases := []cardsLogicTestcase{
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
			cardIDs:  []int{300783, 300712, 100500, 200572, 300686},
			expected: false,
		},
		{
			name:     "vocal motif on average team",
			cardIDs:  []int{300711, 300783, 300077, 300125, 300811},
			expected: false,
		},
	}
	logic := motifWithHighCorrectStat
	testCardsLogic(t, testcases, logic)
}

func TestCardLogic_UseUnevolvedWithoutEvolved(t *testing.T) {
	testcases := []cardsLogicTestcase{
		{
			name:     "correct",
			cardIDs:  []int{300712, 300710, 300718, 300744, 300743},
			expected: true,
		},
		{
			name:     "incorrect",
			cardIDs:  []int{300712, 300710, 300718, 300746, 300743},
			expected: false,
		},
		{
			name:     "irrelevant",
			cardIDs:  []int{300712, 300710, 300718, 300744, 300686},
			expected: true,
		},
	}
	logic := useUnevolvedWithoutEvolved
	testCardsLogic(t, testcases, logic)
}
