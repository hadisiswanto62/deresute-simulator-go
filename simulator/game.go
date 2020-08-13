package simulator

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

const (
	tickRate  = 30
	greatProb = 0.0
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

// GameState represents the state of the game. Including Score and Log for exports
type GameState struct {
	timestamp            int
	Score                int
	leadSkillActive      bool
	guestLeadSkillActive bool
	currentHp            int
	currentNoteIndex     int
	activeSkills         []*activeSkill
	teamAttributes       [6]enum.Attribute
	Log                  []string
	randomGenerator      *rand.Rand
	verbose              bool
}

// PrintLog prints the log for the gamestate
func (gs GameState) PrintLog() {
	for _, item := range gs.Log {
		fmt.Println(item)
	}
}

func (gs *GameState) log(item string) {
	gs.Log = append(gs.Log, item)
}

func (gs *GameState) logf(format string, a ...interface{}) {
	if !gs.verbose {
		return
	}
	text := fmt.Sprintf(format, a...)
	gs.log(text)
	return
}

// Game represents a simulation of a game. (use simulator.NewGame() instead)
type Game struct {
	config *GameConfig

	UseAppealsOnly           bool
	songDifficultyMultiplier float64
	comboBonusMap            map[int]float64
	maxHp                    int
	verbose                  bool
}

func (g *Game) rollSkill(state *GameState) {
	// TODO: remove expired skills and roll skill here
	// remove expired
	var newActiveSkills []*activeSkill
	for _, skill := range state.activeSkills {
		if skill.isActiveOn(state.timestamp) {
			newActiveSkills = append(newActiveSkills, skill)
		} else {
			state.logf("%6d: Skill (%s) is deactivated", state.timestamp, skill)
		}
	}
	state.activeSkills = newActiveSkills
	// roll skill
	// skill can't active in the first loop
	if state.timestamp < tickRate*2 {
		state.logf("%6d: Tried to activate skills but it is on first loop", state.timestamp)
		return
	}
	for i, ocard := range g.config.team.Ocards {
		// if it is not time yet --> skip
		if !(state.timestamp%(ocard.Card.Skill.Timer*1000) < tickRate) {
			// state.logf("Tried to activate %d but it is not time yet", i)
			continue
		}
		// skill can't active within 3 seconds before last note
		if state.timestamp > g.config.song.Notes[g.config.song.NotesCount()-1].TimestampMs-3000 {
			state.logf("%6d: Tried to activate (%d. %s) but it is 3 seconds before last note", state.timestamp,
				i, ocard.Card.Skill.SkillType.Name)
			continue
		}
		// if inactive skill --> skip
		if !ocard.Card.Skill.SkillType.IsActive(state.teamAttributes) {
			state.logf("%6d: Tried to activate (%d. %s) but it is inactive", state.timestamp, i, ocard.Card.Skill.SkillType.Name)
			continue
		}
		// if card is currently active --> skip
		active := false
		for _, activeSkill := range state.activeSkills {
			if ocard == activeSkill.ocard {
				active = true
				break
			}
		}
		if active {
			state.logf("%6d: Tried to activate (%d. %s) but it is currently active", state.timestamp,
				i, ocard.Card.Skill.SkillType.Name)
			continue
		}

		probMultiplier := 1.0
		if (ocard.Card.Idol.Attribute == g.config.song.Attribute) || (g.config.song.Attribute == enum.AttrAll) {
			probMultiplier += 0.3
		}
		if state.leadSkillActive {
			probMultiplier += g.config.team.Leader().LeadSkill.SkillProbBonus(
				g.config.team.Leader().Card.Rarity.Rarity,
				ocard.Card.Idol.Attribute,
			)
		}
		if state.guestLeadSkillActive {
			probMultiplier += g.config.guest.LeadSkill.SkillProbBonus(
				g.config.team.Leader().Card.Rarity.Rarity,
				ocard.Card.Idol.Attribute,
			)
		}
		prob := float64(ocard.SkillProcChance) / 10000.0 * probMultiplier
		if helper.RollSafe(prob, state.randomGenerator) {
			cost := ocard.Card.Skill.ActivationCost
			if cost > 0 {
				hpAfter := state.currentHp - cost
				if hpAfter < 1 {
					hpAfter = 1
				}
				state.currentHp = hpAfter
			}
			state.logf("%6d: %d. %v activated.", state.timestamp, i, ocard.Card.Skill.SkillType.Name)
			state.activeSkills = append(state.activeSkills, &activeSkill{ocard: ocard, cardIndex: i, timestamp: state.timestamp})
		} else {
			state.logf("%6d: Tried to activate (%d. %s) but roll failed", state.timestamp, i, ocard.Card.Skill.SkillType.Name)
		}
	}
}

func (g Game) scoreAndComboBonus(state GameState, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
	maxScoreBonus := 0.0
	maxComboBonus := 0.0
	for _, activeSkill := range state.activeSkills {
		if activeSkill.ocard.Skill.SkillType.IsActive(state.teamAttributes) {
			tmpScoreBonus := activeSkill.ocard.Skill.SkillType.ScoreBonus(
				activeSkill.ocard.Card.Rarity.Rarity,
				g.config.BaseVisual,
				g.config.BaseDance,
				g.config.BaseVocal,
				judgement,
				noteType,
			)
			tmpComboBonus := activeSkill.ocard.Skill.SkillType.ComboBonus(
				activeSkill.ocard.Card.Rarity.Rarity,
				state.currentHp,
				judgement,
				noteType,
			)
			maxScoreBonus = math.Max(maxScoreBonus, tmpScoreBonus)
			maxComboBonus = math.Max(maxComboBonus, tmpComboBonus)
		}
	}
	return (1 + maxScoreBonus) * (1 + maxComboBonus)
}

// Play plays the game and return the final state
func (g Game) Play(seed int64) GameState {
	// defer helper.MeasureTime(time.Now(), "Play")
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	if g.UseAppealsOnly {
		score := int(1.41 * g.songDifficultyMultiplier * float64(g.config.Appeal))
		return GameState{Score: score}
	}

	teamAttributes := g.config.getTeamAttributes()
	state := GameState{
		timestamp:            0,
		leadSkillActive:      g.config.team.Leader().LeadSkill.IsActive(teamAttributes),
		guestLeadSkillActive: g.config.guest.LeadSkill.IsActive(teamAttributes),
		currentNoteIndex:     -1,
		currentHp:            g.config.Hp,
		teamAttributes:       teamAttributes,
		randomGenerator:      rand.New(rand.NewSource(seed)),
		verbose:              g.verbose,
	}
	state.logf("Playing with appeal %d:", g.config.Appeal)
	for state.timestamp < g.config.song.DurationMs {
		g.rollSkill(&state)
		for i := state.currentNoteIndex + 1; i < g.config.song.NotesCount(); i++ {
			if g.config.song.Notes[i].TimestampMs > state.timestamp {
				break
			}

			// Play note here
			judgement := enum.TapJudgementPerfect
			if helper.RollSafe(greatProb, state.randomGenerator) {
				judgement = enum.TapJudgementGreat
			}
			noteType := g.config.song.Notes[i].NoteType
			scoreComboBonus := g.scoreAndComboBonus(state, judgement, noteType)
			noteScoreMultiplier := g.songDifficultyMultiplier *
				getJudgementScoreMultiplier(judgement) *
				g.comboBonusMap[i] *
				scoreComboBonus /
				float64(g.config.song.NotesCount())

			// TODO: Tap heal here
			tapHeal := 0
			for _, activeSkill := range state.activeSkills {
				heal := activeSkill.ocard.Skill.SkillType.TapHeal(
					activeSkill.ocard.Card.Rarity.Rarity,
					judgement, noteType,
				)
				if heal > tapHeal {
					tapHeal = heal
				}
			}
			state.currentHp += tapHeal
			if state.currentHp > g.maxHp {
				state.currentHp = g.maxHp
			}

			score := int(math.Ceil(noteScoreMultiplier * float64(g.config.Appeal)))
			state.Score += score
			state.logf("%6d: Note %d tapped for %d/%d. Hp is %d (scoreComboBonus = %f)",
				state.timestamp, i, score, state.Score, state.currentHp, scoreComboBonus,
			)

			state.currentNoteIndex = i
		}
		state.timestamp += tickRate
	}
	return state
}

// NewGame creates new game
func NewGame(config *GameConfig, verbose bool) *Game {
	return &Game{
		config:                   config,
		songDifficultyMultiplier: getSongDifficultyMultiplier(config.song.Level),
		comboBonusMap:            getComboBonusMap(config.song.NotesCount()),
		maxHp:                    config.Hp * 2,
		verbose:                  verbose,
	}
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
		if progress >= 90.0 {
			comboMap[i] = 2.0
		} else if progress >= 80.0 {
			comboMap[i] = 1.7
		} else if progress >= 70.0 {
			comboMap[i] = 1.5
		} else if progress >= 50.0 {
			comboMap[i] = 1.4
		} else if progress >= 25.0 {
			comboMap[i] = 1.3
		} else if progress >= 10.0 {
			comboMap[i] = 1.2
		} else if progress >= 5.0 {
			comboMap[i] = 1.1
		} else {
			comboMap[i] = 1.0
		}
	}
	return comboMap
}

func getSongDifficultyMultiplier(songLevel int) float64 {
	if songLevel <= 9 {
		return 1 + (0.025 * float64((songLevel - 5)))
	} else if songLevel <= 14 {
		return 1.2 + (0.025 * float64((songLevel - 10)))
	} else if songLevel <= 19 {
		return 1.4 + (0.025 * float64((songLevel - 15)))
	} else if songLevel <= 28 {
		return 1.6 + (0.05 * float64((songLevel - 20)))
	} else if songLevel <= 30 {
		return 2 + (0.1 * float64((songLevel - 28)))
	} else {
		return 1
	}
}
