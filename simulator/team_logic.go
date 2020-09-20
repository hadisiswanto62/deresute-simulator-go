package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type teamLogic struct {
	Name        string
	IsSatisfied func(team *usermodel.Team, song *models.Song) bool
}

type ocardLogic struct {
	Name        string
	IsSatisfied func(ocard *usermodel.OwnedCard, song *models.Song) bool
}

var skillsAreActive = teamLogic{
	Name: "skillsAreActive",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		attributes := [6]enum.Attribute{}
		for i, ocard := range team.Ocards {
			attributes[i] = ocard.Card.Idol.Attribute
		}
		for _, ocard := range team.Ocards {
			if !ocard.Card.Skill.SkillType.IsActive(attributes[:]) {
				return false
			}
		}
		// IsSatisfied when all skills are implemented
		return true
	},
}

var leadSkillIsImplemented = teamLogic{
	Name: "leadSkillIsImplemented",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		// IsSatisfied when lead skill is implemented
		return helper.IsLeadSkillImplemented(team.Leader().Card.LeadSkill.Name)
	},
}

var skillIsImplemented = ocardLogic{
	Name: "skillIsImplemented",
	IsSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		// IsSatisfied when skill is implemented
		return helper.IsSkillImplemented(ocard.Card.Skill.SkillType.Name)
	},
}

var cardIsSSR = ocardLogic{
	Name: "cardIsSSR",
	IsSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		// IsSatisfied when card is SSR
		return ocard.Card.Rarity.Rarity == enum.RaritySSR
	},
}

var skillIsNotConcentration = ocardLogic{
	Name: "skillIsNotConcentration",
	IsSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		// IsSatisfied when skill is not concentration
		return ocard.Card.Skill.SkillType.Name != enum.SkillTypeConcentration
	},
}

var princessUnisonOnUnicolor = teamLogic{
	Name: "princessUnisonOnUnicolor",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		maps := []map[enum.LeadSkill]enum.Attribute{unisonTypeMap, princessTypeMap}
		for _, item := range maps {
			for lskill, attr := range item {
				if team.Leader().LeadSkill.Name == lskill {
					for _, ocard := range team.Ocards {
						if ocard.Card.Idol.Attribute != attr {
							return false
						}
					}
					// IsSatisfied all card's attribute matches the lead skill's unison/princess requirement
					return true
				}
			}
		}
		// or no unison/princess card
		return true
	},
}

var tricolorOnMinimum2Color = teamLogic{
	Name: "tricolorOnMinimum2Color",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		for lskill := range tricolorStatMap {
			if team.Leader().LeadSkill.Name == lskill {
				attributes := make(map[enum.Attribute]bool)
				for _, ocard := range team.Ocards {
					attributes[ocard.Card.Idol.Attribute] = true
				}
				// IsSatisfied when there are >2 unique attributes
				return len(attributes) >= 2
			}
		}
		// or lead skill is not tricolor
		return true
	},
}

var tricolorOnMinimum3Color = teamLogic{
	Name: "tricolorOnMinimum3Color",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		for lskill := range tricolorStatMap {
			if team.Leader().LeadSkill.Name == lskill {
				attributes := make(map[enum.Attribute]bool)
				for _, ocard := range team.Ocards {
					attributes[ocard.Card.Idol.Attribute] = true
				}
				// IsSatisfied when there are 3 unique attributes
				return len(attributes) == 3
			}
		}
		// or lead skill is not tricolor
		return true
	},
}

var coolLeadSkills = []enum.LeadSkill{
	enum.LeadSkillCoolMakeup,
	enum.LeadSkillCoolStep,
	enum.LeadSkillCoolVoice,
	enum.LeadSkillCoolAbility,
	enum.LeadSkillCoolCheer,
	enum.LeadSkillCoolPrincess,
	enum.LeadSkillCoolUnison,
	enum.LeadSkillCoolBrilliance,
	enum.LeadSkillCoolEnergy,
}

var cuteLeadSkills = []enum.LeadSkill{
	enum.LeadSkillCuteMakeup,
	enum.LeadSkillCuteStep,
	enum.LeadSkillCuteVoice,
	enum.LeadSkillCuteAbility,
	enum.LeadSkillCuteCheer,
	enum.LeadSkillCutePrincess,
	enum.LeadSkillCuteUnison,
	enum.LeadSkillCuteBrilliance,
	enum.LeadSkillCuteEnergy,
}

var passionLeadSkills = []enum.LeadSkill{
	enum.LeadSkillPassionMakeup,
	enum.LeadSkillPassionStep,
	enum.LeadSkillPassionVoice,
	enum.LeadSkillPassionAbility,
	enum.LeadSkillPassionCheer,
	enum.LeadSkillPassionPrincess,
	enum.LeadSkillPassionUnison,
	enum.LeadSkillPassionBrilliance,
	enum.LeadSkillPassionEnergy,
}

