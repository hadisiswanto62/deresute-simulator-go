package cardmanager

import (
	"os"
	"testing"
)

func init() {
	os.Chdir("../")
}

func TestCreateDefault(t *testing.T) {
	cm := NewDefault()
	if len(cm.Cards) == 0 {
		t.Errorf("Failed to parse cards.")
	}
}

func BenchmarkCreateDefault10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewDefault()
	}
}
