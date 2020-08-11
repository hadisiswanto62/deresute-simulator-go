package models

import (
	"fmt"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/stretchr/testify/assert"
)

func sampleSong(attr enum.Attribute) Song {
	return NewDefaultSong("Default Song", 26, attr, 120000, 200)
}

func TestDefault(t *testing.T) {
	song := sampleSong(enum.AttrAll)
	assert.Equal(t, song.NotesCount(), 200)
	fmt.Println(song.Notes)
}
