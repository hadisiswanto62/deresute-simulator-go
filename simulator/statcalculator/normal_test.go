package statcalculator

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"

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
func defaultOcard(cards ...*models.Card) []*usermodel.OwnedCard {
	ocards := []*usermodel.OwnedCard{}
	for _, card := range cards {
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			fmt.Errorf("cannot create card: %v", err)
		}
		ocards = append(ocards, ocard)
	}
	return ocards
}
func TestNormalCalculator_Calculate(t *testing.T) {
	cards := cm.Filter().Rarity(enum.RaritySSR).Get()[:10]
	leaderIndex := 0
	guestC := cm.Filter().Rarity(enum.RaritySSR).First()
	supportsC := cm.Filter().Rarity(enum.RaritySSR).Get()[20:30]
	song := models.NewDefaultSong("", 26, enum.AttrAll, 100000, 100)

	ocards := defaultOcard(cards...)
	supports := defaultOcard(supportsC...)
	guest := defaultOcard(guestC)[0]

	calc := normalStatCalculator{}
	calc.SetCards(ocards)
	calc.SetLeaderIndex(leaderIndex)
	calc.SetGuest(guest)
	calc.SetSong(&song)
	calc.SetSupports(supports)
	stats, err := calc.Calculate(0)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 396840, "Wrong appeal!")
	assert.Equal(t, stats.Hp, 504, "Wrong hp!")
	// with bonus Appeal
	stats, err = calc.Calculate(1000)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 397840, "Wrong appeal if use bonus!")
}

func TestNormalCalculator_NoGuest(t *testing.T) {
	cards := cm.Filter().Rarity(enum.RaritySSR).Get()[:10]
	leaderIndex := 0
	supportsC := cm.Filter().Rarity(enum.RaritySSR).Get()[20:30]
	song := models.NewDefaultSong("", 26, enum.AttrAll, 100000, 100)

	ocards := defaultOcard(cards...)
	supports := defaultOcard(supportsC...)

	calc := normalStatCalculator{}
	calc.SetCards(ocards)
	calc.SetLeaderIndex(leaderIndex)
	calc.SetGuest(nil)
	calc.SetSong(&song)
	calc.SetSupports(supports)
	stats, err := calc.Calculate(0)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 316255, "Wrong appeal!")
	assert.Equal(t, stats.Hp, 462, "Wrong hp!")
	// with bonus Appeal
	stats, err = calc.Calculate(1000)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 317255, "Wrong appeal if use bonus!")
}

func TestNormalCalculator_NoSupports(t *testing.T) {
	cards := cm.Filter().Rarity(enum.RaritySSR).Get()[:10]
	leaderIndex := 0
	guestC := cm.Filter().Rarity(enum.RaritySSR).First()
	song := models.NewDefaultSong("", 26, enum.AttrAll, 100000, 100)

	ocards := defaultOcard(cards...)
	guest := defaultOcard(guestC)[0]

	calc := normalStatCalculator{}
	calc.SetCards(ocards)
	calc.SetLeaderIndex(leaderIndex)
	calc.SetGuest(guest)
	calc.SetSong(&song)
	calc.SetSupports(nil)
	stats, err := calc.Calculate(0)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 312520, "Wrong appeal!")
	assert.Equal(t, stats.Hp, 504, "Wrong hp!")
	// with bonus Appeal
	stats, err = calc.Calculate(1000)
	if err != nil {
		t.Errorf("failed to calcaulte: %v", err)
	}
	assert.Equal(t, stats.Appeal, 313520, "Wrong appeal if use bonus!")
}
