package simulator

import (
	"testing"
)

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
