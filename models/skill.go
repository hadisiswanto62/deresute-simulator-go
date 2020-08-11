package models

// Skill represents a card's skill, complete with the duration, proc chance, and effect length.
type Skill struct {
	ID int
	// Timer is every how many `second` the skill rolls.
	Timer int
	// ProcChance is the min and max of skill's proc chance in `ProcChance/10000` (e.g. ProcChance 40000 would be 0.4 chance)
	ProcChance [2]int
	// EffectLength is the duration of the skill if it activates in `centisecond`
	EffectLength   [2]int
	SkillType      *SkillType
	ActivationCost int
}
