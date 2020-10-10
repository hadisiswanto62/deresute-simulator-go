package usermodel_test

import (
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"
)

var cm *cardmanager.CardManager
var defaultCard *models.Card

func init() {
	os.Chdir("../")
	cm, _ = cardmanager.Default()
	defaultCard = cm.Filter().Rarity(enum.RaritySSR).First()
}

func TestOwnedCard_Create(t *testing.T) {
	builder := usermodel.NewOwnedCardBuilder()
	_, err := builder.Card(defaultCard).Build()
	assert.Equal(t, err, nil, "Build card failed")
}

func TestOwnedCard_Defaults(t *testing.T) {
	assert := assert.New(t)
	builder := usermodel.NewOwnedCardBuilder()
	ocard, err := builder.Card(defaultCard).Build()
	assert.Equal(err, nil, "Build card failed")
	assert.Equal(ocard.Level(), ocard.Card.Rarity.MaxLevel, "Wrong default level")
	assert.Equal(ocard.SkillLevel(), 10, "Wrong default level")
	assert.Equal(ocard.StarRank, 1, "Wrong default star rank")
	assert.Equal(ocard.PotVisual(), 0, "Wrong default pot vis")
	assert.Equal(ocard.PotDance(), 0, "Wrong default pot da")
	assert.Equal(ocard.PotVocal(), 0, "Wrong default pot vo")
	assert.Equal(ocard.PotHp(), 0, "Wrong default pot hp")
	assert.Equal(ocard.PotSkill(), 0, "Wrong default pot skil")
}

func TestOwnedCard_Values(t *testing.T) {
	assert := assert.New(t)
	builder := usermodel.NewOwnedCardBuilder()
	ocard, err := builder.Card(defaultCard).Level(10).SkillLevel(5).StarRank(15).
		PotVisual(10).PotDance(10).PotVocal(10).PotHp(10).PotSkill(10).Build()
	assert.Equal(err, nil, "Build card failed")
	assert.Equal(ocard.Card, defaultCard, "Wrong card")
	assert.Equal(ocard.Level(), 10, "Wrong default level")
	assert.Equal(ocard.SkillLevel(), 5, "Wrong default level")
	assert.Equal(ocard.StarRank, 15, "Wrong default star rank")
	assert.Equal(ocard.PotVisual(), 10, "Wrong default pot vis")
	assert.Equal(ocard.PotDance(), 10, "Wrong default pot da")
	assert.Equal(ocard.PotVocal(), 10, "Wrong default pot vo")
	assert.Equal(ocard.PotHp(), 10, "Wrong default pot hp")
	assert.Equal(ocard.PotSkill(), 10, "Wrong default pot skil")
}
