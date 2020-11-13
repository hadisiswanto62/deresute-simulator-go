package simulator

import (
	"fmt"
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator/simulatormodels"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

const (
	tickRate               = 30
	greatProb              = 0.0
	concentrationGreatProb = 0.0
)

type activeSkill struct {
	ocard     *usermodel.OwnedCard
	cardIndex int
	timestamp int
}

func (as activeSkill) isActiveOn(timestamp int) bool {
	endTimestamp := as.ocard.SkillEffectLength*10 + as.timestamp
	return timestamp <= endTimestamp
}

func (as activeSkill) String() string {
	return fmt.Sprintf("%d. %s", as.cardIndex, as.ocard.Card.Skill.SkillType.Name)
}

// Game is a game. USE NewGame TO CREATE.
type Game struct {
	Config simulatormodels.Playable

	songDifficultyMultiplier float64
	comboBonusMap            map[int]float64
}

// NewGame creates, initializes, and returns Game
func NewGame(c simulatormodels.Playable) *Game {
	game := Game{
		Config: c,
	}
	game.songDifficultyMultiplier = helper.GetSongDifficultyMultiplier(c.GetSong().Level)
	game.comboBonusMap = getComboBonusMap(c.GetSong().NotesCount())
	return &game
}

// GameState represents internal state of the game. also includes Score
type GameState struct {
	timestamp        int
	Score            int
	currentNoteIndex int
	currentHp        int

	leadSkillsActive []bool
	skillsActive     []bool
	concentrationOn  bool
	resonantOn       bool

	song *models.Song

	activeSkills            []*activeSkill
	skillActivableCards     []*usermodel.OwnedCard
	leadSkillActivableCards []*usermodel.OwnedCard
	appeal                  int
	baseVisual              int
	baseVocal               int
	baseDance               int
	maxHp                   int
	baseTapScore            float64

	alwaysGoodRolls bool
	caches
}

func (s GameState) printState() {
	// 123000 Note #1: 12345 (hp=10). activeSkillsIndex = [] conc=bool
	activeSkillsIndex := []int{}
	for _, skill := range s.activeSkills {
		activeSkillsIndex = append(activeSkillsIndex, skill.cardIndex)
	}
	fmt.Printf("%d Note #%d: %d (hp=%d). activeSkills = %v conc=%v\n",
		s.timestamp, s.currentNoteIndex, s.Score, s.currentHp,
		activeSkillsIndex, s.concentrationOn,
	)
}

func (g Game) rollSkill(state *GameState, seed int) {
	state.activeSkills = expireOldSkills(state.activeSkills, state.timestamp)

	// skill can't activate in the first loop
	if state.timestamp < tickRate*2 {
		return
	}
	// if within 3 seconds before last note -> skip
	if state.timestamp > state.song.Notes[state.song.NotesCount()-1].TimestampMs-3000 {
		return
	}

	for i, ocard := range state.skillActivableCards {
		// if inactive -> skip
		if !state.skillsActive[i] {
			continue
		}
		// if not time yet -> skip
		if !(state.timestamp%(ocard.Card.Skill.Timer*1000) < tickRate) {
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
		prob := float64(ocard.SkillProcChance) / 10000.0 * probMultiplier
		if !helper.RollFast(prob, 0) {
			if !state.alwaysGoodRolls {
				continue
			}
		}

		hpCost := ocard.Card.Skill.ActivationCost
		if hpCost > 0 {
			hpAfter := state.currentHp - hpCost
			if hpAfter < 1 {
				hpAfter = 1
			}
			state.currentHp = hpAfter
		}
		state.activeSkills = append(state.activeSkills, &activeSkill{
			ocard:     ocard,
			cardIndex: i,
			timestamp: state.timestamp,
		})
	}

	state.concentrationOn = isSkillActive(state.activeSkills, enum.SkillTypeConcentration)
}

// Play plays the game and return the state
func (g Game) Play(alwaysGoodRolls bool, seed int) *GameState {
	state := initConfig(g.Config)
	state.alwaysGoodRolls = alwaysGoodRolls
	for state.timestamp < state.song.DurationMs {
		g.rollSkill(state, 0)
		for i := state.currentNoteIndex + 1; i < state.song.NotesCount(); i++ {
			if state.song.Notes[i].TimestampMs > state.timestamp {
				break
			}

			// Play note here
			judgement := getTapJudgement(state)
			noteType := state.song.Notes[i].NoteType
			scoreComboBonus := g.getScoreAndComboBonus(state, judgement, noteType)
			noteScoreMultiplier := g.songDifficultyMultiplier *
				getJudgementScoreMultiplier(judgement) *
				g.comboBonusMap[i] *
				scoreComboBonus

			score := int(math.Round(noteScoreMultiplier * state.baseTapScore))
			tapHeal := g.getTapHeal(state, judgement, noteType)
			state.Score += score
			state.currentHp += tapHeal

			state.currentNoteIndex = i
			// state.printState()
		}
		state.timestamp += tickRate
	}
	return state
}

func initConfig(c simulatormodels.Playable) *GameState {
	teamAttributes := c.GetTeamAttributesv2()
	teamSkills := c.GetTeamSkillsv2()

	resonantOn := c.IsResonantActive()
	leadSkillsActive := make([]bool, 0, 2)
	for _, ocard := range c.GetLeadSkillActivableCards() {
		active := ocard.LeadSkill.IsActive(teamAttributes, teamSkills)
		leadSkillsActive = append(leadSkillsActive, active)
	}

	skillsActive := make([]bool, 0, 5)
	for _, ocard := range c.GetSkillActivableCards() {
		active := ocard.Skill.SkillType.IsActive(teamAttributes)
		skillsActive = append(skillsActive, active)
	}

	state := GameState{
		timestamp:        0,
		currentNoteIndex: -1,
		currentHp:        c.GetHp(),

		leadSkillsActive:        leadSkillsActive,
		skillsActive:            skillsActive,
		concentrationOn:         false,
		resonantOn:              resonantOn,
		song:                    c.GetSong(),
		skillActivableCards:     c.GetSkillActivableCards(),
		leadSkillActivableCards: c.GetLeadSkillActivableCards(),
		appeal:                  c.GetAppeal(),
		baseVisual:              c.GetBaseVisual(),
		baseVocal:               c.GetBaseVocal(),
		baseDance:               c.GetBaseDance(),
		maxHp:                   c.GetHp() * 2,
		caches: caches{
			alternateScoreBonusCache: map[enum.NoteType]float64{
				enum.NoteTypeFlick: 0.0,
				enum.NoteTypeHold:  0.0,
				enum.NoteTypeSlide: 0.0,
				enum.NoteTypeTap:   0.0,
			},
		},
	}
	state.baseTapScore = float64(state.appeal) / float64(state.song.NotesCount())
	return &state
}

func expireOldSkills(skills []*activeSkill, timestamp int) []*activeSkill {
	newActiveSkills := []*activeSkill{}
	for _, skill := range skills {
		if skill.isActiveOn(timestamp) {
			newActiveSkills = append(newActiveSkills, skill)
		}
	}
	return newActiveSkills
}

func isSkillActive(skills []*activeSkill, skillType enum.SkillType) bool {
	for _, skill := range skills {
		if skill.ocard.Card.Skill.SkillType.Name == skillType {
			return true
		}
	}
	return false
}

func getTapJudgement(state *GameState) enum.TapJudgement {
	judgement := enum.TapJudgementPerfect
	var prob float64
	if state.concentrationOn {
		prob = concentrationGreatProb
	} else {
		prob = greatProb
	}

	if helper.RollFast(prob, 0) && !state.alwaysGoodRolls {
		judgement = enum.TapJudgementGreat
	}
	return judgement
}

func (g Game) getScoreAndComboBonus(state *GameState, judgement enum.TapJudgement, noteTypes []enum.NoteType) float64 {
	maxScoreBonus := 0.0
	maxComboBonus := 0.0
	maxBonusBonus := 0.0

	for _, activeSkill := range state.activeSkills {
		scoreBonus := activeSkill.ocard.Skill.SkillType.ScoreBonus(
			activeSkill.ocard.Card.Rarity.Rarity,
			state.baseVisual,
			state.baseDance,
			state.baseVocal,
			judgement,
			noteTypes,
		)
		comboBonus := activeSkill.ocard.Skill.SkillType.ComboBonus(
			activeSkill.ocard.Card.Rarity.Rarity,
			state.currentHp,
			judgement,
			noteTypes,
		)
		bonusBonus := 0.0
		if activeSkill.ocard.Skill.SkillType.ScoreComboBonusBonus != nil {
			bonusBonus = activeSkill.ocard.Skill.SkillType.ScoreComboBonusBonus(activeSkill.ocard.Card.Idol.Attribute)
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

func (g Game) getTapHeal(state *GameState, judgement enum.TapJudgement, noteTypes []enum.NoteType) int {
	maxHeal := 0
	maxHealBonus := 0.0
	for _, activeSkill := range state.activeSkills {
		heal := activeSkill.ocard.Skill.SkillType.TapHeal(
			activeSkill.ocard.Card.Rarity.Rarity,
			judgement, noteTypes,
		)
		healBonus := 0.0
		if activeSkill.ocard.Skill.SkillType.TapHealBonus != nil {
			healBonus = activeSkill.ocard.Skill.SkillType.TapHealBonus()
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

func getJudgementScoreMultiplier(judgement enum.TapJudgement) float64 {
	switch judgement {
	case enum.TapJudgementPerfect:
		return 1.0
	case enum.TapJudgementGreat:
		return 0.7
	case enum.TapJudgementNice:
		return 0.4
	case enum.TapJudgementBad:
		return 0.1
	case enum.TapJudgementMiss:
		return 0
	}
	return 0
}
func getComboBonusMap(notesCount int) map[int]float64 {
	comboMap := make(map[int]float64)
	for i := 0; i < notesCount; i++ {
		progress := float64(i+2) / float64(notesCount) * 100.0
		if progress > 90.0 {
			comboMap[i] = 2.0
		} else if progress > 80.0 {
			comboMap[i] = 1.7
		} else if progress > 70.0 {
			comboMap[i] = 1.5
		} else if progress > 50.0 {
			comboMap[i] = 1.4
		} else if progress > 25.0 {
			comboMap[i] = 1.3
		} else if progress > 10.0 {
			comboMap[i] = 1.2
		} else if progress > 5.0 {
			comboMap[i] = 1.1
		} else {
			comboMap[i] = 1.0
		}
	}
	return comboMap
}
