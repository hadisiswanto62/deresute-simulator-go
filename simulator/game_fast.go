package simulator

import (
	"math"
	"sort"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
)

func NewGameFast(c Playable) *GameFast {
	game := GameFast{
		Config: c,
	}
	game.songDifficultyMultiplier = getSongDifficultyMultiplier(c.getSong().Level)
	game.comboBonusMap = getComboBonusMap(c.getSong().NotesCount())
	return &game
}

type GameFast struct {
	Config Playable

	songDifficultyMultiplier float64
	comboBonusMap            map[int]float64
}

func (g GameFast) Play(alwaysGoodRolls bool) *GameState {
	state := initConfig(g.Config)
	state.alwaysGoodRolls = alwaysGoodRolls
	activeSkillsData := rollSkill(state)
	activeSkills := activeSkillsData.activeSkillTimestamps
	hpCosts := activeSkillsData.hpCostTimestamps
	for i, note := range state.song.Notes {
		timestamp := note.TimestampMs
		hpCost := g.getHpCost(timestamp, hpCosts)
		state.currentHp -= hpCost
		activeSkillsIndex := g.getActiveSkillsOn(timestamp, activeSkills)

		// Play note
		judgement := getTapJudgement(state)
		noteType := note.NoteType
		scoreComboBonus := g.getScoreAndComboBonus(activeSkillsIndex, state, judgement, noteType)
		noteScoreMultiplier := g.songDifficultyMultiplier *
			getJudgementScoreMultiplier(judgement) *
			g.comboBonusMap[i] *
			scoreComboBonus
		score := int(math.Round(noteScoreMultiplier * state.baseTapScore))
		tapHeal := g.getTapHeal(activeSkillsIndex, state, judgement, noteType)
		state.Score += score
		state.currentHp += tapHeal
		state.printState()
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

func (g GameFast) getTapHeal(activeSkillsIndex []int, state *GameState, judgement enum.TapJudgement, noteType enum.NoteType) int {
	maxHeal := 0
	maxHealBonus := 0.0
	for _, id := range activeSkillsIndex {
		ocard := state.skillActivableCards[id]
		heal := ocard.Skill.SkillType.TapHeal(
			ocard.Card.Rarity.Rarity,
			judgement, noteType,
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

func (g GameFast) getScoreAndComboBonus(activeCardIds []int, state *GameState, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
	maxScoreBonus := 0.0
	maxComboBonus := 0.0
	maxBonusBonus := 0.0

	for _, id := range activeCardIds {
		ocard := state.skillActivableCards[id]
		scoreBonus := ocard.Skill.SkillType.ScoreBonus(
			ocard.Card.Rarity.Rarity,
			state.baseVisual,
			state.baseDance,
			state.baseVocal,
			judgement,
			noteType,
		)
		comboBonus := ocard.Skill.SkillType.ComboBonus(
			ocard.Card.Rarity.Rarity,
			state.currentHp,
			judgement,
			noteType,
		)
		bonusBonus := 0.0
		if ocard.Skill.SkillType.ScoreComboBonusBonus != nil {
			bonusBonus = ocard.Skill.SkillType.ScoreComboBonusBonus(
				ocard.Card.Idol.Attribute,
			)
		}
		if state.resonantOn {
			maxScoreBonus += scoreBonus
			maxComboBonus += comboBonus
			maxBonusBonus += bonusBonus
		} else {
			maxScoreBonus = math.Max(maxScoreBonus, scoreBonus)
			maxComboBonus = math.Max(maxComboBonus, comboBonus)
			maxBonusBonus = math.Max(maxBonusBonus, bonusBonus)
		}
	}
	maxScoreBonus = math.Ceil(maxScoreBonus*(1+maxBonusBonus)*100.0) / 100.0
	maxComboBonus = math.Ceil(maxComboBonus*(1+maxBonusBonus)*100.0) / 100.0
	return (1 + maxScoreBonus) * (1 + maxComboBonus)
}

// assuming allSkillTimestamps is sorted by startTimestamp
func (g GameFast) getActiveSkillsOn(timestamp int, allSkillTimestamps []*activeSkillTimestamp) []int {
	ret := []int{}
	for _, activeSkill := range allSkillTimestamps {
		if activeSkill.startTimestamp > timestamp {
			break
		}
		if activeSkill.endTimestamp >= timestamp {
			ret = append(ret, activeSkill.cardIndex)
		}
	}
	return ret
}

func rollSkill(state *GameState) activeSkillData {
	activeSkillTimestamps := []*activeSkillTimestamp{}
	hpCostTimestamps := []*hpCostTimestamp{}

	timestampLimit := state.song.DurationMs
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
		for timestamp := timer; timestamp < timestampLimit-3000; timestamp += timer {
			if !helper.RollFast(prob) {
				if !state.alwaysGoodRolls {
					continue
				}
			}
			activeSkillTimestamps = append(activeSkillTimestamps, &activeSkillTimestamp{
				cardIndex:      i,
				startTimestamp: timestamp,
				endTimestamp:   timestamp + duration,
			})
			if hpCost != 0 {
				hpCostTimestamps = append(hpCostTimestamps, &hpCostTimestamp{
					timestamp: timestamp,
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
