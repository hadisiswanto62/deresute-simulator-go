package usermodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbum2_Next(t *testing.T) {
	album := NewAlbum2(sampleOcards(10))
	var i int
	for i = 0; album.Next(); i++ {
		album.GetCards()
	}
	assert.Equal(t, i, 252, "wrong number of team generated!")
}

func benchmarkAlbum2_Next(b *testing.B, i int) {
	for j := 0; j < b.N; j++ {
		album := NewAlbum2(sampleOcards(i))
		for album.Next() {
			album.GetCards()
		}
	}
}

func BenchmarkAlbum2_Next6(b *testing.B)  { benchmarkAlbum_Next(b, 6) }
func BenchmarkAlbum2_Next7(b *testing.B)  { benchmarkAlbum_Next(b, 7) }
func BenchmarkAlbum2_Next8(b *testing.B)  { benchmarkAlbum_Next(b, 8) }
func BenchmarkAlbum2_Next9(b *testing.B)  { benchmarkAlbum_Next(b, 9) }
func BenchmarkAlbum2_Next10(b *testing.B) { benchmarkAlbum_Next(b, 10) }
