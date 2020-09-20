package enum

import "fmt"

// TapJudgement represents a tap's judgement
type TapJudgement string

// All tap judgements
const (
	TapJudgementPerfect TapJudgement = "Perfect"
	TapJudgementGreat   TapJudgement = "Great"
	TapJudgementNice    TapJudgement = "Nice"
	TapJudgementBad     TapJudgement = "Bad"
	TapJudgementMiss    TapJudgement = "Miss"
)

// NoteType is the type of a note
type NoteType string

// All note types
const (
	NoteTypeTap   NoteType = "Tap"
	NoteTypeHold  NoteType = "Hold"
	NoteTypeFlick NoteType = "Flick"
	NoteTypeSlide NoteType = "Slide"
)

type SongDifficulty int

const (
	SongDifficultyDebut      SongDifficulty = 1
	SongDifficultyRegular    SongDifficulty = 2
	SongDifficultyPro        SongDifficulty = 3
	SongDifficultyMaster     SongDifficulty = 4
	SongDifficultyMasterPlus SongDifficulty = 5
)

var allDifficulties = []SongDifficulty{
	SongDifficultyDebut, SongDifficultyRegular,
	SongDifficultyPro, SongDifficultyMaster,
	SongDifficultyMasterPlus,
}

func GetSongDifficulty(i int) (*SongDifficulty, error) {
	for _, diff := range allDifficulties {
		if int(diff) == i {
			return &diff, nil
		}
	}
	return nil, fmt.Errorf("difficulty %d not found. ", i)
}
