package usermodel

import (
	"sort"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"gonum.org/v1/gonum/stat/combin"
)

// Album is collections of OwnedCard that the user have.
// Use `NewAlbum()` if need to use combination functionalities.
type Album struct {
	Ocards               []*OwnedCard
	combinationGenerator *combin.CombinationGenerator
	currentIndices       []int
	currentLeaderIndex   int
}

// NewAlbum creates Album object from owned cards
func NewAlbum(ocards []*OwnedCard) *Album {
	gen := combin.NewCombinationGenerator(len(ocards), 5)
	album := Album{Ocards: ocards, combinationGenerator: gen, currentLeaderIndex: -1}
	return &album
}

// Next moves the team generator to next team
func (a *Album) Next() bool {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	if (a.currentLeaderIndex == 4) || (a.currentLeaderIndex == -1) {
		a.currentLeaderIndex = 0
		return a.combinationGenerator.Next()
	}
	a.currentLeaderIndex++
	return true
}

// GetTeam returns a new Team from generator
func (a *Album) GetTeam() *Team {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	indices := make([]int, 5)
	// if currentLeaderIndex = 0, that means generator will return new combination (from album.Next())
	// so get the new combination and store it
	// else, that means generator only change leaderIndex so just retrieve the stored combination
	if a.currentLeaderIndex == 0 {
		a.combinationGenerator.Combination(indices)
		a.currentIndices = indices
	} else {
		indices = a.currentIndices
	}
	ocards := [5]*OwnedCard{}
	for i, index := range indices {
		ocards[i] = a.Ocards[index]
	}
	return &Team{Ocards: ocards, LeaderIndex: a.currentLeaderIndex}
}

func (a Album) MaxTeamID() int {
	return combin.Binomial(len(a.Ocards), 5)*5 - 1
}

func (a Album) FindSupportsFor(team *Team, attr enum.Attribute) [10]*OwnedCard {
	ocards := a.Ocards
	sort.SliceStable(ocards, func(i, j int) bool {
		return ocards[i].Appeal > ocards[j].Appeal
	})
	var ret [10]*OwnedCard
	nextIndex := 0
	for _, ocard := range ocards {
		if (attr != enum.AttrAll) && (attr != ocard.Card.Idol.Attribute) {
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
	return ret
}
