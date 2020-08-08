package enum

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
