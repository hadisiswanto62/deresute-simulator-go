package simulator

import (
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// not to be updated outside this file
type caches struct {
	scoreBonusCache map[enum.NoteType]float64
	comboBonusCache map[enum.NoteType]float64

	refCache map[int]map[enum.NoteType][2]float64
}

// does not check if key exist
func updateCache(cache *caches, noteTypes []enum.NoteType, sb, cb float64) {
	for _, noteType := range noteTypes {
		cache.scoreBonusCache[noteType] = math.Max(cache.scoreBonusCache[noteType], sb)
		cache.comboBonusCache[noteType] = math.Max(cache.comboBonusCache[noteType], cb)
	}
}

// TODO: currently dont care about judgement
func handleAlternate(cache *caches,
	judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
	sb := 0.0
	for _, noteType := range noteTypes {
		if cache.scoreBonusCache[noteType] > sb {
			sb = cache.scoreBonusCache[noteType]
		}
	}
	return sb * 1.5
}

func resetRefrainCache(cache *caches, id int) {
	delete(cache.refCache, id)
}

func handleRefrain(cache *caches, judgement enum.TapJudgement, noteTypes []enum.NoteType, id int) (float64, float64) {
	selectedID := -1
	for idInCache := range cache.refCache {
		if idInCache != id {
			continue
		}
		selectedID = id
		break
	}
	if selectedID == -1 {
		cache.refCache[id] = map[enum.NoteType][2]float64{
			enum.NoteTypeFlick: [2]float64{cache.scoreBonusCache[enum.NoteTypeFlick], cache.comboBonusCache[enum.NoteTypeFlick]},
			enum.NoteTypeHold:  [2]float64{cache.scoreBonusCache[enum.NoteTypeHold], cache.comboBonusCache[enum.NoteTypeHold]},
			enum.NoteTypeSlide: [2]float64{cache.scoreBonusCache[enum.NoteTypeSlide], cache.comboBonusCache[enum.NoteTypeSlide]},
			enum.NoteTypeTap:   [2]float64{cache.scoreBonusCache[enum.NoteTypeTap], cache.comboBonusCache[enum.NoteTypeTap]},
		}
		selectedID = id
	}
	sb := 0.0
	cb := 0.0
	for _, noteType := range noteTypes {
		bonus := cache.refCache[selectedID][noteType]
		sb = math.Max(sb, bonus[0])
		cb = math.Max(cb, bonus[1])
	}
	return sb, cb
}
