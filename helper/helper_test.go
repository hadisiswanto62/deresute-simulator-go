package helper

import "testing"

func TestScale(t *testing.T) {
	x := Scale(2465, 4180, 50, 0)
	if x != 2465 {
		t.Errorf("Scale error! want = %d have = %d", 2465, x)
	}
}
