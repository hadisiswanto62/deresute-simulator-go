package simulatormodels

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// Playable is an interface for game configs that can be played.
// for implementers: make sure all O(1)
type Playable interface {
	GetSkillActivableCards() []*usermodel.OwnedCard
	GetLeadSkillActivableCards() []*usermodel.OwnedCard
	GetSong() *models.Song
	GetBaseVisual() int
	GetBaseVocal() int
	GetBaseDance() int
	GetAppeal() int
	GetHp() int

	GetTeamAttributesv2() []enum.Attribute
	// note that teamSkills != skills from getSkillActivableCards
	// teamSkills is used for reso (which include guest's skill that cant be active)
	GetTeamSkillsv2() []enum.SkillType
	IsResonantActive() bool
	GetCards() []*usermodel.OwnedCard
	GetLeaderIndex() int
	GetGuest() *usermodel.OwnedCard
}

var _ Playable = GameConfig{}
