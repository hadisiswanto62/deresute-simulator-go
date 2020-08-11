package enum

// SkillType represents a skill type's name
type SkillType string

// All rarities
const (
	SkillTypeBase          SkillType = "Base Skill"
	SkillTypeScoreBonus    SkillType = "Score Bonus"
	SkillTypeComboBonus    SkillType = "Combo Bonus"
	SkillTypeConcentration SkillType = "Concentration"
)

/*
note when adding new SkillType:
	- Make sure name = name in json file
	- Add const on enum/skills.go
	- Add to models/skill_type.go/GetSkillType()
	- Add implementation to models/skill_type_impl.go
*/
