package csvmodels

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

type TmpOwnedCardRawData struct {
	CardID     int `csv:"card_id"`
	SkillLevel int `csv:"skill_level"`
	StarRank   int `csv:"star_rank"`
	PotVisual  int `csv:"pot_visual"`
	PotDance   int `csv:"pot_dance"`
	PotVocal   int `csv:"pot_vocal"`
	PotHp      int `csv:"pot_hp"`
	PotSkill   int `csv:"pot_skill"`
}

type TmpNoteRawData struct {
	ID        int     `csv:"id"`
	Sec       float64 `csv:"sec"`
	Type      int     `csv:"type"`
	StartPos  int     `csv:"startPos"`
	FinishPos int     `csv:"finishPos"`
	Status    int     `csv:"status"`
	Sync      int     `csv:"sync"`
	GroupID   int     `csv:"groupId"`
}

func (note TmpNoteRawData) toNote() *models.Note {
	/* TODO:
	handle
		- hold and slide
		- slide and hold
		- slide and flick
	*/
	noteType := []enum.NoteType{}
	if note.Type == 1 {
		if note.Status == 0 {
			noteType = append(noteType, enum.NoteTypeTap)
		} else {
			noteType = append(noteType, enum.NoteTypeFlick)
		}
	} else if note.Type == 2 {
		noteType = append(noteType, enum.NoteTypeHold)
		if note.Status == 1 || note.Status == 2 {
			noteType = append(noteType, enum.NoteTypeFlick)
		}
	} else if note.Type == 3 {
		noteType = append(noteType, enum.NoteTypeSlide)
	} else {
		return nil
	}
	return &models.Note{
		TimestampMs: int(note.Sec * 1000),
		NoteType:    noteType,
	}
}

type TmpSongDifficultyId struct {
	SongID      int `csv:"id"`
	MusicDataID int `csv:"music_data_id"`
	Type        int `csv:"type"`
}

type TmpSongDifficultyLevel struct {
	SongDifficultyID int `csv:"id"`
	SongID           int `csv:"live_data_id"`
	DifficultyInt    int `csv:"difficulty_type"`
	Level            int `csv:"level_vocal"`
}

type TmpMusicNameData struct {
	MusicDataID int    `csv:"id"`
	Name        string `csv:"name"`
	Length      int    `csv:"sound_length"`
}
