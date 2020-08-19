package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type gameConfigLogic struct {
	Name       string
	IsViolated func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool
}

var motifStatMap = map[enum.Stat]enum.SkillType{
	enum.StatDance:  enum.SkillTypeDanceMotif,
	enum.StatVocal:  enum.SkillTypeVocalMotif,
	enum.StatVisual: enum.SkillTypeVisualMotif,
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

var statUpLeadSkillMap = map[enum.Stat][3]enum.LeadSkill{
	enum.StatDance: [3]enum.LeadSkill{
		enum.LeadSkillCuteStep,
		enum.LeadSkillCoolStep,
		enum.LeadSkillPassionStep,
	},
	enum.StatVocal: [3]enum.LeadSkill{
		enum.LeadSkillCuteVoice,
		enum.LeadSkillCoolVoice,
		enum.LeadSkillPassionVoice,
	},
	enum.StatVisual: [3]enum.LeadSkill{
		enum.LeadSkillCuteMakeup,
		enum.LeadSkillCoolMakeup,
		enum.LeadSkillPassionMakeup,
	},
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
		attributes := [6]enum.Attribute{}
		for i, ocard := range team.Ocards {
			attributes[i] = ocard.Card.Idol.Attribute
		}
		attributes[5] = guest.Card.Idol.Attribute
		// IsViolated when any of lead skill is not active
		return !(team.Leader().LeadSkill.IsActive(attributes) && guest.LeadSkill.IsActive(attributes))
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

var guestPrincessUnisonCorrectStat = gameConfigLogic{
	Name: "guestPrincessUnisonCorrectStat",
	IsViolated: func(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
		guestIsPrincessOrUnison := false
		maps := []map[enum.LeadSkill]enum.Attribute{unisonTypeMap, princessTypeMap}
		for _, item := range maps {
			for lskill, _ := range item {
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
			for _, lskill := range lskillList {
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

var logics = []gameConfigLogic{
	unisonInCorrectSongType,
	bothLeadSkillIsActive,
	guestPrincessUnisonOnUnicolor,
	princessWhenShouldBeUnison,
	guestTriColorCorrectStat,
	guestPrincessUnisonCorrectStat,
}

func isGameConfigOk(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) bool {
	for _, logic := range logics {
		if logic.IsViolated(team, song, guest) {
			if logic.Name != "guestPrincessUnisonCorrectStat" {
				return false
			}
			return false
		}
	}
	return true
}

func isGameConfigOkDebug(team *usermodel.Team, song *models.Song, guest *usermodel.OwnedCard) string {
	for _, logic := range logics {
		if logic.IsViolated(team, song, guest) {
			return logic.Name
		}
	}
	return ""
}
