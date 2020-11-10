package enum

// SkillType represents a skill type's name
type SkillType string

// All rarities
const (
	SkillTypeBase             SkillType = "Base Skill"
	SkillTypeScoreBonus       SkillType = "Score Bonus"
	SkillTypeComboBonus       SkillType = "Combo Bonus"
	SkillTypeConcentration    SkillType = "Concentration"
	SkillTypeHealer           SkillType = "Healer"
	SkillTypeAllRound         SkillType = "All-Round"
	SkillTypeLifeSparkle      SkillType = "Life Sparkle"
	SkillTypeTricolorSynergy  SkillType = "Tricolor Synergy"
	SkillTypeCoordinate       SkillType = "Coordinate"
	SkillTypeOverload         SkillType = "Overload"
	SkillTypeCuteFocus        SkillType = "Cute Focus"
	SkillTypeCoolFocus        SkillType = "Cool Focus"
	SkillTypePassionFocus     SkillType = "Passion Focus"
	SkillTypeTuning           SkillType = "Tuning"
	SkillTypeDanceMotif       SkillType = "Dance Motif"
	SkillTypeVisualMotif      SkillType = "Visual Motif"
	SkillTypeVocalMotif       SkillType = "Vocal Motif"
	SkillTypePerfectLock      SkillType = "Perfect Lock"
	SkillTypeComboGuard       SkillType = "Combo Guard"
	SkillTypeEncore           SkillType = "Encore"
	SkillTypeAlternate        SkillType = "Alternate"
	SkillTypeSkillBoost       SkillType = "Skill Boost"
	SkillTypeTricolorSymphony SkillType = "Tricolor Symphony"
	SkillTypeLifeGuard        SkillType = "Life Guard"
	SkillTypeCuteEnsemble     SkillType = "Cute Ensemble"
	SkillTypeCoolEnsemble     SkillType = "Cool Ensemble"
	SkillTypePassionEnsemble  SkillType = "Passion Ensemble"
	SkillTypeFlickAct         SkillType = "Flick Act"
	SkillTypeHoldAct          SkillType = "Long Act"
	SkillTypeSlideAct         SkillType = "Slide Act"
)

/*
note when adding new SkillType:
	- Make sure name = name in json file
	- Add const on enum/skills.go
	- Add to models/skill_type.go/GetSkillType()
	- Add implementation to models/skill_type_impl.go
*/
var MotifStatMap = map[Stat]SkillType{
	StatDance:  SkillTypeDanceMotif,
	StatVocal:  SkillTypeVocalMotif,
	StatVisual: SkillTypeVisualMotif,
}
