package helper

import (
	"testing"
)

func TestUseConcentration(t *testing.T) {
	flags["use-concentration"] = true
	use := Features.UseConcentration()
	if !use {
		t.Errorf("not activated!")
	}
	flags["use-concentration"] = false
	use = Features.UseConcentration()
	if use {
		t.Errorf("activated!")
	}
	delete(flags, "use-concentration")
	use = Features.UseConcentration()
	if use {
		t.Errorf("activated!")
	}
}
