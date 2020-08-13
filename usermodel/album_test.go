package usermodel

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/stretchr/testify/assert"
)

func sampleAlbum(n int) *Album {
	var ocards []*OwnedCard
	for i := 0; i < n; i++ {
		randAppeal := helper.RandInt(1000, 10000)
		attrIndex := helper.RandInt(0, 3)
		randAttr := enum.AttrForIdol[attrIndex]
		starRank := 3
		ocards = append(ocards, &OwnedCard{
			Card: &models.Card{
				ID: i,
				Idol: &models.Idol{
					Attribute: randAttr,
				},
			},
			Appeal:   randAppeal,
			StarRank: starRank,
		})
	}
	return NewAlbum(ocards)
}

func TestFindSupportsMultipleStarRank(t *testing.T) {
	n := 300
	album := sampleAlbum(n)
	album.Next()
	team := album.GetTeam()
	supports, err := album.FindSupportsFor(team, enum.AttrPassion)
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	// sampleAlbum() generates cards with starRank=3 so should have 4 unique cards in supports
	uniques := make(map[int]bool)
	for _, ocard := range supports {
		uniques[ocard.Card.ID] = true
	}
	assert.Equal(t, 4, len(uniques), "Too many (or too less) unique card!")
	assert.Equal(t, 10, len(supports), "Too many (or too less) supports!")
}

func TestFindSupportsNotPlaying(t *testing.T) {
	assertion := assert.New(t)
	album := sampleAlbum(5)
	ocards := album.Ocards
	ocards = append(ocards, &OwnedCard{
		Card: &models.Card{
			ID: 100,
			Idol: &models.Idol{
				Attribute: enum.AttrAll,
			},
		},
		Appeal:   1,
		StarRank: 20,
	})
	album = NewAlbum(ocards)
	album.Next()
	team := album.GetTeam()
	supports, _ := album.FindSupportsFor(team, enum.AttrAll)
	for _, support := range supports {
		assertion.Equal(100, support.Card.ID, "Currently playing card is included in supports!")
	}
}

func TestFindSupports(t *testing.T) {
	assertion := assert.New(t)
	threshold := 9000 * 10
	n := 300
	album := sampleAlbum(n)
	album.Next()
	team := album.GetTeam()
	supports, err := album.FindSupportsFor(team, enum.AttrAll)
	if err != nil {
		t.Errorf("test error: %v", err)
	}
	sum := 0
	for _, ocard := range supports {
		sum += ocard.Appeal
	}
	assertion.Greater(sum, threshold, "Appeal sum is suspiciously low! FindSupportsFor not sorted?")

	supports, err = album.FindSupportsFor(team, enum.AttrCute)
	if err != nil {
		t.Errorf("test error: %v", err)
	}
	sum = 0
	for _, ocard := range supports {
		sum += ocard.Appeal
		assertion.Equal(ocard.Card.Idol.Attribute, enum.AttrCute, "FindSupportsFor returns non-song-attribute card!")
	}
	assertion.Greater(sum, threshold, "Appeal sum is suspiciously low! FindSupportsFor not sorted?")

	supports, err = album.FindSupportsFor(team, enum.AttrCool)
	if err != nil {
		t.Errorf("test error: %v", err)
	}
	sum = 0
	for _, ocard := range supports {
		sum += ocard.Appeal
		assertion.Equal(ocard.Card.Idol.Attribute, enum.AttrCool, "FindSupportsFor returns non-song-attribute card!")
	}
	assertion.Greater(sum, threshold, "Appeal sum is suspiciously low! FindSupportsFor not sorted?")

	supports, err = album.FindSupportsFor(team, enum.AttrPassion)
	if err != nil {
		t.Errorf("test error: %v", err)
	}
	sum = 0
	for _, ocard := range supports {
		sum += ocard.Appeal
		assertion.Equal(ocard.Card.Idol.Attribute, enum.AttrPassion, "FindSupportsFor returns non-song-attribute card!")
	}
	assertion.Greater(sum, threshold, "Appeal sum is suspiciously low! FindSupportsFor not sorted?")
}

func TestMaxTeamID(t *testing.T) {
	n := 20
	album := sampleAlbum(n)
	assert.Equal(t, album.MaxTeamID(), 77519, "Wrong max team id!")
}

func TestMakeTeam(t *testing.T) {
	assertion := assert.New(t)
	n := 7
	album := sampleAlbum(n)
	i := 0
	for album.Next() {
		team := album.GetTeam()
		assertion.NotNil(team, "Team not generated!")
		i++
	}
	assertion.Equal(album.MaxTeamID()+1, i, "Not enough team generated?")
}
