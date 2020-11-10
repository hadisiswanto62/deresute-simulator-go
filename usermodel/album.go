package usermodel

import (
	"fmt"
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
	indices              []int
	sortedAll            []*OwnedCard
	sortedCool           []*OwnedCard
	sortedCute           []*OwnedCard
	sortedPassion        []*OwnedCard
}

func (a Album) String() string {
	return fmt.Sprintf("(Album of %d cards)", len(a.Ocards))
}

// NewAlbum creates Album object from owned cards
func NewAlbum(ocards []*OwnedCard) *Album {
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
	album := Album{
		Ocards: ocards, combinationGenerator: gen, currentLeaderIndex: -1,
		sortedCute: sortedCute, sortedCool: sortedCool, sortedPassion: sortedPassion,
		sortedAll: sortedAll, indices: make([]int, 5),
	}
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
	// if currentLeaderIndex = 0, that means generator will return new combination (from album.Next())
	// so get the new combination and store it
	// else, that means generator only change leaderIndex so just retrieve the stored combination
	if a.currentLeaderIndex == 0 {
		a.combinationGenerator.Combination(a.indices)
		a.currentIndices = a.indices
	} else {
		a.indices = a.currentIndices
	}
	ocards := [5]*OwnedCard{}
	for i, index := range a.indices {
		ocards[i] = a.Ocards[index]
	}
	return &Team{Ocards: ocards, LeaderIndex: a.currentLeaderIndex}
}

type baseTeamData struct {
	ids         []int
	leaderIndex int
}

func (a *Album) GetTeamDebug() *Team {
	data := a.getTeamDebug()
	ocards := [5]*OwnedCard{}
	for i, index := range data.ids {
		ocards[i] = a.Ocards[index]
	}
	return &Team{Ocards: ocards, LeaderIndex: data.leaderIndex}
}

// GetTeamDebug returns a new Team from generator
func (a *Album) getTeamDebug() *baseTeamData {
	if a.combinationGenerator == nil {
		panic("combination generator not set")
	}
	// if currentLeaderIndex = 0, that means generator will return new combination (from album.Next())
	// so get the new combination and store it
	// else, that means generator only change leaderIndex so just retrieve the stored combination
	if a.currentLeaderIndex == 0 {
		a.combinationGenerator.Combination(a.indices)
		a.currentIndices = a.indices
	} else {
		a.indices = a.currentIndices
	}
	return &baseTeamData{ids: a.indices, leaderIndex: a.currentLeaderIndex}
	// ocards := [5]*OwnedCard{}
	// for i, index := range a.indices {
	// 	ocards[i] = a.Ocards[index]
	// }
	// return &Team{Ocards: ocards, LeaderIndex: a.currentLeaderIndex}
}

func (a Album) MaxTeamID() int {
	return combin.Binomial(len(a.Ocards), 5)*5 - 1
}

func (a Album) FindSupportsFor(team *Team, attr enum.Attribute) ([10]*OwnedCard, error) {
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

	ids := [5]int{}
	for i, ocard := range team.Ocards {
		ids[i] = ocard.Card.ID
	}
	for _, ocard := range usedSlice {
		isInTeam := false
		for _, id := range ids {
			if id == ocard.Card.ID {
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

func (a Album) FindSupportsFor2(team *Team, attr enum.Attribute) ([10]*OwnedCard, error) {
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
