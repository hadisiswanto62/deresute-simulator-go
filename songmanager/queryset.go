package songmanager

import (
	"strings"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

type QuerySet struct {
	songs []*models.Song
}

func (q *QuerySet) Attribute(attr enum.Attribute) *QuerySet {
	result := []*models.Song{}
	for i := range q.songs {
		if q.songs[i].Attribute == attr {
			result = append(result, q.songs[i])
		}
	}
	q.songs = result
	return q
}

func (q *QuerySet) NameLike(name string) *QuerySet {
	result := []*models.Song{}
	for i := range q.songs {
		if strings.Contains(
			strings.ToLower(q.songs[i].Name),
			strings.ToLower(name),
		) {
			result = append(result, q.songs[i])
		}
	}
	q.songs = result
	return q
}

func (q *QuerySet) Difficulty(diff enum.SongDifficulty) *QuerySet {
	result := []*models.Song{}
	for i := range q.songs {
		if q.songs[i].Difficulty == diff {
			result = append(result, q.songs[i])
		}
	}
	q.songs = result
	return q
}

// Get gets all cards that matches current filter
func (q *QuerySet) Get() []*models.Song {
	return q.songs
}

// First gets the first card that matches current filter
func (q *QuerySet) First() *models.Song {
	if len(q.songs) == 0 {
		return nil
	}
	return q.songs[0]
}
