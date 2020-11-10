package simulator

import (
	"fmt"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"

	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"github.com/hadisiswanto62/deresute-simulator-go/songmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

var sm *songmanager.SongManager

func init() {
	// os.Chdir("../")
	var err error
	cm, err = cardmanager.Default()
	if err != nil {
		panic(err)
	}
	sm, err = songmanager.Default()
	if err != nil {
		panic(err)
	}
}

func TestSimulate_All100Percent(t *testing.T) {
	cardIds := []int{200726, 300830, 300236, 300572, 200314}
	leaderIndex := 2
	guestId := 100298
	songName := "M@GIC"
	result, err := getAlbum(songName)

	song := result.song
	ocards := [5]*usermodel.OwnedCard{}
	for i, ID := range cardIds {
		ocard, err := result.album.GetCardByID(ID)
		if err != nil {
			t.Errorf("cannot get card: %v", err)
		}
		ocards[i] = ocard
	}
	var guest *usermodel.OwnedCard
	for _, ocard := range result.guests {
		if ocard.Card.ID == guestId {
			guest = ocard
		}
	}
	team := &usermodel.Team{ocards, leaderIndex}
	supports, err := result.album.FindSupportsFor(team, song.Attribute)
	if err != nil {
		t.Errorf("cannot find supports: %v", err)
	}
	gc := simulatormodels.NewGameConfig(
		ocards[:], leaderIndex, supports[:],
		guest, song, 0, statcalculator.NormalStatCalculator)

	res := Simulate(gc, 100)
	fmt.Println(res.Report())
	fmt.Println(res.ReportOneline())
}
