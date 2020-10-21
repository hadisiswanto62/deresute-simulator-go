package logic

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogic struct {
	name       string
	isViolated func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool
}

var unisonInCorrectSongType = &gameConfigLogic{
	name: "unisonInCorrectSongType",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {

		for _, ocards := range leadSkillActivableCards {
			for attr, lskill := range enum.LeadSkillUnisonMap {
				if lskill != ocards.LeadSkill.Name {
					continue
				}
				// if use unison:
				// violated if song is ALL (princess is objectively better)
				if song.Attribute == enum.AttrAll {
					return true
				}
				// violated if unison attr does not match song attr
				if attr != song.Attribute {
					return true
				}
			}
		}
		return false
	},
}

var allLeadSkillsActive = &gameConfigLogic{
	name: "allLeadSkillsActive",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		attrs := []enum.Attribute{guest.Card.Idol.Attribute}
		skills := []enum.SkillType{guest.Card.Skill.SkillType.Name}
		for _, ocard := range team.Ocards {
			attrs = append(attrs, ocard.Card.Idol.Attribute)
			skills = append(skills, ocard.Card.Skill.SkillType.Name)
		}
		for _, ocard := range leadSkillActivableCards {
			// violated if any lead skill is inactive
			if !ocard.LeadSkill.IsActive(attrs, skills) {
				return true
			}
		}
		return false
	},
}

var princessWhenShouldBeUnison = &gameConfigLogic{
	name: "princessWhenShouldBeUnison",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		for attr, lskill := range enum.LeadSkillPrincessMap {
			// only check guest because leader should be checked on teamLogic (TODO: )
			if guest.LeadSkill.Name == lskill {
				if song.Attribute == attr {
					// IsViolated when using princess on the attribute song (because princess is +50% while unison is +55%)
					return true
				}
				return false
			}
		}
		return false
	},
}

var skillsAreActive = &gameConfigLogic{
	name: "skillsAreActive",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		attrs := []enum.Attribute{guest.Card.Idol.Attribute}
		for _, ocard := range team.Ocards {
			attrs = append(attrs, ocard.Card.Idol.Attribute)
		}
		for _, ocard := range team.Ocards {
			if !ocard.Card.Skill.SkillType.IsActive(attrs) {
				return true
			}
		}
		return false
	},
}

var guestTriColorCorrectStat = &gameConfigLogic{
	name: "guestTriColorCorrectStat",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		for stat, lskill := range enum.LeadSkillTricolorMap {
			if guest.LeadSkill.Name == lskill {
				da, vo, vi := 0, 0, 0
				for _, ocard := range team.Ocards {
					da += ocard.Dance
					vo += ocard.Vocal
					vi += ocard.Visual
				}

				if stat == enum.StatDance && da >= vo && da >= vi {
					return false
				} else if stat == enum.StatVocal && vo >= da && vo >= vi {
					return false
				} else if stat == enum.StatVisual && vi >= da && vi >= vo {
					return false
				}
				// IsViolated when the team's max stat does not match the guest's tricolor stat
				return true
			}
		}
		return false
	},
}

var _relevant = []enum.LeadSkill{
	enum.LeadSkillCoolPrincess, enum.LeadSkillCutePrincess, enum.LeadSkillPassionPrincess,
	enum.LeadSkillCoolUnison, enum.LeadSkillCuteUnison, enum.LeadSkillPassionUnison,
	enum.LeadSkillTricolorMakeup, enum.LeadSkillTricolorStep, enum.LeadSkillTricolorVoice,
	enum.LeadSkillTricolorAbility,
}

var guestPrincessUnisonCorrectStat = &gameConfigLogic{
	name: "guestPrincessUnisonCorrectStat",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		// if guest is princess/unison/trico:
		// guest highest stat should be = team's highest stat
		found := false
		for _, lskill := range _relevant {
			if guest.Card.LeadSkill.Name == lskill {
				found = true
			}
		}
		if !found {
			return false
		}
		teamStat := make(map[enum.Stat]int)
		total := 0
		for _, ocard := range team.Ocards {
			for stat, value := range ocard.Stats() {
				_, ok := teamStat[stat]
				if !ok {
					teamStat[stat] = 0
				}
				teamStat[stat] += value
				total += value
			}
		}
		var maxStat enum.Stat
		maxValue := 0
		for stat, value := range teamStat {
			if value > maxValue {
				maxValue = value
				maxStat = stat
			}
		}

		return maxStat != guest.MaxStat()
	},
}

var tricolorCorrectColor = &gameConfigLogic{
	name: "tricolorCorrectColor",
	isViolated: func(team *usermodel.Team, leadSkillActivableCards []*usermodel.OwnedCard, guest *usermodel.OwnedCard, song *models.Song) bool {
		// if team is one color -> always false
		guestIsTriColor := false
		for _, lskill := range enum.LeadSkillTricolorMap {
			if guest.LeadSkill.Name == lskill {
				guestIsTriColor = true
			}
		}
		// Not violated if guest is not tricolor
		if !guestIsTriColor {
			return false
		}
		teamAttrs := map[enum.Attribute]bool{
			enum.AttrCool:    false,
			enum.AttrCute:    false,
			enum.AttrPassion: false,
		}
		attrCount := 0
		for _, ocard := range team.Ocards {
			attr := ocard.Card.Idol.Attribute
			res, _ := teamAttrs[attr]
			if !res {
				attrCount++
			}
			teamAttrs[ocard.Card.Idol.Attribute] = true
		}

		guestAttr := guest.Card.Idol.Attribute
		// IsViolated when team is 2 color AND guest is tricolor AND guest is not the remaining color
		if attrCount == 2 {
			var leftover enum.Attribute
			for attr, val := range teamAttrs {
				if !val {
					leftover = attr
					break
				}
			}
			return guestAttr != leftover
		}
		// IsViolated when team is already 3 color AND guest is tricolor AND guest attr is not CUTE (because attr should not matter)
		// TODO: check if lead is attr specific stat up
		if attrCount == 3 {
			return guestAttr != enum.AttrCute
		}
		// Else not violated
		return false
	},
}
