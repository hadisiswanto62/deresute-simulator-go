package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/logger"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogic struct {
	Name       string
	IsViolated func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool
}

var unisonTypeMap = map[enum.LeadSkill]enum.Attribute{
	enum.LeadSkillCuteUnison:    enum.AttrCute,
	enum.LeadSkillCoolUnison:    enum.AttrCool,
	enum.LeadSkillPassionUnison: enum.AttrPassion,
}

var princessTypeMap = map[enum.LeadSkill]enum.Attribute{
	enum.LeadSkillCutePrincess:    enum.AttrCute,
	enum.LeadSkillCoolPrincess:    enum.AttrCool,
	enum.LeadSkillPassionPrincess: enum.AttrPassion,
}

var tricolorStatMap = map[enum.LeadSkill]enum.Stat{
	enum.LeadSkillTricolorMakeup: enum.StatVisual,
	enum.LeadSkillTricolorStep:   enum.StatDance,
	enum.LeadSkillTricolorVoice:  enum.StatVocal,
}

var statUpLeadSkillMap = map[enum.Stat][]enum.LeadSkill{
	enum.StatDance: []enum.LeadSkill{
		enum.LeadSkillCuteStep,
		enum.LeadSkillCoolStep,
		enum.LeadSkillPassionStep,
	},
	enum.StatVocal: []enum.LeadSkill{
		enum.LeadSkillCuteVoice,
		enum.LeadSkillCoolVoice,
		enum.LeadSkillPassionVoice,
	},
	enum.StatVisual: []enum.LeadSkill{
		enum.LeadSkillCuteMakeup,
		enum.LeadSkillCoolMakeup,
		enum.LeadSkillPassionMakeup,
	},
}

var resoLeadSkillMap = map[enum.Stat]enum.LeadSkill{
	enum.StatDance:  enum.LeadSkillResonantStep,
	enum.StatVocal:  enum.LeadSkillResonantVoice,
	enum.StatVisual: enum.LeadSkillResonantMakeup,
}

var unisonInCorrectSongType = gameConfigLogic{
	Name: "unisonInCorrectSongType",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		for lskill, attr := range unisonTypeMap {
			if guest.LeadSkill.Name == lskill {
				// IsViolated when song attribute does not match unison's requirement
				return song.Attribute != attr
			}
		}
		return false
	},
}

var bothLeadSkillIsActive = gameConfigLogic{
	Name: "bothLeadSkillIsActive",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		attributes := []enum.Attribute{}
		skills := []enum.SkillType{}
		for _, ocard := range team.Ocards {
			attributes = append(attributes, ocard.Card.Idol.Attribute)
			skills = append(skills, ocard.Card.Skill.SkillType.Name)
		}
		attributes = append(attributes, guest.Card.Idol.Attribute)
		skills = append(skills, guest.Card.Skill.SkillType.Name)
		// fmt.Println(attributes, team.Leader().LeadSkill.IsActive(attributes, skills), guest.LeadSkill.IsActive(attributes, skills))
		// fmt.Println(team.Leader().LeadSkill.Name, guest.LeadSkill.Name)
		// IsViolated when any of lead skill is not active
		return !(team.Leader().LeadSkill.IsActive(attributes, skills) && guest.LeadSkill.IsActive(attributes, skills))
	},
}

var guestPrincessUnisonOnUnicolor = gameConfigLogic{
	Name: "guestPrincessUnisonOnUnicolor",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		maps := []map[enum.LeadSkill]enum.Attribute{unisonTypeMap, princessTypeMap}
		for _, item := range maps {
			for lskill, attr := range item {
				if guest.LeadSkill.Name == lskill {
					for _, ocard := range team.Ocards {
						if ocard.Card.Idol.Attribute != attr {
							// IsViolated when any card does not match the princess/unison's attr

							return true
						}
					}
					return false
				}
			}
		}
		return false
	},
}

