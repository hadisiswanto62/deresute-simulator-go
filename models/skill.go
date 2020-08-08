package models

// Skill represents a card's skill, complete with the duration, proc chance, and effect length.
type Skill struct {
	ID             int
	Timer          int
	ProcChance     [2]uint16
	EffectLength   [2]uint16
	SkillType      *SkillType
	ActivationCost int
}
