package enum

// Stat represents a card's stat key
type Stat string

// All stats
const (
	StatVisual Stat = "Visual"
	StatVocal  Stat = "Vocal"
	StatDance  Stat = "Dance"
)

// AllStats is all valid stats
var AllStats = [3]Stat{StatVisual, StatVocal, StatDance}
