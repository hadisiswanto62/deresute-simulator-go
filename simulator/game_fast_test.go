package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/statcalculator"
	"github.com/hadisiswanto62/deresute-simulator-go/songmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

func getGc() *simulatormodels.GameConfig {
	return getGcCustomCalc(0, statcalculator.NormalStatCalculator)
}

// do not change card ids, game_fast_test tests the score
func getGcCustomCalc(bonusAppeal int, calcType statcalculator.StatCalculatorType) *simulatormodels.GameConfig {
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
	return simulatormodels.NewGameConfig(ocards[:], leaderIndex, supports[:], guest, song, bonusAppeal, calcType)
}

func TestGameFast(t *testing.T) {
	gc := getGc()
	game := NewGameFast(gc)
	result := game.Play(true)
	if want, have := 1369869, result.Score; want != have {
		t.Errorf("Score should be %d. (it is %d)", want, have)
	}
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
	for _, note := range gc.GetSong().Notes {
		notesTimestamp = append(notesTimestamp, note.TimestampMs)
	}
	state := initConfig(game.Config)
	state.alwaysGoodRolls = true
	skills := rollSkill(state)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, timestamp := range notesTimestamp {
			game.getActiveSkillsOn(timestamp, &skills.activeSkillTimestamps, enum.NoteTypeTap)
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

type miniCardData struct {
	name string
	skLv int
	stRa int
	vo   int
	vi   int
	da   int
	hp   int
	poSk int
}

func (data miniCardData) toOwnedCard(cm *cardmanager.CardManager) *usermodel.OwnedCard {
	card := cm.Filter().SsrNameID(data.name).First()
	builder := usermodel.NewOwnedCardBuilder()
	ocard, err := builder.Card(card).SkillLevel(data.skLv).StarRank(data.stRa).PotVocal(data.vo).
		PotVisual(data.vi).PotDance(data.da).PotHp(data.hp).PotSkill(data.poSk).Build()
	if err != nil {
		return nil
	}
	return ocard
}

func TestGameFast_CorrectScore(t *testing.T) {
	testcases := []struct {
		guestData       miniCardData
		cardsData       []miniCardData
		leadIndex       int
		supportsData    []miniCardData
		supportAppeals  int
		statCalc        statcalculator.StatCalculatorType
		songName        string
		diff            enum.SongDifficulty
		expectedAppeals int
		expectedScore   int
		windowAbuse     bool
		skillAlwaysOn   bool
	}{
		{
			guestData: miniCardData{name: "sachiko2"},
			cardsData: []miniCardData{
				miniCardData{name: "karen4", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "karen2", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "kanade2", skLv: 10, poSk: 10},
				miniCardData{name: "aiko4", skLv: 10, poSk: 10, da: 3},
				miniCardData{name: "sato3", skLv: 10, poSk: 10},
			},
			supportsData:    nil,
			leadIndex:       2,
			supportAppeals:  102727,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "M@GIC",
			diff:            enum.SongDifficultyMaster,
			expectedAppeals: 283375,
			expectedScore:   1189581,
		},
		{
			guestData: miniCardData{name: "hotaru1", vo: 10, da: 2, vi: 10},
			cardsData: []miniCardData{
				miniCardData{name: "karen4", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "karen2", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "kanade2", skLv: 10, poSk: 10},
				miniCardData{name: "aiko4", skLv: 10, poSk: 10, da: 3},
				miniCardData{name: "sato3", skLv: 10, poSk: 10},
			},
			supportsData:    nil,
			leadIndex:       2,
			supportAppeals:  102727,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "M@GIC",
			diff:            enum.SongDifficultyRegular,
			expectedAppeals: 244271,
			expectedScore:   673762,
		},
		{
			guestData: miniCardData{name: "hotaru1", vo: 10, da: 2, vi: 10},
			cardsData: []miniCardData{
				miniCardData{name: "karen4", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "karen2", skLv: 10, poSk: 10, da: 9},
				miniCardData{name: "kanade2", skLv: 10, poSk: 10},
				miniCardData{name: "aiko4", skLv: 10, poSk: 10, da: 3},
				miniCardData{name: "sato3", skLv: 10, poSk: 10},
			},
			supportsData:    nil,
			leadIndex:       2,
			supportAppeals:  102727,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "M@GIC",
			diff:            enum.SongDifficultyDebut,
			expectedAppeals: 244271,
			expectedScore:   557660,
		},
		{
			guestData: miniCardData{name: "kaede2", vo: 10, da: 10, vi: 10, hp: 5},
			cardsData: []miniCardData{
				miniCardData{name: "mayu5", skLv: 10, poSk: 10, da: 5, vo: 10, vi: 10},
				miniCardData{name: "mio4", skLv: 10, poSk: 10, da: 10, vo: 5, vi: 10},
				miniCardData{name: "nina4", skLv: 10, poSk: 10, da: 6, vi: 10},
				miniCardData{name: "yoshino3", skLv: 10, poSk: 10, da: 10, vi: 10},
				miniCardData{name: "yoshino3u", skLv: 10, poSk: 10, da: 10, vi: 10},
			},
			supportsData:    nil,
			leadIndex:       2,
			supportAppeals:  113290,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "M@GIC",
			diff:            enum.SongDifficultyMaster,
			expectedAppeals: 263468,
			expectedScore:   1934552,
			windowAbuse:     true,
		},
		{
			guestData: miniCardData{name: "uzuki2", vo: 10, da: 10, vi: 10},
			cardsData: []miniCardData{
				miniCardData{name: "anastasia4", skLv: 10, poSk: 10, da: 4, vo: 0, vi: 10},
				miniCardData{name: "sakuma5", skLv: 10, poSk: 10, da: 0, vo: 1, vi: 10},
				miniCardData{name: "nina4", skLv: 10, poSk: 10, da: 3, vo: 0, vi: 10},
				miniCardData{name: "yoshino3", skLv: 10, poSk: 10, da: 10, vi: 10},
				miniCardData{name: "yoshino3u", skLv: 10, poSk: 10, da: 10, vi: 10},
			},
			supportsData:    nil,
			leadIndex:       2,
			supportAppeals:  113290,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "M@GIC",
			diff:            enum.SongDifficultyMaster,
			expectedAppeals: 262916,
			expectedScore:   1909443, // actual = 1924782
		},
		{
			// Confirmed with reddit score sim (!)
			guestData: miniCardData{name: "kaede2", vo: 10, da: 10, vi: 10}, // trico make
			cardsData: []miniCardData{
				miniCardData{name: "sae4", skLv: 10, vo: 2, vi: 10, da: 0, hp: 0, poSk: 0},     // visual motif/reso [motif% idk]
				miniCardData{name: "chieri4", skLv: 10, vo: 0, vi: 10, da: 9, hp: 0, poSk: 0},  // trico synergy [16%/15%]
				miniCardData{name: "yoshino3", skLv: 10, vo: 8, vi: 10, da: 0, hp: 0, poSk: 0}, // trico synergy [16%/15%] [11s, 7.5s]
				miniCardData{name: "rika4", skLv: 10, vo: 8, vi: 10, da: 0, hp: 0, poSk: 0},    // trico symphony
				miniCardData{name: "mio4", skLv: 10, vo: 0, vi: 5, da: 0, hp: 0, poSk: 0},      // coordinate [10%/15%]
			},
			supportsData:    nil,
			leadIndex:       0,
			supportAppeals:  113290 + 2735,
			statCalc:        statcalculator.NormalStatCalculator,
			songName:        "印象",
			diff:            enum.SongDifficultyMaster,
			expectedAppeals: 270000,
			expectedScore:   1582088,
			skillAlwaysOn:   true,
		},
	}
	cm, err := cardmanager.Default()
	if err != nil {
		t.Errorf("cannot initialize cardmanager: %v", err)
	}
	sm, err := songmanager.Default()
	if err != nil {
		panic(err)
	}

	/**
	difference in score is quite big, possibly because some delay in tapping.
	For example: if note is on 90011, and a skill ends at 90000, one can simply tap the note at
	90000 to get larger score.

	Timing window:
	diff: perfect/great/nice/bad
	DB: 80ms / 120ms / 150ms / 180ms
	RG: 80ms / 120ms / 150ms / 180ms
	PR: 70ms / 90ms / 110ms / 140ms
	MS: 60ms / 80ms / 100ms / 130ms
	RMS: 60ms / 80ms / 100ms / 130ms
	holds are 150ms iirc just like flicks
	Flick notes window: 150ms / 180ms / 190ms / 200ms
	Slide checkpoint/entry/exit window: 200ms <-- note (!) slide can only be perfect/miss
	So note at 90000 can actually still be perfect when tapped on 89940 - 90060.
	Expected score is taken from rehearsal demo mode. Maybe in that mode, note is not exactly tap'd
	on the note timestamp?
	Also for some test, data is taken from leaderboard, so the guy might not be very optimal with their timing window abuse
	*/
	scoreThreshold := 10000.0
	appealThreshold := 50.0

	for i, tc := range testcases {
		helper.Features.SetWindowAbuse(tc.windowAbuse)
		// if i != len(testcases)-1 {
		// 	continue
		// }
		guest := tc.guestData.toOwnedCard(cm)
		ocards := []*usermodel.OwnedCard{}
		for _, card := range tc.cardsData {
			ocards = append(ocards, card.toOwnedCard(cm))
		}
		supports := []*usermodel.OwnedCard{}
		for _, card := range tc.supportsData {
			supports = append(supports, card.toOwnedCard(cm))
		}
		song := sm.Filter().NameLike(tc.songName).Difficulty(tc.diff).First()
		gc := simulatormodels.NewGameConfig(ocards, tc.leadIndex, supports, guest, song, tc.supportAppeals, tc.statCalc)
		assert.InDeltaf(t, tc.expectedAppeals, gc.GetAppeal(), appealThreshold, "Wrong appeal for test #%d", i)
		game := NewGameFast(gc)
		result := game.Play(tc.skillAlwaysOn)
		assert.InDeltaf(t, tc.expectedScore, result.Score, scoreThreshold, "Wrong score for test #%d", i)
	}
}
