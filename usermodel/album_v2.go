package usermodel

import (
	"fmt"
	"sort"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"gonum.org/v1/gonum/stat/combin"
)

// Album2 is collections of OwnedCard that the user have.
// Use `NewAlbum2()` if need to use combination functionalities.
type Album2 struct {
	Ocards               []*OwnedCard
	combinationGenerator *combin.CombinationGenerator

	indices []int

	sortedAll     []*OwnedCard
	sortedCool    []*OwnedCard
	sortedCute    []*OwnedCard
	sortedPassion []*OwnedCard
}

func (a Album2) String() string {
	return fmt.Sprintf("(Album2 of %d cards)", len(a.Ocards))
}

// NewAlbum2 creates Album2 object from owned cards
func NewAlbum2(ocards []*OwnedCard) *Album2 {
	sort.SliceStable(ocards, func(i, j int) bool {
		return ocards[i].Appeal > ocards[j].Appeal
	})
	gen := combin.NewCombinationGenerator(len(ocards), 5)
	sortedCute := []*OwnedCard{}
	sortedCool := []*OwnedCard{}
	sortedPassion := []*OwnedCard{}
	sortedAll := []*OwnedCard{}
	for _, ocard := range ocards {
		sortedAll = append(sortedAll, ocard)
		switch ocard.Card.Idol.Attribute {
		case enum.AttrCute:
			sortedCute = append(sortedCute, ocard)
		case enum.AttrCool:
			sortedCool = append(sortedCool, ocard)
		case enum.AttrPassion:
			sortedPassion = append(sortedPassion, ocard)
		}
	}
	Album2 := Album2{
		Ocards: ocards, combinationGenerator: gen,
		sortedCute: sortedCute, sortedCool: sortedCool, sortedPassion: sortedPassion,
		sortedAll: sortedAll, indices: make([]int, 5),
	}
	return &Album2
}

// Next moves the team generator to next team
func (a *Album2) Next() bool {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	return a.combinationGenerator.Next()
}

func (a *Album2) GetCards() [5]*OwnedCard {
	var ocards [5]*OwnedCard
	for i, index := range a.getIndexes() {
		ocards[i] = a.Ocards[index]
	}
	return ocards
}

// GetTeam returns a new Team from generator
func (a *Album2) getIndexes() []int {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	// if currentLeaderIndex = 0, that means generator will return new combination (from Album2.Next())
	// so get the new combination and store it
	// else, that means generator only change leaderIndex so just retrieve the stored combination
	a.combinationGenerator.Combination(a.indices)
	return a.indices
}
func (a Album2) MaxTeamID() int {
	return combin.Binomial(len(a.Ocards), 5) - 1
}

func (a Album2) FindSupportsFor(team *Team, attr enum.Attribute) ([10]*OwnedCard, error) {
	var ret [10]*OwnedCard
	nextIndex := 0
	var usedSlice []*OwnedCard
	switch attr {
	case enum.AttrAll:
		usedSlice = a.sortedAll
	case enum.AttrCute:
		usedSlice = a.sortedCute
	case enum.AttrCool:
		usedSlice = a.sortedCool
	case enum.AttrPassion:
		usedSlice = a.sortedPassion
	}
	for _, ocard := range usedSlice {
		isInTeam := false
		for _, teamOcard := range team.Ocards {
			if ocard == teamOcard {
				isInTeam = true
				break
			}
		}
		if isInTeam {
			continue
		}
		for i := 0; i < ocard.StarRank; i++ {
			ret[nextIndex] = ocard
			nextIndex++
			if nextIndex == 10 {
				break
			}
		}
		if nextIndex == 10 {
			break
		}
	}
	if nextIndex == 10 {
		return ret, nil
	} else {
		err := fmt.Errorf("can only find %d supports", nextIndex)
		return [10]*OwnedCard{}, err
	}
}
