package models

import "fmt"

// Card is a card
type Card struct {
	ID        int
	SeriesID  int   `json:"series_id"`
	Idol      *Idol `json:"chara"`
	Rarity    *Rarity
	LeadSkill *LeadSkill
	Skill     *Skill

	BonusDance  int `json:"bonus_dance"`
	BonusHp     int `json:"bonus_hp"`
	BonusVisual int `json:"bonus_visual"`
	BonusVocal  int `json:"bonus_vocal"`
	DanceMax    int `json:"dance_max"`
	DanceMin    int `json:"dance_min"`
	HpMax       int `json:"hp_max"`
	HpMin       int `json:"hp_min"`
	VisualMax   int `json:"visual_max"`
	VisualMin   int `json:"visual_min"`
	VocalMax    int `json:"vocal_max"`
	VocalMin    int `json:"vocal_min"`
}

func (c Card) String() string {
	return fmt.Sprintf("<Card %d (%s); %s; %s; %s>",
		c.ID, c.Idol.Name, c.Rarity.Rarity, c.LeadSkill.Name, c.Skill.SkillType.Name)
}
