package models

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/config"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// LeadSkill represents a leader skill. (Do not create new instance. Use GetLeadSkill(name) instead)
type LeadSkill struct {
	Name           enum.LeadSkill
	IsActive       func(attr [6]enum.Attribute, skills [6]enum.SkillType) bool
	StatBonus      func(rarity enum.Rarity, cardAttr enum.Attribute, stat enum.Stat, songAttr enum.Attribute) float64
	SkillProbBonus func(rarity enum.Rarity, cardAttr enum.Attribute) float64
	HpBonus        func(rarity enum.Rarity, cardAttr enum.Attribute) float64
}

var leadSkillMap = map[string]*LeadSkill{
	string(enum.LeadSkillCuteMakeup):    &LeadSkillCuteMakeup,
	string(enum.LeadSkillPassionMakeup): &LeadSkillPassionMakeup,
	string(enum.LeadSkillCoolMakeup):    &LeadSkillCoolMakeup,

	string(enum.LeadSkillCuteStep):    &LeadSkillCuteStep,
	string(enum.LeadSkillPassionStep): &LeadSkillPassionStep,
	string(enum.LeadSkillCoolStep):    &LeadSkillCoolStep,

	string(enum.LeadSkillCuteVoice):    &LeadSkillCuteVoice,
	string(enum.LeadSkillPassionVoice): &LeadSkillPassionVoice,
	string(enum.LeadSkillCoolVoice):    &LeadSkillCoolVoice,

	string(enum.LeadSkillCuteAbility):    &LeadSkillCuteAbility,
	string(enum.LeadSkillPassionAbility): &LeadSkillPassionAbility,
	string(enum.LeadSkillCoolAbility):    &LeadSkillCoolAbility,

	string(enum.LeadSkillCuteCheer):    &LeadSkillCuteCheer,
	string(enum.LeadSkillPassionCheer): &LeadSkillPassionCheer,
	string(enum.LeadSkillCoolCheer):    &LeadSkillCoolCheer,

	string(enum.LeadSkillCutePrincess):    &LeadSkillCutePrincess,
	string(enum.LeadSkillPassionPrincess): &LeadSkillPassionPrincess,
	string(enum.LeadSkillCoolPrincess):    &LeadSkillCoolPrincess,

	string(enum.LeadSkillCuteUnison):    &LeadSkillCuteUnison,
	string(enum.LeadSkillPassionUnison): &LeadSkillPassionUnison,
	string(enum.LeadSkillCoolUnison):    &LeadSkillCoolUnison,

	string(enum.LeadSkillCuteBrilliance):    &LeadSkillCuteBrilliance,
	string(enum.LeadSkillPassionBrilliance): &LeadSkillPassionBrilliance,
	string(enum.LeadSkillCoolBrilliance):    &LeadSkillCoolBrilliance,

	string(enum.LeadSkillCuteEnergy):    &LeadSkillCuteEnergy,
	string(enum.LeadSkillPassionEnergy): &LeadSkillPassionEnergy,
	string(enum.LeadSkillCoolEnergy):    &LeadSkillCoolEnergy,

	string(enum.LeadSkillTricolorMakeup):  &LeadSkillTricolorMakeup,
	string(enum.LeadSkillTricolorStep):    &LeadSkillTricolorStep,
	string(enum.LeadSkillTricolorVoice):   &LeadSkillTricolorVoice,
	string(enum.LeadSkillTricolorAbility): &LeadSkillTricolorAbility,

	string(enum.LeadSkillShinyMakeup): &LeadSkillShinyMakeup,
	string(enum.LeadSkillShinyStep):   &LeadSkillShinyStep,
	string(enum.LeadSkillShinyVoice):  &LeadSkillShinyVoice,

	string(enum.LeadSkillCuteCrossCool):    &LeadSkillCuteCrossCool,
	string(enum.LeadSkillCuteCrossPassion): &LeadSkillCuteCrossPassion,

	string(enum.LeadSkillCoolCrossCute):    &LeadSkillCoolCrossCute,
	string(enum.LeadSkillCoolCrossPassion): &LeadSkillCoolCrossPassion,

	string(enum.LeadSkillPassionCrossCool): &LeadSkillPassionCrossCool,
	string(enum.LeadSkillPassionCrossCute): &LeadSkillPassionCrossCute,

	string(enum.LeadSkillResonantMakeup): &LeadSkillResonantMakeup,
	string(enum.LeadSkillResonantStep):   &LeadSkillResonantStep,
	string(enum.LeadSkillResonantVoice):  &LeadSkillResonantVoice,
}

// GetLeadSkill returns pointer to lead skill with the requested name
func GetLeadSkill(name string) (*LeadSkill, error) {
	skill, ok := leadSkillMap[name]
	if ok {
		return skill, nil
	}

	var irrelevants = [5]string{
		"フォーチュンプレゼント",
		"シンデレラチャーム",
		"シンデレラエール",
		"クリスマスプレゼント",
		"ワールドレベル",
	}
	for _, lskillName := range irrelevants {
		if name == lskillName {
			return &LeadSkillIrrelevant, nil
		}
	}

	if config.DebugMode {
		if name != "" {
			fmt.Println(name)
		}
		return &LeadSkillBase, nil
	}
	err := fmt.Errorf("invalid skill name: %s", name)
	return &LeadSkillBase, err
}

// GetLeadSkillFromEnum returns pointer to lead skill with the requested name
// (but in enum.LeadSkill)
func GetLeadSkillFromEnum(name enum.LeadSkill) (*LeadSkill, error) {
	return GetLeadSkill(string(name))
}
