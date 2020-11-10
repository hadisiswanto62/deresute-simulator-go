package simulator

import (
	"fmt"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/songmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodelmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/csvmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
)

var cm *cardmanager.CardManager

func init() {
	var err error
	cm, err = cardmanager.Default()
	if err != nil {
		panic(fmt.Errorf("cannot init cardmanager: %v", err))
	}
}

func TestOptimizer2_FilterConfigs(t *testing.T) {
	cardNames := []string{
		"aiko4", "nana2", "megumi1", "shiki2", "nina1", "yuu1",
		"michiru1", "chieri1",
	}
	guestNames := []string{
		"kako3", "yuuki4", "yoshino4",
		"kaede2", "syuko2", "nana2",
	}
	song := models.NewDefaultSong("", 26, enum.AttrAll, 1000, 1000)

	ocards := []*usermodel.OwnedCard{}
	for _, name := range cardNames {
		card := cm.Filter().SsrNameID(name).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).StarRank(15).Build()
		if err != nil {
			t.Errorf("cannot create card: %v", err)
		}
		ocards = append(ocards, ocard)
	}
	album := usermodel.NewAlbum2(ocards)
	guests := []*usermodel.OwnedCard{}
	for _, name := range guestNames {
		card := cm.Filter().SsrNameID(name).First()
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create card: %v", err)
		}
		guests = append(guests, ocard)
	}

	count := 0
	for range GetFilteredGameConfigs(album, guests, &song) {
		count++
	}
	fmt.Println(count)
}

type msg struct {
	album  *usermodel.Album2
	guests []*usermodel.OwnedCard
	song   *models.Song
}

func getAlbum(songName string) (msg, error) {
	cardsPath := "userdata/cards.csv"
	guestPath := "userdata/guest tricolor.csv"
	sm, err := songmanager.Default()
	if err != nil {
		panic("")
	}
	song := sm.Filter().NameLike(songName).Difficulty(enum.SongDifficultyMaster).First()

	params := usermodelmanager.CustomOwnedCardParameters{
		SkillLevel: 10,
		PotSkill:   10,
	}

	dp := csvmodels.CSVDataParser{}
	baseOcards, err := usermodelmanager.ParseOwnedCard(dp, cardsPath, &params)
	if err != nil {
		return msg{}, fmt.Errorf("cannot create ocards: %v", err)
	}
	guests, err := usermodelmanager.ParseOwnedCard(dp, guestPath, nil)
	if err != nil {
		return msg{}, fmt.Errorf("cannot create guests: %v", err)
	}
	ocards := []*usermodel.OwnedCard{}
	if song.Attribute == enum.AttrAll {
		ocards = baseOcards
	} else {
		for _, ocard := range baseOcards {
			if ocard.Card.Idol.Attribute == song.Attribute {
				ocards = append(ocards, ocard)
			}
		}
	}
	return msg{
		album:  usermodel.NewAlbum2(ocards),
		guests: guests,
		song:   song,
	}, nil
}

func TestOptimizer2_FilterConfigsRealData(t *testing.T) {
	result, err := getAlbum("TOKIMEKI")
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	album := result.album
	guests := result.guests
	song := result.song

	count := 0
	for range GetFilteredGameConfigs(album, guests, song) {
		count++
	}
	fmt.Println(count)
}

func TestOptimizer2_Optimize(t *testing.T) {
	// result, err := getAlbum("O-Ku-Ri-Mo-No")
	// result, err := getAlbum("M@GIC")
	result, err := getAlbum("M@GIC")
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	album := result.album
	guests := result.guests
	song := result.song

	generator := GetFilteredGameConfigs(album, guests, song)
	Optimize(generator, 10, "test_all.txt")
}
