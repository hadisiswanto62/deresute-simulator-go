package simulator

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
)

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
	skillAlwaysActive    bool
	concentrationOn      bool
	resonantOn           bool
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
	state.concentrationOn = false
	var newActiveSkills []*activeSkill
	for _, skill := range state.activeSkills {
		if skill.isActiveOn(state.timestamp) {
			newActiveSkills = append(newActiveSkills, skill)
			if skill.ocard.Skill.SkillType.Name == enum.SkillTypeConcentration {
				state.concentrationOn = true
			}
			// } else {
			// state.logf("%6d: Skill (%s) is deactivated", state.timestamp, skill)
		}
	}
	state.activeSkills = newActiveSkills
	// roll skill
	// skill can't active in the first loop
	if state.timestamp < tickRate*2 {
		// state.logf("%6d: Tried to activate skills but it is on first loop", state.timestamp)
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
			// state.logf("%6d: Tried to activate (%d. %s) but it is 3 seconds before last note", state.timestamp,
			// 	i, ocard.Card.Skill.SkillType.Name)
			continue
		}
		// if inactive skill --> skip
		if !ocard.Card.Skill.SkillType.IsActive(state.teamAttributes[:]) {
			// state.logf("%6d: Tried to activate (%d. %s) but it is inactive", state.timestamp, i, ocard.Card.Skill.SkillType.Name)
			continue
		}
		// if card is currently active --> skip
		active := false
		for _, activeSkill := range state.activeSkills {
			// if ocard == activeSkill.ocard {
			// 	active = true
			// 	break
			// }
			if i == activeSkill.cardIndex {
				active = true
				break
			}
		}
		if active {
			// state.logf("%6d: Tried to activate (%d. %s) but it is currently active", state.timestamp,
			// 	i, ocard.Card.Skill.SkillType.Name)
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

		if !state.skillAlwaysActive {
			if !helper.RollFast(prob) {
				continue
			}
		}
		cost := ocard.Card.Skill.ActivationCost
		if cost > 0 {
			hpAfter := state.currentHp - cost
			if hpAfter < 1 {
				hpAfter = 1
			}
			state.currentHp = hpAfter
		}
		// state.logf("%6d: %d. %v activated.", state.timestamp, i, ocard.Card.Skill.SkillType.Name)
		state.activeSkills = append(state.activeSkills, &activeSkill{ocard: ocard, cardIndex: i, timestamp: state.timestamp})
	}
}

func (g Game) scoreAndComboBonus(state GameState, judgement enum.TapJudgement, noteType enum.NoteType) float64 {
	maxScoreBonus := 0.0
	maxComboBonus := 0.0

	maxBonusBonus := 0.0
	for _, activeSkill := range state.activeSkills {
		if activeSkill.ocard.Skill.SkillType.IsActive(state.teamAttributes[:]) {
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
			bonusBonus := 0.0
			if activeSkill.ocard.Skill.SkillType.ScoreComboBonusBonus != nil {
				bonusBonus = activeSkill.ocard.Skill.SkillType.ScoreComboBonusBonus()
			}
			if state.resonantOn {
				maxScoreBonus += tmpScoreBonus
				maxComboBonus += tmpComboBonus
			} else {
				maxScoreBonus = math.Max(maxScoreBonus, tmpScoreBonus)
				maxComboBonus = math.Max(maxComboBonus, tmpComboBonus)
			}
			maxBonusBonus = math.Max(maxBonusBonus, bonusBonus)
		}
	}
	maxScoreBonus = math.Ceil(maxScoreBonus*(1+maxBonusBonus)*100.0) / 100.0
	maxComboBonus = math.Ceil(maxComboBonus*(1+maxBonusBonus)*100.0) / 100.0
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
	teamSkills := g.config.getTeamSkills()
	state := GameState{
		timestamp:            0,
		leadSkillActive:      g.config.team.Leader().LeadSkill.IsActive(teamAttributes[:], teamSkills[:]),
		guestLeadSkillActive: g.config.guest.LeadSkill.IsActive(teamAttributes[:], teamSkills[:]),
		currentNoteIndex:     -1,
		currentHp:            g.config.Hp,
		teamAttributes:       teamAttributes,
		randomGenerator:      nil, // rand.New(rand.NewSource(seed)),
		verbose:              g.verbose,
		skillAlwaysActive:    helper.GetSkillAlwaysActive(),
		concentrationOn:      false,
		resonantOn:           g.config.resonantOn(),
	}
	// state.logf("Playing with appeal %d:", g.config.Appeal)
	for state.timestamp < g.config.song.DurationMs {
		g.rollSkill(&state)
		for i := state.currentNoteIndex + 1; i < g.config.song.NotesCount(); i++ {
			if g.config.song.Notes[i].TimestampMs > state.timestamp {
				break
			}

			// Play note here
			judgement := enum.TapJudgementPerfect

			var prob float64
			if state.concentrationOn {
				prob = concentrationGreatProb
			} else {
				prob = greatProb
			}

			if helper.RollFast(prob) {
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
			tapHealBonus := 0.0
			for _, activeSkill := range state.activeSkills {
				heal := activeSkill.ocard.Skill.SkillType.TapHeal(
					activeSkill.ocard.Card.Rarity.Rarity,
					judgement, noteType,
				)
				tmp := 0.0
				if activeSkill.ocard.Skill.SkillType.TapHealBonus != nil {
					tmp = activeSkill.ocard.Skill.SkillType.TapHealBonus()
				}
				tapHealBonus = math.Max(tapHealBonus, tmp)
				if heal > tapHeal {
					tapHeal = heal
				}
			}
			tapHeal = int(math.Ceil(float64(tapHeal) * (1.0 + tapHealBonus)))

			state.currentHp += tapHeal
			if state.currentHp > g.maxHp {
				state.currentHp = g.maxHp
			}

			score := int(math.Round(noteScoreMultiplier * float64(g.config.Appeal)))
			state.Score += score
			// state.logf("%6d: Note %d tapped for %d/%d. (from combo = %.2f, scoreComboBonus = %.2f)",
			// 	state.timestamp, i, score, state.Score, g.comboBonusMap[i], scoreComboBonus,
			// )

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
