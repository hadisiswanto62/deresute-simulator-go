package jsonmodels

import (
	"testing"
)

var sampleLeadSkill = TmpLeadSkill{
	Name: "キュートメイク",
}

var sampleInvalidLeadSkill = TmpLeadSkill{
	Name: "TESTSETTE",
}

func TestMakeLeadSkill(t *testing.T) {
	_, err := makeLeadSkill(sampleLeadSkill)
	if err != nil {
		t.Errorf("test failed: %v", err)
	}
	_, err = makeLeadSkill(sampleInvalidLeadSkill)
	if err == nil {
		t.Errorf("error should not be nil: %v", err)
	}
}

func TestMakeLeadSkillSameInstance(t *testing.T) {
	lskill, _ := makeLeadSkill(sampleLeadSkill)
	lskill2, _ := makeLeadSkill(sampleLeadSkill)
	if lskill != lskill2 {
		t.Errorf("makeLeadSkill returns new object (addr = %p and %p)", lskill, lskill2)
	}
}
