package statcalculator

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type IStatCalculator interface {
	SetCards([]*usermodel.OwnedCard)
	SetLeaderIndex(int)
	SetGuest(*usermodel.OwnedCard)
	SetSupports([]*usermodel.OwnedCard)
	SetSong(*models.Song)
	Calculate(bonusAppeal int) (*GameConfigStats, error)
}

type GameConfigStats struct {
	Appeal         int
	Hp             int
	TeamAttributes []enum.Attribute
	TeamSkills     []enum.SkillType
	resonantStat   enum.Stat
}

func (stat GameConfigStats) IsResonantOn() bool {
	return stat.resonantStat != ""
}

type StatCalculatorType int

const (
	NormalStatCalculator StatCalculatorType = iota
)

func CalculatorDispatcher(type_ StatCalculatorType) IStatCalculator {
	switch type_ {
	case NormalStatCalculator:
		return &normalStatCalculator{}
	}
	return &normalStatCalculator{}
}
