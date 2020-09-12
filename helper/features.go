package helper

var flags = map[string]bool{
	"use-concentration": true,
	"do-simulate":       false,
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

var Features feature
