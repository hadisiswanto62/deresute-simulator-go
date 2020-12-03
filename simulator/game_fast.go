package simulator

import (
	"fmt"
	"math"
	"sort"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
)

func NewGameFast(c simulatormodels.Playable) *GameFast {
	game := GameFast{
		Config: c,
	}
	game.songDifficultyMultiplier = helper.GetSongDifficultyMultiplier(c.GetSong().Level)
	game.comboBonusMap = getComboBonusMap(c.GetSong().NotesCount())
	game.windowAbuse = helper.Features.UseWindowAbuse()
	return &game
}

type GameFast struct {
	Config simulatormodels.Playable

	songDifficultyMultiplier float64
	comboBonusMap            map[int]float64
	windowAbuse              bool
}

func (g GameFast) printState(state *GameState, i, timestamp int,
	activeSkillsIndex []int, noteType []enum.NoteType, tapScore int, sb, cb,
	baseComboBonus float64) string {
	return fmt.Sprintf("%d Note #%d [%v]: %d/%d (hp=%d), (skill=+%.2f/%.2f) (combo=+%.5f). activeSkills = %v\n",
		timestamp, i, noteType, tapScore, state.Score, state.currentHp,
		sb, cb, baseComboBonus, activeSkillsIndex,
	)
}

func (g GameFast) Play(alwaysGoodRolls bool, seed int) *GameState {
	state := initConfig(g.Config)
	state.alwaysGoodRolls = alwaysGoodRolls
	activeSkillsData := rollSkill(state, seed)
	activeSkills := activeSkillsData.activeSkillTimestamps
	hpCosts := activeSkillsData.hpCostTimestamps
	for i, note := range state.song.Notes {
		timestamp := note.TimestampMs
		noteType := note.NoteType
		hpCost := g.getHpCost(timestamp, hpCosts)
		state.currentHp -= hpCost
		if state.currentHp <= 0 {
			state.currentHp = 1
		}

		activeSkillsIndex := g.getActiveSkillsOn(timestamp, &activeSkills, noteType)
		state.concentrationOn = false
		for _, ID := range activeSkillsIndex {
			skill := state.skillActivableCards[ID].Card.Skill.SkillType.Name
			if skill == enum.SkillTypeConcentration {
				state.concentrationOn = true
			}
		}

		// Play note
		judgement := getTapJudgement(state)
		tapHeal := g.getTapHeal(activeSkillsIndex, state, judgement, noteType)
		state.currentHp += tapHeal
		sb, cb := g.getScoreAndComboBonus(activeSkillsIndex, state, judgement, noteType)
		noteScoreMultiplier := g.songDifficultyMultiplier *
			getJudgementScoreMultiplier(judgement) *
			g.comboBonusMap[i] *
			float64(sb*cb) / 10000.0
		score := int(math.Round(noteScoreMultiplier * state.baseTapScore))
		state.Score += score
		// fmt.Println(g.printState(state, i, timestamp, activeSkillsIndex, noteType, score, float64(sb), float64(cb), g.comboBonusMap[i]))
	}
	return state
}

func (g GameFast) getHpCost(timestamp int, hpCosts []*hpCostTimestamp) int {
	if len(hpCosts) == 0 {
		return 0
	}
	cost := 0
	for len(hpCosts) > 0 && hpCosts[0].timestamp < timestamp {
		cost += hpCosts[0].hpCost
		hpCosts = hpCosts[1:]
	}
	return cost
}

func (g GameFast) getTapHeal(activeSkillsIndex []int, state *GameState, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
	maxHeal := 0
	maxHealBonus := 0.0
	for _, id := range activeSkillsIndex {
		ocard := state.skillActivableCards[id]
		heal := ocard.Skill.SkillType.TapHeal(
			ocard.Card.Rarity.Rarity,
			judgement, noteTypes,
		)
		healBonus := 0.0
		if ocard.Skill.SkillType.TapHealBonus != nil {
			healBonus = ocard.Skill.SkillType.TapHealBonus()
		}
		if heal > maxHeal {
			maxHeal = heal
		}
		maxHealBonus = math.Max(maxHealBonus, healBonus)
	}
	heal := int(math.Ceil(float64(maxHeal) * (1.0 + maxHealBonus)))

	if state.currentHp+heal > state.maxHp {
		heal = state.maxHp - state.currentHp
	}
	return heal
}

