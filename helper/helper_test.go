package helper

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	count := 0
	for i := 0; i < 100; i++ {
		if Roll(0.1) {
			count++
		}
	}
	if (count > 30) || (count == 0) {
		t.Errorf("RNG is suspiciously wrong. want = %d, have = %d", 30, count)
	}
}

func TestRandomSafe(t *testing.T) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for i := 0; i < 100; i++ {
		if RollSafe(0.1, generator) {
			count++
		}
	}
	if (count > 30) || (count == 0) {
		t.Errorf("RNG is suspiciously wrong. want = %d, have = %d", 30, count)
	}
}

func TestScale(t *testing.T) {
	x := Scale(2465, 4180, 50, 1)
	if x != 2465 {
		t.Errorf("Scale error! want = %d have = %d", 2465, x)
	}

	x = Scale(2317, 4471, 90, 90)
	if x != 4471 {
		t.Errorf("Rounding error! want = %d have = %d", 4471, x)
	}
}
