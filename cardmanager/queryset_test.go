package cardmanager

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/stretchr/testify/assert"
)

var cm *CardManager

func init() {
	cm, _ = Default()
}

func TestQuerySetSameInstance(t *testing.T) {
	card := cm.Filter().Rarity(enum.RaritySSR).First()
	card2 := cm.Filter().Rarity(enum.RaritySSR).First()
	assert.Equal(t, &card, &card2, "QuerySet creates new card object!")
}

func TestAttribute(t *testing.T) {
	for _, attr := range enum.AttrForIdol {
		card := cm.Filter().Attribute(attr).First()
		if card.Idol.Attribute != attr {
			t.Errorf("Incorrect card attribute! %v != %v", card.Idol.Attribute, attr)
		}
	}
}

func TestRarity(t *testing.T) {
	assert := assert.New(t)
	for _, rarity := range enum.AllRarities {
		card := cm.Filter().Rarity(rarity).First()
		assert.Equal(card.Rarity.Rarity, rarity, "Incorrect rarity!")
	}
}

func BenchmarkNameLike(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cm.Filter().NameLike("kaede").First()
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cm.Filter().Name("Takagaki Kaede").First()
	}
}

func TestLeadSkill(t *testing.T) {
	assert := assert.New(t)
	testcases := [1]enum.LeadSkill{
		enum.LeadSkillCuteMakeup,
	}
	for _, lskill := range testcases {
		card := cm.Filter().LeadSkill(lskill).First()
		assert.Equal(card.LeadSkill.Name, lskill, "Incorrect lead skill!")
	}
}

func TestSkillType(t *testing.T) {
	assert := assert.New(t)
	testcases := [1]enum.SkillType{
		enum.SkillTypeScoreBonus,
	}
	for _, skill := range testcases {
		card := cm.Filter().SkillType(skill).First()
		assert.Equal(card.Skill.SkillType.Name, skill, "Incorrect skill!")
	}
}

func TestEvolved(t *testing.T) {
	assert := assert.New(t)
	for _, evolveStatus := range [2]bool{true, false} {
		card := cm.Filter().IsEvolved(evolveStatus).First()
		assert.Equal(card.Rarity.IsEvolved, evolveStatus, "Incorrect evolve status!")
	}
}

func TestID(t *testing.T) {
	assert := assert.New(t)
	testcases := [5]int{
		100001,
		200609,
		300706,
		200195,
		100064,
	}
	for _, id := range testcases {
		card := cm.Filter().ID(id).First()
		assert.Equal(card.ID, id, "Incorrect ID")
	}
}
