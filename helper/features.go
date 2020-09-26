package helper

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

var flags = map[string]bool{
	// should be TRUE for any valid sims
	"use-concentration": true,
	"do-simulate":       true,
	"limit-appeals":     true,
	"limit-score":       true, //true this only when optimizing
	"use-game-fast":     false,

	// should be FALSE for any valid sims
	"always-good-rolls": true,

	// false for now
	"use-reso": false,
}

type feature struct{}

func checkFlag(flag string) bool {
	val, _ := flags[flag]
	return val == true
}

func (f feature) UseConcentration() bool {
	return checkFlag("use-concentration")
}

func (f feature) ReallySimulate() bool {
	return checkFlag("do-simulate")
}

func (f feature) UseReso() bool {
	return checkFlag("use-reso")
}

func (f feature) LimitAppeals() bool {
	return checkFlag("limit-appeals")
}

func (f feature) LimitScore() bool {
	return checkFlag("limit-score")
}

func (f feature) AlwaysGoodRolls() bool {
	return checkFlag("always-good-rolls")
}

func (f feature) GetScoreLimitForAttr(attr enum.Attribute) int {
	limit, ok := scoreLimit[attr]
	if !ok {
		limit = 1000000
	}
	return limit
}

func (f feature) UseFastGame() bool {
	return checkFlag("use-game-fast")
}

var scoreLimit = map[enum.Attribute]int{
	enum.AttrAll:     1250000,
	enum.AttrCool:    1150000,
	enum.AttrPassion: 1100000,
	enum.AttrCute:    1000000,
}

var Features feature
