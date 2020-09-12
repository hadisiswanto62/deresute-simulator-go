package helper

var flags = map[string]bool{
	"use-concentration": true,
	"do-simulate":       false,
	"use-reso":          true,
	"limit-appeals":     true,
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

var Features feature
