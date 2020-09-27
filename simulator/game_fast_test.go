package simulator

import (
	"fmt"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/songmanager"
)

func getGc() *GameConfig {
	//
	songName := "M@GIC"
	cardIds := [5]int{
		200726, 300830, 300236, 200314, 100282,
	}
	leaderIndex := 2
	guestId := 200346
	//

	sm, err := songmanager.Default()
	if err != nil {
		panic(err)
	}
	song := sm.Filter().NameLike(songName).Difficulty(enum.SongDifficultyMaster).First()

	cm, err := cardmanager.Default()
	if err != nil {
		panic(err)
	}
	ocards := [5]*usermodel.OwnedCard{}
	for i, id := range cardIds {
		request := usermodel.OwnedCardRequest{
			Card:       cm.Filter().ID(id).First(),
			SkillLevel: 10,
			// PotSkill:   10,
			StarRank: 1,
		}
		ocards[i] = usermodel.NewOwnedCard2(request)
	}
	guest := usermodel.NewOwnedCard2(usermodel.OwnedCardRequest{
		Card:       cm.Filter().ID(guestId).First(),
		SkillLevel: 10,
		PotSkill:   10,
		StarRank:   1,
	})
	supports := [10]*usermodel.OwnedCard{}
	for i := 0; i < 10; i++ {
		supports[i] = usermodel.NewOwnedCard2(usermodel.OwnedCardRequest{
			Card:     cm.Filter().ID(guestId).First(),
			StarRank: 15,
		})
	}
	return NewGameConfig(ocards[:], leaderIndex, supports[:], guest, song)
}

func TestGameFast(t *testing.T) {
	gc := getGc()
	game := NewGameFast(gc)
	result := game.Play(true)
	fmt.Println(result.Score)
}

func BenchmarkRoll(b *testing.B) {
	gc := getGc()
	game := NewGameFast(gc)
	state := initConfig(game.Config)
	state.alwaysGoodRolls = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rollSkill(state)
	}
}

func BenchmarkGetSkillActive(b *testing.B) {
	gc := getGc()
	game := NewGameFast(gc)
	notesTimestamp := []int{}
	for _, note := range gc.song.Notes {
		notesTimestamp = append(notesTimestamp, note.TimestampMs)
	}
	state := initConfig(game.Config)
	state.alwaysGoodRolls = true
	skills := rollSkill(state)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, timestamp := range notesTimestamp {
			game.getActiveSkillsOn(timestamp, &skills.activeSkillTimestamps)
		}
	}
}

func BenchmarkPlay(b *testing.B) {
	gc := getGc()
	game := NewGameFast(gc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Play(false)
	}
}
