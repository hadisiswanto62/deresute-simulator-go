package logic

import (
	"fmt"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/models"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func makeOcard(ID int) *usermodel.OwnedCard {
	card := cm.Filter().ID(ID).First()
	ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
	if err != nil {
		panic(fmt.Errorf("cannot create ocard: %v", err))
	}
	return ocard
}

func makeOcards(IDs []int) []*usermodel.OwnedCard {
	ret := []*usermodel.OwnedCard{}
	for _, ID := range IDs {
		ret = append(ret, makeOcard(ID))
	}
	return ret
}

type gameConfigLogicTestcase struct {
	name        string
	cardIds     []int
	leaderIndex int
	guestID     int
	songAttr    enum.Attribute
	expected    bool
}

func testGameConfigLogic(t *testing.T, testcases []gameConfigLogicTestcase, logic *gameConfigLogic) bool {
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIds)
		guest := makeOcard(tc.guestID)
		song := models.NewDefaultSong("", 26, tc.songAttr, 100, 100)

		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)

		team := &usermodel.Team{Ocards: ocards, LeaderIndex: tc.leaderIndex}
		activable := []*usermodel.OwnedCard{
			team.Leader(), guest,
		}

		actual := logic.isViolated(team, activable, guest, &song)
		result := assert.Equalf(t, tc.expected, actual, "Wrong result for tc %s", tc.name)
		if !result {
			return result
		}
	}
	return true
}

func TestGameConfigHandler_IsOK(t *testing.T) {
	testcases := []gameConfigLogicTestcase{
		{
			name:        "Current trico team",
			cardIds:     []int{200726, 300830, 300236, 300572, 200314},
			leaderIndex: 2,
			guestID:     100298,
			songAttr:    enum.AttrAll,
			expected:    true,
		},
		{
			name:        "Current cool team",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200810,
			songAttr:    enum.AttrCool,
			expected:    true,
		},
		{
			name:        "Current cool team in all song",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200810,
			songAttr:    enum.AttrAll,
			expected:    false,
		},
		{
			name:        "Current cool team in all song but with princess guest",
			cardIds:     []int{200726, 200378, 200624, 200584, 200314},
			leaderIndex: 4,
			guestID:     200538,
			songAttr:    enum.AttrAll,
			expected:    true,
		},
	}
	handler := NewGameConfigLogicHandler()
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIds)
		guest := makeOcard(tc.guestID)
		song := models.NewDefaultSong("", 26, tc.songAttr, 100, 100)

		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)

		team := &usermodel.Team{Ocards: ocards, LeaderIndex: tc.leaderIndex}
		assert.Equalf(t, tc.expected, handler.IsOk(team, guest, &song), "Wrong for tc: %s", tc.name)
	}
}

func TestGameConfig_UnisonInCorrectSongType(t *testing.T) {
	testcases := []gameConfigLogicTestcase{
		{
			name:        "Guest unison correct color",
			cardIds:     []int{300830, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     200698,
			songAttr:    enum.AttrCool,
			expected:    false,
		},
		{
			name:        "Lead unison correct color",
			cardIds:     []int{200698, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     300830,
			songAttr:    enum.AttrCool,
			expected:    false,
		},
		{
			name:        "Lead unison wrong color",
			cardIds:     []int{200698, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     300830,
			songAttr:    enum.AttrCute,
			expected:    true,
		},
		{
			name:        "Two unison, one wrong",
			cardIds:     []int{200698, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     100773,
			songAttr:    enum.AttrCool,
			expected:    true,
		},
		{
			name:        "Unison on all color",
			cardIds:     []int{300830, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     200698,
			songAttr:    enum.AttrAll,
			expected:    true,
		},
		{
			name:        "No unison",
			cardIds:     []int{300830, 100344, 200834, 100282, 300124},
			leaderIndex: 0,
			guestID:     100344,
			songAttr:    enum.AttrCool,
			expected:    false,
		},
	}
	logic := unisonInCorrectSongType
	testGameConfigLogic(t, testcases, logic)
}
func TestGameConfig_AllLeadSkillsActive(t *testing.T) {
	testcases := []gameConfigLogicTestcase{
		{
			name:        "Both active",
			cardIds:     []int{300830, 100344, 300236, 200834, 100730},
			leaderIndex: 2,
			guestID:     300412,
			songAttr:    enum.AttrCool,
			expected:    false,
		},
		{
			name:        "Only lead active",
			cardIds:     []int{300830, 100344, 300236, 200834, 100730},
			leaderIndex: 2,
			guestID:     200856,
			songAttr:    enum.AttrCool,
			expected:    true,
		},
		{
			name:        "Only guest active",
			cardIds:     []int{300830, 100344, 300572, 200834, 100730},
			leaderIndex: 2,
			guestID:     200856,
			songAttr:    enum.AttrCool,
			expected:    true,
		},
		{
			name:        "Trico in twocolor team + diff color active guest",
			cardIds:     []int{300830, 100344, 300236, 100578, 100730},
			leaderIndex: 2,
			guestID:     200254,
			songAttr:    enum.AttrCool,
			expected:    false,
		},
		{
			name:        "Trico in twocolor team + same color active guest",
			cardIds:     []int{300830, 100344, 300236, 100578, 100730},
			leaderIndex: 2,
			guestID:     300280,
			songAttr:    enum.AttrCool,
			expected:    true,
		},
	}
	logic := allLeadSkillsActive
	testGameConfigLogic(t, testcases, logic)
}
