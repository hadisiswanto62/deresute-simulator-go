package simulator

import (
	"math"
	"testing"
)

func TestNewGame(t *testing.T) {
	config := sampleGameConfig()
	game := NewGame(config, false)
	if game == nil {
		t.Errorf("Cannot create a game!")
	}
}

func Test_getSongDifficultyMultiplier(t *testing.T) {
	should := map[int]float64{
		5: 1, 6: 1.025, 7: 1.05, 8: 1.075, 9: 1.1,
		10: 1.2, 11: 1.225, 12: 1.25, 13: 1.275, 14: 1.3,
		15: 1.4, 16: 1.425, 17: 1.45, 18: 1.475, 19: 1.5,
		20: 1.6, 21: 1.65, 22: 1.7, 23: 1.75, 24: 1.8, 25: 1.85, 26: 1.9, 27: 1.95, 28: 2, 29: 2.1, 30: 2.2,
	}
	for level, want := range should {
		if have := getSongDifficultyMultiplier(level); math.Abs(want-have) >= 0.00001 {
			t.Errorf("Invalid song difficulty multiplier for level (%d). want = %f have = %f", level, want, have)
		}
	}
}

func Benchmark_getComboBonusMap(b *testing.B) {
	for j := 0; j < b.N; j++ {
		noteCount := 477
		getComboBonusMap(noteCount)

	}
}
