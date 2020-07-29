package parser

import (
	"os"
	"testing"
)

func init() {
	os.Chdir("../")
}
func BenchmarkTest10(b *testing.B) {
	for i := 0; i < 100; i++ {
		Parse()
	}
}
