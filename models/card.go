package models

// Card is a card
type Card struct {
	ID       int
	SeriesID int `json:"series_id"`
	Idol     `json:"chara"`

	BonusDance   uint16 `json:"bonus_dance"`
	BonusHp      uint16 `json:"bonus_hp"`
	BonusVisual  uint16 `json:"bonus_visual"`
	BonusVocal   uint16 `json:"bonus_vocal"`
	DanceMax     uint16 `json:"dance_max"`
	DanceMin     uint16 `json:"dance_min"`
	HpMax        uint16 `json:"hp_max"`
	HpMin        uint16 `json:"hp_min"`
	OverallMax   uint16 `json:"overall_max"`
	OverallMin   uint16 `json:"overall_min"`
	OverallBonus uint16 `json:"overall_bonus"`
	VisualMax    uint16 `json:"visual_max"`
	VisualMin    uint16 `json:"visual_min"`
	VocalMax     uint16 `json:"vocal_max"`
	VocalMin     uint16 `json:"vocal_min"`
}

// missing fields:
// lead_skill, rarity, skill
