package enum

// LeadSkill represents a lead skill's name
type LeadSkill string

// All rarities
const (
	LeadSkillBase       LeadSkill = "Base Lead Skill"
	LeadSkillIrrelevant LeadSkill = "Irrelevant Lead Skill"

	LeadSkillCuteMakeup        LeadSkill = "キュートメイク"
	LeadSkillCuteStep          LeadSkill = "キュートステップ"
	LeadSkillCuteVoice         LeadSkill = "キュートボイス"
	LeadSkillCuteAbility       LeadSkill = "キュートアビリティ"
	LeadSkillCuteCheer         LeadSkill = "キュートチアー"
	LeadSkillCutePrincess      LeadSkill = "キュートプリンセス"
	LeadSkillCuteUnison        LeadSkill = "キュートユニゾン"
	LeadSkillCuteBrilliance    LeadSkill = "キュートブリリアンス"
	LeadSkillCuteEnergy        LeadSkill = "キュートエナジー"
	LeadSkillPassionMakeup     LeadSkill = "パッションメイク"
	LeadSkillPassionStep       LeadSkill = "パッションステップ"
	LeadSkillPassionVoice      LeadSkill = "パッションボイス"
	LeadSkillPassionAbility    LeadSkill = "パッションアビリティ"
	LeadSkillPassionCheer      LeadSkill = "パッションチアー"
	LeadSkillPassionPrincess   LeadSkill = "パッションプリンセス"
	LeadSkillPassionUnison     LeadSkill = "パッションユニゾン"
	LeadSkillPassionBrilliance LeadSkill = "パッションブリリアンス"
	LeadSkillPassionEnergy     LeadSkill = "パッションエナジー"
	LeadSkillCoolMakeup        LeadSkill = "クールメイク"
	LeadSkillCoolStep          LeadSkill = "クールステップ"
	LeadSkillCoolVoice         LeadSkill = "クールボイス"
	LeadSkillCoolAbility       LeadSkill = "クールアビリティ"
	LeadSkillCoolCheer         LeadSkill = "クールチアー"
	LeadSkillCoolPrincess      LeadSkill = "クールプリンセス"
	LeadSkillCoolUnison        LeadSkill = "クールユニゾン"
	LeadSkillCoolBrilliance    LeadSkill = "クールブリリアンス"
	LeadSkillCoolEnergy        LeadSkill = "クールエナジー"
	LeadSkillTricolorMakeup    LeadSkill = "トリコロール・メイク"
	LeadSkillTricolorVoice     LeadSkill = "トリコロール・ボイス"
	LeadSkillTricolorStep      LeadSkill = "トリコロール・ステップ"
	LeadSkillTricolorAbility   LeadSkill = "トリコロール・アビリティ"
	LeadSkillShinyStep         LeadSkill = "シャイニーステップ"
	LeadSkillShinyVoice        LeadSkill = "シャイニーボイス"
	LeadSkillShinyMakeup       LeadSkill = "シャイニーメイク"

	LeadSkillCuteCrossPassion LeadSkill = "キュート・クロス・パッション"
	LeadSkillCuteCrossCool    LeadSkill = "キュート・クロス・クール"

	LeadSkillCoolCrossCute    LeadSkill = "クール・クロス・キュート"
	LeadSkillCoolCrossPassion LeadSkill = "クール・クロス・パッション"

	LeadSkillPassionCrossCute LeadSkill = "パッション・クロス・キュート"
	LeadSkillPassionCrossCool LeadSkill = "パッション・クロス・クール"

	LeadSkillResonantStep   LeadSkill = "レゾナンス・ステップ"
	LeadSkillResonantVoice  LeadSkill = "レゾナンス・ボイス"
	LeadSkillResonantMakeup LeadSkill = "レゾナンス・メイク"
)

var ResonantMap = map[Stat]LeadSkill{
	StatDance:  LeadSkillResonantStep,
	StatVisual: LeadSkillResonantMakeup,
	StatVocal:  LeadSkillResonantVoice,
}
