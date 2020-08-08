package jsonmodels

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// TmpIdol is temporary struct for Idol
type TmpIdol struct {
	ID        int            `json:"chara_id"`
	Name      string         `json:"conventional"`
	Attribute enum.Attribute `json:"type"`
}

// TmpLeadSkill is temporary struct for LeadSkill
type TmpLeadSkill struct {
	// ID          int
	Name string
	// Description string `json:"explain_en"`
}

// TmpSkill is temporary struct for Skill
type TmpSkill struct {
	ID                int
	Condition         int
	EffectLength      [2]int `json:"effect_length"`
	Description       string `json:"explain_en"`
	ProcChance        [2]int `json:"proc_chance"`
	SkillTriggerType  int    `json:"skill_trigger_type"`
	SkillTriggerValue int    `json:"skill_trigger_value"`
	SkillType         string `json:"skill_type"`
	Value             int
}

// TmpRarity is temporary struct for Rarity
type TmpRarity struct {
	Rarity       uint8
	BaseMaxLevel uint8 `json:"base_max_level"`
}

// IsEvolved returns if the rarity belongs to an evolved card or not.
func (r TmpRarity) IsEvolved() bool {
	return r.Rarity%2 == 0
}

// ToEnum converts the object to enum.Rarity
func (r TmpRarity) ToEnum() enum.Rarity {
	switch r.Rarity {
	case 1, 2:
		return enum.RarityN
	case 3, 4:
		return enum.RarityR
	case 5, 6:
		return enum.RaritySR
	case 7, 8:
		return enum.RaritySSR
	}
	return enum.RarityN
}

// TmpCard is a temporary struct for Card
type TmpCard struct {
	ID           int
	SeriesID     int `json:"series_id"`
	Title        string
	TmpIdol      `json:"chara"`
	TmpLeadSkill `json:"lead_skill"`
	TmpSkill     `json:"skill"`
	TmpRarity    `json:"rarity"`

	BonusDance  uint16 `json:"bonus_dance"`
	BonusHp     uint16 `json:"bonus_hp"`
	BonusVisual uint16 `json:"bonus_visual"`
	BonusVocal  uint16 `json:"bonus_vocal"`
	DanceMax    uint16 `json:"dance_max"`
	DanceMin    uint16 `json:"dance_min"`
	HpMax       uint16 `json:"hp_max"`
	HpMin       uint16 `json:"hp_min"`
	VisualMax   uint16 `json:"visual_max"`
	VisualMin   uint16 `json:"visual_min"`
	VocalMax    uint16 `json:"vocal_max"`
	VocalMin    uint16 `json:"vocal_min"`
}
