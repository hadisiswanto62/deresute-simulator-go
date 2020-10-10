package statcalculator

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type normalStatCalculator struct {
	cards       []*usermodel.OwnedCard
	leaderIndex int
	guest       *usermodel.OwnedCard
	supports    []*usermodel.OwnedCard
	song        *models.Song
}

func (sc *normalStatCalculator) SetCards(cards []*usermodel.OwnedCard) {
	sc.cards = cards
}

func (sc *normalStatCalculator) SetLeaderIndex(index int) {
	sc.leaderIndex = index
}

func (sc *normalStatCalculator) SetGuest(guest *usermodel.OwnedCard) {
	sc.guest = guest
}

func (sc *normalStatCalculator) SetSupports(supports []*usermodel.OwnedCard) {
	sc.supports = supports
}

func (sc *normalStatCalculator) SetSong(song *models.Song) {
	sc.song = song
}

func (sc normalStatCalculator) Calculate(bonusAppeal int) (*GameConfigStats, error) {
	if sc.cards == nil {
		return nil, fmt.Errorf("cards not set")
	}

	leadSkillsActivableCards := []*usermodel.OwnedCard{
		sc.cards[sc.leaderIndex], sc.guest,
	}
	ocards := append(sc.cards, sc.guest)

	data := &GameConfigStats{}
	populateTeamAttributesSkills(data, ocards)
	populateResonant(data, leadSkillsActivableCards)
	populateAppealHp(data, ocards, leadSkillsActivableCards, sc.song)

	data.Appeal += calcSupportAppeals(data, sc.supports)
	data.Appeal += bonusAppeal
	return data, nil
}
