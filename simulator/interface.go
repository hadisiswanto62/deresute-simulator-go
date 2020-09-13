package simulator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// Interface for game configs that can be played.
// for implementers: make sure all O(1)
type Playable interface {
	getSkillActivableCards() []*usermodel.OwnedCard
	getLeadSkillActivableCards() []*usermodel.OwnedCard
	getSong() *models.Song
	getBaseVisual() int
	getBaseVocal() int
	getBaseDance() int
	getAppeal() int
	getHp() int

	getTeamAttributesv2() []enum.Attribute
	// note that teamSkills != skills from getSkillActivableCards
	// teamSkills is used for reso (which include guest's skill that cant be active)
	getTeamSkillsv2() []enum.SkillType
	isResonantActive() bool
	getCards() []*usermodel.OwnedCard
	getLeaderIndex() int
	getGuest() *usermodel.OwnedCard
}

var _ Playable = GameConfig{}
var _ Playable = GameConfig2{}
