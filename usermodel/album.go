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
}

func NewAlbum(ocards []*OwnedCard) *Album {
	gen := combin.NewCombinationGenerator(len(ocards), 5)
	album := Album{Ocards: ocards, combinationGenerator: gen}
	return &album
}

func (a *Album) Next() bool {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	return a.combinationGenerator.Next()
}

func (a *Album) GetTeam() *Team {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	indices := make([]int, 5)
	a.combinationGenerator.Combination(indices)
	ocards := [5]*OwnedCard{}
	for i, index := range indices {
		ocards[i] = a.Ocards[index]
	}
	return &Team{Ocards: ocards, LeaderIndex: 2}
}

func (a Album) MaxTeamID() int {
	return combin.Binomial(len(a.Ocards), 5) - 1
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
