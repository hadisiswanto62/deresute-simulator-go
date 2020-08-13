package csvmodels

type TmpOwnedCardRawData struct {
	CardID     int `csv:"card_id"`
	SkillLevel int `csv:"skill_level"`
	StarRank   int `csv:"star_rank"`
	PotVisual  int `csv:"pot_visual"`
	PotDance   int `csv:"pot_dance"`
	PotVocal   int `csv:"pot_vocal"`
	PotHp      int `csv:"pot_hp"`
	PotSkill   int `csv:"pot_skill"`
}
