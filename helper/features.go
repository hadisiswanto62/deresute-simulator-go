package helper

import "github.com/hadisiswanto62/deresute-simulator-go/enum"

var flags = map[string]bool{
	// should be TRUE for any valid sims
	"use-concentration": true,
	"do-simulate":       true,
	"limit-appeals":     true,
	"limit-score":       true, //true this only when optimizing
	"use-game-fast":     true,
	"allow-two-colors":  true,

	// should be FALSE for any valid sims
	"always-good-rolls": false,

	// false for now
	"use-reso": false,

	// depends
	"use-window-abuse": false,

	// for debug (if not debug -> false)
	"debug-no-logic": false,
}

type feature struct{}

func checkFlag(flag string) bool {
	val, _ := flags[flag]
	return val == true
}

func (f feature) DebugNoLogic() bool {
	return checkFlag("debug-no-logic")
}

func (f feature) AllowTwoColors() bool {
	return checkFlag("allow-two-colors")
}

func (f feature) UseConcentration() bool {
	return checkFlag("use-concentration")
}

func (f feature) UseWindowAbuse() bool {
	return checkFlag("use-window-abuse")
}

func (f feature) SetWindowAbuse(val bool) {
	flags["use-window-abuse"] = val
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

func (f feature) SetLimitScore(val bool) {
	flags["limit-score"] = val

}

func (f feature) AlwaysGoodRolls() bool {
	return checkFlag("always-good-rolls")
}

func (f feature) GetScoreLimitForAttr(attr enum.Attribute, level int) int {
	multiplier := GetSongDifficultyMultiplier(level)
	limit, ok := scoreLimit[attr]
	if !ok {
		limit = 500000
	}
	return int(float64(limit) * multiplier)
}

func (f feature) UseFastGame() bool {
	return checkFlag("use-game-fast")
}

var scoreLimit = map[enum.Attribute]int{
	enum.AttrAll:     650000,
	enum.AttrCool:    575000,
	enum.AttrPassion: 550000,
	enum.AttrCute:    500000,
}

var Features feature
