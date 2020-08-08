package jsonmodels

import (
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
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
		if !helper.DebugMode {
			t.Errorf("error should not be nil: %v", err)
		}
	}
}

func TestMakeLeadSkillSameInstance(t *testing.T) {
	lskill, _ := makeLeadSkill(sampleLeadSkill)
	lskill2, _ := makeLeadSkill(sampleLeadSkill)
	if lskill != lskill2 {
		t.Errorf("makeLeadSkill returns new object (addr = %p and %p)", lskill, lskill2)
	}
}

func TestMakeRarity(t *testing.T) {
	r := makeRarity(TmpRarity{
		Rarity:       6,
		BaseMaxLevel: 70,
	})
	if r == nil {
		t.Errorf("Cannot make rarity")
	}
}

func TestMakeIdol(t *testing.T) {
	i := makeIdol(TmpIdol{
		ID:        1,
		Name:      "Sample Idol",
		Attribute: enum.AttrCool,
	})
	if i == nil {
		t.Errorf("Cannot make idol")
	}
}

func TestMakeSkill(t *testing.T) {
	x := TmpSkill{
		ID:           300591,
		Condition:    11,
		EffectLength: [2]int{500, 750},
		ProcChance:   [2]int{4000, 6000},
		SkillType:    "Perfect Score Bonus",
	}
	s, err := makeSkill(x)
	if err != nil {
		t.Errorf("Cannot make skill: %v", err)
	}
	if s == nil {
		t.Errorf("Cannot make skill!")
	}
}