var attrSpecificLeadSkillOnUnicolor = teamLogic{
	Name: "attrSpecificLeadSkillOnUnicolor",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		skillMap := map[enum.Attribute][]enum.LeadSkill{
			enum.AttrCool:    coolLeadSkills,
			enum.AttrCute:    cuteLeadSkills,
			enum.AttrPassion: passionLeadSkills,
		}
		for attr, lskillList := range skillMap {
			for _, lskill := range lskillList {
				if team.Leader().LeadSkill.Name == lskill {
					for _, ocard := range team.Ocards {
						if ocard.Card.Idol.Attribute != attr {
							return false
						}
					}
					// IsSatisfied when all cards matches the lead skills' attr requirement
					return true
				}
			}
		}
		// or when lead skill is not attr specific
		return true
	},
}

var unicolorOnColoredSong = ocardLogic{
	Name: "unicolorOnColoredSong",
	IsSatisfied: func(ocard *usermodel.OwnedCard, song *models.Song) bool {
		if song.Attribute == enum.AttrAll {
			return true
		}
		return ocard.Card.Idol.Attribute == song.Attribute
	},
}

var noDuoColor = teamLogic{
	Name: "noDuoColor",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		attributes := make(map[enum.Attribute]bool)
		for _, ocard := range team.Ocards {
			attributes[ocard.Card.Idol.Attribute] = true
		}
		return len(attributes) != 2
	},
}
var motifWithHighCorrectStat = teamLogic{
	Name: "motifWithHighCorrectStat",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		for stat, skill := range motifStatMap {
			for _, ocard := range team.Ocards {
				if ocard.Card.Skill.SkillType.Name == skill {
					da, vo, vi := 0, 0, 0
					for _, ocard := range team.Ocards {
						da += ocard.Dance
						vo += ocard.Vocal
						vi += ocard.Visual
					}
					sum := float64(da + vo + vi)
					// IsSatisfied when the stat required by the motif is > 0.4*totalAppeal
					if stat == enum.StatDance && float64(da)/sum > 0.4 {
						return true
					} else if stat == enum.StatVocal && float64(vo)/sum > 0.4 {
						return true
					} else if stat == enum.StatVisual && float64(vi)/sum > 0.4 {
						return true
					}
					return false
				}
			}
		}
		return true
	},
}
var useUnevolvedWithoutEvolved = teamLogic{
	Name: "useUnevolvedWithoutEvolved",
	IsSatisfied: func(team *usermodel.Team, song *models.Song) bool {
		for _, ocard := range team.Ocards {
			if ocard.Card.Rarity.IsEvolved {
				continue
			}
			foundEvolved := false
			for _, ocard2 := range team.Ocards {
				if ocard2.Card.Rarity.IsEvolved {
					if ocard2.Card.ID == ocard.Card.ID+1 {
						foundEvolved = true
					}
				}
			}
			if !foundEvolved {
				return false
			}
		}
		return true
	},
}
var teamLogics = []teamLogic{
	leadSkillIsImplemented,
	princessUnisonOnUnicolor,
	tricolorOnMinimum3Color,
	skillsAreActive,
	attrSpecificLeadSkillOnUnicolor,
	noDuoColor,
	motifWithHighCorrectStat,
	useUnevolvedWithoutEvolved,
	// tricolorOnMinimum2Color,

	// for info only:
	// skillIsNotConcentration,
}

var ocardLogics = []ocardLogic{
	cardIsSSR,
	skillIsImplemented,
	unicolorOnColoredSong,
}

func isTeamOk(team *usermodel.Team, song *models.Song) bool {
	for _, logic := range teamLogics {
		if !logic.IsSatisfied(team, song) {
			return false
		}
	}

	ocardLogic := ocardLogics
	if !helper.Features.UseConcentration() {
		ocardLogic = append(ocardLogic, skillIsNotConcentration)
	}
	for _, ocard := range team.Ocards {
		for _, logic := range ocardLogic {
			if !logic.IsSatisfied(ocard, song) {
				return false
			}
		}
	}
	return true
}

func isTeamOkDebug(team *usermodel.Team, song *models.Song) string {
	for _, logic := range teamLogics {
		if !logic.IsSatisfied(team, song) {
			return logic.Name
		}
	}
	for _, ocard := range team.Ocards {
		for _, logic := range ocardLogics {
			if !logic.IsSatisfied(ocard, song) {
				return logic.Name
			}
		}
	}
	return "team looks ok"
}
