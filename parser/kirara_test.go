package parser

import (
	"os"
	"testing"
)

func init() {
	os.Chdir("../")
}
func BenchmarkTest100(b *testing.B) {
	for i := 0; i < 100; i++ {
		Parse()
	}
}