func (g GameFast) getScoreAndComboBonus(activeCardIds []int, state *GameState,
	judgement enum.TapJudgement, noteTypes []enum.NoteType) (int, int) {
	DELTA := 0.0001
	maxScoreBonus := 0.0
	maxComboBonus := 0.0
	maxBonusBonus := 0.0
	// get bonusBonus first because for reso, we multiply by bonusBonus first before adding
	for _, id := range activeCardIds {
		ocard := state.skillActivableCards[id]
		bonusBonus := 0.0
		if ocard.Skill.SkillType.ScoreComboBonusBonus != nil {
			bonusBonus = ocard.Skill.SkillType.ScoreComboBonusBonus(
				ocard.Card.Idol.Attribute,
			)
		}
		if state.resonantOn {
			maxBonusBonus += bonusBonus
		} else {
			maxBonusBonus = math.Max(bonusBonus, maxBonusBonus)
		}
	}

	altIndices := []int{}
	refIndices := []int{}
	sbForCache := 0.0
	cbForCache := 0.0
	// process other skills
	for _, id := range activeCardIds {
		ocard := state.skillActivableCards[id]
		if ocard.Skill.SkillType.Name == enum.SkillTypeAlternate {
			altIndices = append(altIndices, id)
			continue
		} else if ocard.Skill.SkillType.Name == enum.SkillTypeRefrain {
			refIndices = append(refIndices, id)
			continue
		}

		scoreBonus := ocard.Skill.SkillType.ScoreBonus(
			ocard.Card.Rarity.Rarity,
			state.baseVisual,
			state.baseDance,
			state.baseVocal,
			judgement,
			noteTypes,
		)
		comboBonus := ocard.Skill.SkillType.ComboBonus(
			ocard.Card.Rarity.Rarity,
			state.currentHp,
			judgement,
			noteTypes,
		)
		sb := math.Ceil(scoreBonus*(1+maxBonusBonus)*100.0-DELTA) / 100
		cb := math.Ceil(comboBonus*(1+maxBonusBonus)*100.0-DELTA) / 100
		sbForCache = math.Max(sbForCache, scoreBonus)
		cbForCache = math.Max(cbForCache, comboBonus)
		if state.resonantOn {
			maxScoreBonus += sb
			maxComboBonus += cb
		} else {
			maxScoreBonus = math.Max(sb, maxScoreBonus)
			maxComboBonus = math.Max(cb, maxComboBonus)
		}
	}
	// Remove non-activateds from refrain cache
	for id := range state.caches.refCache {
		found := false
		for _, activatedID := range activeCardIds {
			if id == activatedID {
				found = true
				break
			}
		}
		if !found {
			resetRefrainCache(&state.caches, id)
		}
	}
	// Process Refrain
	for q := 0; q < len(refIndices); q++ {
		sb, cb := handleRefrain(&state.caches, judgement, noteTypes, refIndices[q])
		sb = math.Ceil(sb*(1+maxBonusBonus)*100.0-DELTA) / 100
		cb = math.Ceil(cb*(1+maxBonusBonus)*100.0-DELTA) / 100
		if state.resonantOn {
			maxScoreBonus += sb
			maxComboBonus += cb
		} else {
			maxScoreBonus = math.Max(sb, maxScoreBonus)
			maxComboBonus = math.Max(cb, maxComboBonus)
		}
	}
	// update cache
	updateCache(&state.caches, noteTypes, sbForCache, cbForCache)
	// process Alternate
	for q := 0; q < len(altIndices); q++ {
		sb := handleAlternate(&state.caches, judgement, noteTypes)
		sb = math.Ceil(sb*(1+maxBonusBonus)*100.0-DELTA) / 100
		if state.resonantOn {
			maxScoreBonus += sb
		} else {
			maxScoreBonus = math.Max(sb, maxScoreBonus)
		}
		if maxComboBonus == 0.0 {
			maxComboBonus = -0.2
		}
	}
	sb := 100 + int(math.Round(maxScoreBonus*100))
	cb := 100 + int(math.Round(maxComboBonus*100))
	return sb, cb
}

