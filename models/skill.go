package models

// Skill represents a card's skill, complete with the duration, proc chance, and effect length.
type Skill struct {
	ID             int
	Duration       int
	ProcChance     [2]int
	EffectLength   [2]int
	SkillType      *SkillType
	ActivationCost int
}
