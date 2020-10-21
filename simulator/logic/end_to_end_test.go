package logic

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"
)

type endToEndTestcase struct {
	name        string
	cardIds     []int
	leaderIndex int
	guestID     int
	songAttr    enum.Attribute
	expected    []bool
}

func TestLogic_All(t *testing.T) {
	assert := assert.New(t)
	testcases := []endToEndTestcase{
		{
			name:        "Current trico team",
			cardIds:     []int{200726, 300830, 300236, 300572, 200314},
			leaderIndex: 2,
			guestID:     100298,
			songAttr:    enum.AttrAll,
			expected:    []bool{true, true, true},
		},
		{
			name:        "Current cool team",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200810,
			songAttr:    enum.AttrCool,
			expected:    []bool{true, true, true},
		},
		{
			name:        "Current cool team in all song",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200810,
			songAttr:    enum.AttrAll,
			expected:    []bool{true, true, false},
		},
		{
			name:        "Current cool team in all song but with princess guest",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200538,
			songAttr:    enum.AttrAll,
			expected:    []bool{true, true, true},
		},
		{
			name:        "Tricolor lead on monocolor",
			cardIds:     []int{300280, 300236, 300432, 300148, 300096},
			leaderIndex: 0,
			guestID:     200856,
			songAttr:    enum.AttrAll,
			expected:    []bool{true, false},
		},
		{
			name:        "Resonant lead, 2color team, diff color guest (false because no guest can satisfy all condition)",
			cardIds:     []int{300712, 300830, 100536, 100344, 300280},
			leaderIndex: 0,
			guestID:     100344,
			songAttr:    enum.AttrAll,
			expected:    []bool{true, true, false},
		},
		{
			name:        "tricolor skill on monocolor",
			cardIds:     []int{300830, 300280, 300362, 300432, 300124},
			leaderIndex: 4,
			guestID:     300788,
			songAttr:    enum.AttrAll,
			expected:    []bool{false},
		},
	}
	cardHandler := NewCardLogicHandler()
	teamHandler := NewTeamLogicHandler()
	gcHandler := NewGameConfigLogicHandler()
	for i, tc := range testcases {
		if i != 6 {
			continue
		}
		i += 0
		ocardsSlice := makeOcards(tc.cardIds)
		guest := makeOcard(tc.guestID)
		song := models.NewDefaultSong("", 26, tc.songAttr, 100, 100)

		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)

		team := &usermodel.Team{Ocards: ocards, LeaderIndex: tc.leaderIndex}
		result := []bool{
			cardHandler.IsOk(ocards, &song),
			teamHandler.IsOk(team, &song),
			gcHandler.IsOk(team, guest, &song),
		}
		for i, r := range result {
			assert.Equalf(tc.expected[i], r, "Wrong for #%d handler on tc: %s", i, tc.name)
			if tc.expected[i] == r && !r {
				break
			}
		}
	}
}
