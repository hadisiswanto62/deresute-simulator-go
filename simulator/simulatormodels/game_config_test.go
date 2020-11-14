package simulatormodels

import (
	"fmt"
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"
	"github.com/hadisiswanto62/deresute-simulator-go/songmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("../../")
}

func defaultOcard(cards ...*models.Card) []*usermodel.OwnedCard {
	ocards := []*usermodel.OwnedCard{}
	for _, card := range cards {
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			panic(fmt.Errorf("cannot create card: %v", err))
		}
		ocards = append(ocards, ocard)
	}
	return ocards
}

func getGc() *GameConfig {
	return getGcCustomCalc(0, statcalculator.NormalStatCalculator)
}

// do not change card ids, game_fast_test tests the score
func getGcCustomCalc(bonusAppeal int, calcType statcalculator.StatCalculatorType) *GameConfig {
	//
	songName := "M@GIC"
	cardIds := [5]int{
		200726, 300830, 300236, 200314, 100282,
	}
	leaderIndex := 2
	guestId := 200346
	//

	sm, err := songmanager.Default()
	if err != nil {
		panic(err)
	}
	song := sm.Filter().NameLike(songName).Difficulty(enum.SongDifficultyMaster).First()

	cm, err := cardmanager.Default()
	if err != nil {
		panic(err)
	}
	ocards := [5]*usermodel.OwnedCard{}
	for i, id := range cardIds {
		request := usermodel.OwnedCardRequest{
			Card:       cm.Filter().ID(id).First(),
			SkillLevel: 10,
			// PotSkill:   10,
			StarRank: 1,
		}
		ocards[i] = usermodel.NewOwnedCard2(request)
	}
	guest := usermodel.NewOwnedCard2(usermodel.OwnedCardRequest{
		Card:       cm.Filter().ID(guestId).First(),
		SkillLevel: 10,
		PotSkill:   10,
		StarRank:   1,
	})
	supports := [10]*usermodel.OwnedCard{}
	for i := 0; i < 10; i++ {
		supports[i] = usermodel.NewOwnedCard2(usermodel.OwnedCardRequest{
			Card:     cm.Filter().ID(guestId).First(),
			StarRank: 15,
		})
	}
	return NewGameConfig(ocards[:], leaderIndex, supports[:], guest, song, bonusAppeal, calcType)
}

func TestGameConfig_NormalCalc_Appeal(t *testing.T) {
	gc := getGcCustomCalc(0, statcalculator.NormalStatCalculator)
	assert.Equal(t, gc.Appeal(), 324682, "Wrong appeals")
	gc = getGcCustomCalc(1000, statcalculator.NormalStatCalculator)
	assert.Equal(t, gc.Appeal(), 325682, "Wrong appeals")
}

func TestGameConfig_NormalCalc_WithoutGuestAndSupports(t *testing.T) {
	cm, _ := cardmanager.Default()
	cards := cm.Filter().Rarity(enum.RaritySSR).Get()[:10]

	ocards := defaultOcard(cards...)
	leaderIndex := 2
	song := models.NewDefaultSong("", 26, enum.AttrAll, 100000, 100)

	gc := NewGameConfig(ocards, leaderIndex, nil, nil, &song, 0, statcalculator.NormalStatCalculator)
	assert.Equal(t, gc.Appeal(), 219473, "Doesn't work! (maybe cm.Filter() returns different card, check!)")

	supports := []*usermodel.OwnedCard{}
	var guest *usermodel.OwnedCard
	gc = NewGameConfig(ocards, leaderIndex, supports, guest, &song, 0, statcalculator.NormalStatCalculator)
	assert.Equal(t, gc.Appeal(), 219473, "Doesn't work! (maybe cm.Filter() returns different card, check!)")
}
