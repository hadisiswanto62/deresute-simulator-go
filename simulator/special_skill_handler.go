package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// not to be updated outside this file
type caches struct {
	alternateScoreBonusCache map[enum.NoteType]float64
}

// does not check if key exist
func updateAlternateCache(cache *caches, noteTypes []enum.NoteType, sb float64) {
	for _, noteType := range noteTypes {
		if cache.alternateScoreBonusCache[noteType] < sb {
			cache.alternateScoreBonusCache[noteType] = sb
		}
	}
}

func handleAlternate(cache *caches,
	judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
	sb := 0.0
	for _, noteType := range noteTypes {
		if cache.alternateScoreBonusCache[noteType] > sb {
			sb = cache.alternateScoreBonusCache[noteType]
		}
	}
	return sb * 1.5
}
