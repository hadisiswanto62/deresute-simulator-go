package helper

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d, ", RandInt(10, 13))
	}
}
