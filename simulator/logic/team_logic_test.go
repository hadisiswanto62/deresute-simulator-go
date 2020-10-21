package logic

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func TestTeamLogic_LeadSkillIsImplemented(t *testing.T) {
	testcases := []struct {
		leadSkill enum.LeadSkill
		expected  bool
	}{
		{leadSkill: enum.LeadSkillCoolAbility, expected: true},
		{leadSkill: enum.LeadSkillIrrelevant, expected: false},
	}
	logic := leadSkillIsImplemented

	otherCards := cm.Filter().Rarity(enum.RaritySSR).Get()[:4]
	otherOcards := []*usermodel.OwnedCard{}
	for _, card := range otherCards {
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		otherOcards = append(otherOcards, ocard)
	}

	for _, tc := range testcases {
		card := cm.Filter().LeadSkill(tc.leadSkill).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		ocards := [5]*usermodel.OwnedCard{ocard}
		for i, ocard := range otherOcards {
			ocards[i+1] = ocard
		}
		team := &usermodel.Team{ocards, 0}

		actual := logic.isSatisfied(team, nil)
		assert.Equalf(t, tc.expected, actual, "Wrong result for lead skill %s", tc.leadSkill)
	}
}

type teamLogicTestcase struct {
	name        string
	cardIDs     []int
	leaderIndex int
	expected    bool
}

func testTeamLogic(t *testing.T, testcases []teamLogicTestcase,
	logic *teamLogic) bool {
	for _, tc := range testcases {
		ocardsSlice := makeOcards(tc.cardIDs)
		var ocards [5]*usermodel.OwnedCard
		copy(ocards[:5], ocardsSlice)
		team := usermodel.Team{Ocards: ocards, LeaderIndex: tc.leaderIndex}

		actual := logic.isSatisfied(&team, nil)
		result := assert.Equalf(t, tc.expected, actual, "Wrong for %s", tc.name)
		if !result {
			return result
		}
	}
	return true
}

func TestTeamLogic_TwoCardSameLeadSkillUseLowerID(t *testing.T) {
	testcases := []teamLogicTestcase{
		{
			name:        "use lower id",
			cardIDs:     []int{300845, 300811, 300815, 300817, 300829},
			leaderIndex: 1,
			expected:    true,
		},
		{
			name:        "use higher id",
			cardIDs:     []int{300845, 300811, 300815, 300817, 300829},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "irrelevant",
			cardIDs:     []int{300845, 300811, 300815, 300817, 300829},
			leaderIndex: 2,
			expected:    true,
		},
	}
	logic := twoCardSameLeadSkillThenUseLowerID
	testTeamLogic(t, testcases, logic)
}

func TestTeamLogic_AttrStatUpLeadSkillOnUnicolorTeamOnly(t *testing.T) {
	testcases := []teamLogicTestcase{
		{
			name:        "correct",
			cardIDs:     []int{300349, 300361, 300367, 300375, 300377},
			leaderIndex: 0,
			expected:    true,
		},
		{
			name:        "incorrect",
			cardIDs:     []int{300349, 200361, 300367, 300375, 300377},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "irrelevant",
			cardIDs:     []int{300349, 200361, 300367, 300375, 300377},
			leaderIndex: 4,
			expected:    true,
		},
	}
	logic := attrStatUpLeadSkillOnUnicolorTeamOnly
	testTeamLogic(t, testcases, logic)
}

func TestTeamLogic_TricolorMin2Color(t *testing.T) {
	testcases := []teamLogicTestcase{
		{
			name:        "tricolor 3 color",
			cardIDs:     []int{300236, 300361, 300367, 100252, 200354},
			leaderIndex: 0,
			expected:    true,
		},
		{
			name:        "tricolor 2 color",
			cardIDs:     []int{300236, 300361, 300367, 100252, 100842},
			leaderIndex: 0,
			expected:    true,
		},
		{
			name:        "tricolor 1 color",
			cardIDs:     []int{300236, 300361, 300367, 300368, 300362},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "not tricolor",
			cardIDs:     []int{200726, 300361, 300367, 300368, 300362},
			leaderIndex: 0,
			expected:    true,
		},
	}
	logic := tricolorMin2Color
	testTeamLogic(t, testcases, logic)
}

func TestTeamLogic_TricolorMin3Color(t *testing.T) {
	testcases := []teamLogicTestcase{
		{
			name:        "tricolor 3 color",
			cardIDs:     []int{300236, 300361, 300367, 100252, 200354},
			leaderIndex: 0,
			expected:    true,
		},
		{
			name:        "tricolor 2 color",
			cardIDs:     []int{300236, 300361, 300367, 100252, 100842},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "tricolor 1 color",
			cardIDs:     []int{300236, 300361, 300367, 300368, 300362},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "not tricolor",
			cardIDs:     []int{200726, 300361, 300367, 300368, 300362},
			leaderIndex: 0,
			expected:    true,
		},
	}
	logic := tricolorMin3Color
	testTeamLogic(t, testcases, logic)
}

func TestTeamLogic_ResoOnThreeSkills(t *testing.T) {
	testcases := []teamLogicTestcase{
		{
			name:        "reso 3 skills",
			cardIDs:     []int{300712, 300096, 300152, 300080, 300108},
			leaderIndex: 0,
			expected:    false,
		},
		{
			name:        "reso 4 skills",
			cardIDs:     []int{300712, 300096, 300152, 300080, 300120},
			leaderIndex: 0,
			expected:    true,
		},
		{
			name:        "not reso 3 skills",
			cardIDs:     []int{300712, 300096, 300152, 300080, 300108},
			leaderIndex: 1,
			expected:    true,
		},
	}
	logic := cardsResoOn3UniqueSkills
	testTeamLogic(t, testcases, logic)
}
