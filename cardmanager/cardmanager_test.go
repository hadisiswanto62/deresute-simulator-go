package cardmanager

import (
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func init() {
	os.Chdir("../")
}

func TestCreateDefault(t *testing.T) {
	cm, err := Default()
	if err != nil {
		t.Errorf("Failed to parse cards. %v", err)
	}
	if len(cm.Cards) == 0 {
		t.Errorf("Failed to parse cards.")
	}
}
func TestCardFields(t *testing.T) {
	cm, _ := Default()
	card := cm.Filter().ID(100001).First()
	if have, want := card.ID, 100001; want != have {
		t.Errorf("Error on fields ID! want = %v have = %v", want, have)
	}
	if have, want := card.SeriesID, 100001; want != have {
		t.Errorf("Error on fields SeriesID! want = %v have = %v", want, have)
	}
	if have, want := card.Idol.ID, 101; want != have {
		t.Errorf("Error on fields Idol ID! want = %v have = %v", want, have)
	}
	if have, want := card.Idol.Attribute, enum.AttrCute; want != have {
		t.Errorf("Error on fields Idol Attr! want = %v have = %v", want, have)
	}
	if have, want := card.Idol.Name, "Shimamura Uzuki"; want != have {
		t.Errorf("Error on fields Idol name! want = %v have = %v", want, have)
	}
	if have, want := card.Rarity.ID, 3; want != have {
		t.Errorf("Error on fields Rarity ID! want = %v have = %v", want, have)
	}
	if have, want := card.Rarity.Rarity, enum.RarityR; want != have {
		t.Errorf("Error on fields Rarity! want = %v have = %v", want, have)
	}
	if have, want := card.Rarity.IsEvolved, false; want != have {
		t.Errorf("Error on fields IsEvolved! want = %v have = %v", want, have)
	}
	if have, want := card.Rarity.MaxLevel, 40; want != have {
		t.Errorf("Error on fields Max Level! want = %v have = %v", want, have)
	}
	if have, want := card.LeadSkill.Name, "Cute Makeup"; want != have {
		t.Errorf("Error on fields LeadSkill name! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.ID, 100001; want != have {
		t.Errorf("Error on fields skill id! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.Timer, 6; want != have {
		t.Errorf("Error on fields skill timer! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.ProcChance[1], 4500; want != have {
		t.Errorf("Error on fields skill proc chance! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.EffectLength[1], 600; want != have {
		t.Errorf("Error on fields effect length! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.SkillType.Name, "Score Bonus"; want != have {
		t.Errorf("Error on fields skilltype name! want = %v have = %v", want, have)
	}
	if have, want := card.Skill.ActivationCost, 0; want != have {
		t.Errorf("Error on fields skill activation cost! want = %v have = %v", want, have)
	}
	if have, want := card.BonusDance, 106; want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.BonusHp, 2; want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.BonusVisual, 167; want != have {
		t.Errorf("Error on fields BonusHp! want = %v have = %v", want, have)
	}
	if have, want := card.BonusVocal, 87; want != have {
		t.Errorf("Error on fields BonusVocal! want = %v have = %v", want, have)
	}
	if have, want := card.DanceMax, 2115; want != have {
		t.Errorf("Error on fields DanceMax! want = %v have = %v", want, have)
	}
	if have, want := card.DanceMin, 1248; want != have {
		t.Errorf("Error on fields DanceMin! want = %v have = %v", want, have)
	}
	if have, want := card.HpMax, 25; want != have {
		t.Errorf("Error on fields HpMax! want = %v have = %v", want, have)
	}
	if have, want := card.HpMin, 25; want != have {
		t.Errorf("Error on fields HpMin! want = %v have = %v", want, have)
	}
	if have, want := card.VisualMax, 3330; want != have {
		t.Errorf("Error on fields VisualMax! want = %v have = %v", want, have)
	}
	if have, want := card.VisualMin, 1965; want != have {
		t.Errorf("Error on fields VisualMin! want = %v have = %v", want, have)
	}
	if have, want := card.VocalMax, 1728; want != have {
		t.Errorf("Error on fields VocalMax! want = %v have = %v", want, have)
	}
	if have, want := card.VocalMin, 1020; want != have {
		t.Errorf("Error on fields VocalMin! want = %v have = %v", want, have)
	}
}

func TestSameInstance(t *testing.T) {
	cm, _ := Default()
	cm2, _ := Default()
	if cm != cm2 {
		t.Errorf("cardmanager.Default() returns different CardManager object")
	}
}

func BenchmarkCreateDefault10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Default()
	}
}
