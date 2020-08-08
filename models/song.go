package models

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

// Note is a single note in the song
type Note struct {
	Timestamp float64
	NoteType  enum.NoteType
}

// Song is a song
type Song struct {
	Name      string
	Level     uint8
	Attribute enum.Attribute
	Duration  float64
	Notes     []Note
}

// NotesCount is the count of notes in the song
func (s Song) NotesCount() int {
	return len(s.Notes)
}

// NewDefaultSong generates a Song with linearly distributed Notes
func NewDefaultSong(name string, level uint8, attr enum.Attribute, duration float64, notesCount int) Song {
	startBuffer := 1.0
	endBuffer := 1.0
	effectiveDuration := duration - startBuffer - endBuffer
	var notes []Note
	for i := 0; i < notesCount; i++ {
		note := Note{
			Timestamp: startBuffer + (effectiveDuration / float64((notesCount-1)*i)),
			NoteType:  enum.NoteTypeTap,
		}
		notes = append(notes, note)
	}
	return Song{
		Name:      name,
		Level:     level,
		Attribute: attr,
		Duration:  duration,
		Notes:     notes,
	}
}