// assuming allSkillTimestamps is sorted by startTimestamp
func (g GameFast) getActiveSkillsOn(timestamp int, allSkillTimestamps *[]*activeSkillTimestamp, noteType []enum.NoteType) []int {
	windowAbuse := 0
	// TODO: check this
	if g.windowAbuse {
		switch noteType[0] {
		case enum.NoteTypeTap:
			windowAbuse = 60
		case enum.NoteTypeHold:
			windowAbuse = 150
		case enum.NoteTypeFlick:
			windowAbuse = 150
		case enum.NoteTypeSlide:
			windowAbuse = 200
		}
	}
	ret := []int{}
	// if skill ends in the past, it is inactive
	// windowAbuse -> can tap on timestamp-windowAbuse to get active skill
	for len(*allSkillTimestamps) > 0 && (*allSkillTimestamps)[0].endTimestamp < timestamp-windowAbuse {
		*allSkillTimestamps = (*allSkillTimestamps)[1:]
	}
	for _, activeSkill := range *allSkillTimestamps {
		// because it is sorted, if we reach to future skill, break
		// windowAbuse -> can tap on timestamp+windowAbuse to get active skill
		if activeSkill.startTimestamp > timestamp+windowAbuse {
			break
		}
		// if skill ends in the future, it is active (!)
		//windowAbuse -> can tap on timestamp-windowAbuse to get active skill
		if activeSkill.endTimestamp >= timestamp-windowAbuse {
			ret = append(ret, activeSkill.cardIndex)
		}
	}
	return ret
}

func rollSkill(state *GameState, seed int) activeSkillData {
	activeSkillTimestamps := []*activeSkillTimestamp{}
	hpCostTimestamps := []*hpCostTimestamp{}

	timestampLimit := state.song.Notes[state.song.NotesCount()-1].TimestampMs
	for i, ocard := range state.skillActivableCards {
		if !state.skillsActive[i] {
			continue
		}

		probMultiplier := 1.0
		if ocard.Card.Idol.Attribute == state.song.Attribute || state.song.Attribute == enum.AttrAll {
			probMultiplier += 0.3
		}
		for _, leader := range state.leadSkillActivableCards {
			probMultiplier += leader.LeadSkill.SkillProbBonus(
				leader.Card.Rarity.Rarity,
				ocard.Card.Idol.Attribute,
			)
		}

		hpCost := ocard.Card.Skill.ActivationCost
		prob := float64(ocard.SkillProcChance) / 10000.0 * probMultiplier
		duration := ocard.SkillEffectLength * 10
		timer := ocard.Skill.Timer * 1000
		times := 1
		for timestamp := timer; timestamp < timestampLimit-3000; timestamp += timer {
			times++
			if !helper.RollFast(prob, seed+times*131) {
				if !state.alwaysGoodRolls {
					continue
				}
			}
			// this is to match DEMO live, since skill activation is technically after tap,
			// if skill timer and tap at the same tick -> skill should not active (workaround = add 1 millisecond)
			t := timestamp + 1
			activeSkillTimestamps = append(activeSkillTimestamps, &activeSkillTimestamp{
				cardIndex:      i,
				startTimestamp: t,
				endTimestamp:   t + duration,
			})
			if hpCost != 0 {
				hpCostTimestamps = append(hpCostTimestamps, &hpCostTimestamp{
					timestamp: t,
					hpCost:    hpCost,
				})
			}
		}
	}
	sort.SliceStable(activeSkillTimestamps, func(i, j int) bool {
		return activeSkillTimestamps[i].startTimestamp <
			activeSkillTimestamps[j].startTimestamp
	})

	sort.SliceStable(hpCostTimestamps, func(i, j int) bool {
		return hpCostTimestamps[i].timestamp <
			hpCostTimestamps[j].timestamp
	})
	return activeSkillData{
		activeSkillTimestamps: activeSkillTimestamps,
		hpCostTimestamps:      hpCostTimestamps,
	}
}
