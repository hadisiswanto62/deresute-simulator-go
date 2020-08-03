package helper

import "testing"

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