var princessWhenShouldBeUnison = gameConfigLogic{
	Name: "princessWhenShouldBeUnison",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		for lskill, attr := range princessTypeMap {
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

var guestTriColorCorrectStat = gameConfigLogic{
	Name: "guestTriColorCorrectStat",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		for lskill, stat := range tricolorStatMap {
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

var guestResonantCorrectStat = gameConfigLogic{
	Name: "guestResonantCorrectStat",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		for stat, lskill := range enum.LeadSkillResonantMap {
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

var guestPrincessUnisonCorrectStat = gameConfigLogic{
	Name: "guestPrincessUnisonCorrectStat",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		guestIsPrincessOrUnison := false
		maps := []map[enum.LeadSkill]enum.Attribute{unisonTypeMap, princessTypeMap}
		for _, item := range maps {
			for lskill := range item {
				if guest.LeadSkill.Name == lskill {
					guestIsPrincessOrUnison = true
					break
				}
			}
			if guestIsPrincessOrUnison {
				break
			}
		}
		if !guestIsPrincessOrUnison {
			return false
		}
		for stat, lskillList := range statUpLeadSkillMap {
			lskills := lskillList
			resoLeadSkill, _ := resoLeadSkillMap[stat]
			lskills = append(lskills, resoLeadSkill)
			for _, lskill := range lskills {
				if team.Leader().LeadSkill.Name == lskill {
					da, vo, vi := guest.Dance, guest.Vocal, guest.Visual
					if stat == enum.StatDance && da >= vo && da >= vi {
						return false
					} else if stat == enum.StatVisual && vi >= vo && vi >= da {
						return false
					} else if stat == enum.StatVocal && vo >= vi && vo >= da {
						return false
					}
					return true
				}
			}
		}
		// If guest is princess AND lead is not statUp skills, choose one princess only (DANCE)
		da, vo, vi := guest.Dance, guest.Vocal, guest.Visual
		if da >= vo && da >= vi {
			return false
		}
		return true
	},
}

var doNotUseReso = gameConfigLogic{
	Name: "doNotUseReso",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		lskills := []*models.LeadSkill{
			team.Leader().LeadSkill,
			guest.LeadSkill,
		}
		for _, resoLSkill := range resoLeadSkillMap {
			for _, lskill := range lskills {
				if lskill.Name == resoLSkill {
					return true
				}
			}
		}
		return false
	},
}

var tricolorCorrectColor = gameConfigLogic{
	Name: "tricolorCorrectColor",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		// if team is one color -> always false
		guestIsTriColor := false
		for lskill := range tricolorStatMap {
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
			violated := guestAttr != leftover
			if violated {
				logger.Logf("Violated: %s", team)
				logger.Logf("%s %s", team.Leader().LeadSkill.Name, guest.LeadSkill.Name)
			}
			return violated
		}
		// IsViolated when team is already 3 color AND guest is tricolor AND guest attr is not CUTE (because attr should not matter)
		if attrCount == 3 {
			return guestAttr != enum.AttrCute
		}
		// Else not violated
		return false
	},
}

var skillsAreActive = gameConfigLogic{
	Name: "skillsAreActive",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		attributes := [6]enum.Attribute{}
		for i, ocard := range team.Ocards {
			attributes[i] = ocard.Card.Idol.Attribute
		}
		attributes[5] = guest.Card.Idol.Attribute
		for _, ocard := range team.Ocards {
			if !ocard.Card.Skill.SkillType.IsActive(attributes[:]) {
				return true
			}
		}
		return false
	},
}

var logics = []gameConfigLogic{
	unisonInCorrectSongType,
	bothLeadSkillIsActive,
	guestPrincessUnisonOnUnicolor,
	princessWhenShouldBeUnison,
	guestTriColorCorrectStat,
	guestPrincessUnisonCorrectStat,
	guestResonantCorrectStat,
	skillsAreActive,

	// handled by toggle
	// doNotUseReso,
	// tricolorCorrectColor

}

func isGameConfigOk(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
	logicsUpdated := logics
	if !helper.Features.UseReso() {
		logicsUpdated = append(logicsUpdated, doNotUseReso)
	}
	if helper.Features.AllowTwoColors() {
		logicsUpdated = append(logicsUpdated, tricolorCorrectColor)
	}
	for _, logic := range logicsUpdated {
		if logic.IsViolated(team, song, guest) {
			return false
		}
	}
	return true
}

func isGameConfigOkDebug(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) string {
	logicsUpdated := logics
	if !helper.Features.UseReso() {
		logicsUpdated = append(logicsUpdated, doNotUseReso)
	}
	if helper.Features.AllowTwoColors() {
		logicsUpdated = append(logicsUpdated, tricolorCorrectColor)
	}
	for _, logic := range logicsUpdated {
		if logic.IsViolated(team, song, guest) {
			return logic.Name
		}
	}
	return "passed!"
}
