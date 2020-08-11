package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

// Note is a single note in the song
type Note struct {
	TimestampMs int
	NoteType    enum.NoteType
}

// Song is a song
type Song struct {
	Name       string
	Level      int
	Attribute  enum.Attribute
	DurationMs int
	Notes      []Note
}

// NotesCount is the count of notes in the song
func (s Song) NotesCount() int {
	return len(s.Notes)
}

// NewDefaultSong generates a Song with linearly distributed Notes
func NewDefaultSong(name string, level int, attr enum.Attribute, durationMs int, notesCount int) Song {
	startBuffer := 1000
	endBuffer := 1000
	effectiveDuration := durationMs - startBuffer - endBuffer
	var notes []Note
	for i := 0; i < notesCount; i++ {
		note := Note{
			TimestampMs: startBuffer + (effectiveDuration * i / (notesCount - 1)),
			NoteType:    enum.NoteTypeTap,
		}
		notes = append(notes, note)
	}
	return Song{
		Name:       name,
		Level:      level,
		Attribute:  attr,
		DurationMs: durationMs,
		Notes:      notes,
	}
}
