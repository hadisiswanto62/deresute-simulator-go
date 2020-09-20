package songmanager

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/jsonmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

type SongManager struct {
	Songs    []*models.Song
	querySet *QuerySet
}

type dataParser interface {
	ParseSong(string) ([]*models.Song, error)
}

func (sm *SongManager) Filter() *QuerySet {
	sm.querySet.songs = sm.Songs
	return sm.querySet
}

func Default() (*SongManager, error) {
	var dp dataParser
	dp = jsonmodels.JSONDataParser{}
	if instance == nil {
		songs, err := dp.ParseSong("")
		if err != nil {
			return nil, fmt.Errorf("cannot parse cards: %v", err)
		}
		instance = &SongManager{songs, &QuerySet{songs}}
	}
	return instance, nil
}

var instance *SongManager
