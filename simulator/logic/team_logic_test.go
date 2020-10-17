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
