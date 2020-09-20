package shortcut

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/csvmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodelmanager"
)

type BaseGameConfig struct {
	CardIDs     [5]int
	GuestID     int
	Song        *models.Song
	SongAttr    enum.Attribute
	LeaderIndex int
}

func (c BaseGameConfig) getSong() *models.Song {
	return c.Song
}

func (c BaseGameConfig) getSongAttr() enum.Attribute {
	return c.SongAttr
}

type BaseOptimizeConfig struct {
	CardsPath     string
	GuestsPath    string
	Song          *models.Song
	SongAttr      enum.Attribute
	SimulateTimes int
}

func (c BaseOptimizeConfig) getSong() *models.Song {
	if c.Song == nil {
		var err error
		c.Song, err = makeSong(c)
		if err != nil {
			panic(err)
		}
	}
	return c.Song
}

func (c BaseOptimizeConfig) getSongAttr() enum.Attribute {
	return c.SongAttr
}

func ToGameConfig(config BaseGameConfig, customCardParams *usermodelmanager.CustomOwnedCardParameters,
	useDefaultCards bool) (*simulator.GameConfig, error) {
	if config.Song == nil {
		var err error
		config.Song, err = makeSong(config)
		if err != nil {
			return nil, fmt.Errorf("Cannot make song: %f", err)
		}
	}
	if customCardParams == nil {
		customCardParams = &usermodelmanager.CustomOwnedCardParameters{}
	}

	cm, err := cardmanager.Default()
	if err != nil {
		panic(err)
	}
	dp := csvmodels.CSVDataParser{}
	albumCards, err := usermodelmanager.ParseOwnedCard(dp, "userdata/cards.csv", nil)
	if err != nil {
		panic(err)
	}

	ocards := [5]*usermodel.OwnedCard{}
	if useDefaultCards {
		cards := [5]*models.Card{}
		for i, id := range config.CardIDs {
			card := cm.Filter().ID(id).First()
			cards[i] = card
		}
		for i, card := range cards {
			request := usermodel.OwnedCardRequest{
				Card:       card,
				SkillLevel: customCardParams.SkillLevel,
				PotDance:   customCardParams.PotDance,
				PotSkill:   customCardParams.PotSkill,
				PotVisual:  customCardParams.PotVisual,
				PotVocal:   customCardParams.PotVocal,
				StarRank:   1,
			}
			ocards[i] = usermodel.NewOwnedCard2(request)
		}
	} else {
		for i, id := range config.CardIDs {
			ocards[i] = usermodelmanager.FindById(albumCards, id)
		}
	}

	guest := usermodel.NewOwnedCard2(usermodel.OwnedCardRequest{
		Card:       cm.Filter().ID(config.GuestID).First(),
		SkillLevel: customCardParams.SkillLevel,
		PotDance:   customCardParams.PotDance,
		PotSkill:   customCardParams.PotSkill,
		PotVisual:  customCardParams.PotVisual,
		PotVocal:   customCardParams.PotVocal,
	})

	album := usermodel.NewAlbum(albumCards)
	team := usermodel.Team{ocards, config.LeaderIndex}
	supports, err := album.FindSupportsFor(&team, config.Song.Attribute)
	if err != nil {
		panic(err)
	}
	return simulator.NewGameConfig(ocards[:], config.LeaderIndex, supports[:], guest, config.Song), nil
}

type makeSongable interface {
	getSongAttr() enum.Attribute
}

func makeSong(m makeSongable) (*models.Song, error) {
	if m.getSongAttr() != "" {
		song := models.NewDefaultSong("", 26, m.getSongAttr(), 127000, 736)
		return &song, nil
	}
	return nil, fmt.Errorf("Invalid song")
}
