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

func TestGameConfig_UnisonInCorrectSongType(t *testing.T) {
	testcases := []struct {
		name        string
		cardIds     []int
		leaderIndex int
		guestID     int
		songAttr    enum.Attribute
		expected    bool
	}{
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

		actual := logic.isViolated(team, activable, &song)
		assert.Equalf(t, tc.expected, actual, "Wrong result for tc %s", tc.name)

	}
}
