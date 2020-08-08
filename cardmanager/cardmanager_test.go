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
	card := cm.Filter().ID(300083).First()
	if have, want := card.ID, 300083; want != have {
		t.Errorf("Error on fields ID! want = %v have = %v", want, have)
	}
	if have, want := card.Attribute, enum.AttrPassion; want != have {
		t.Errorf("Error on fields Attr! want = %v have = %v", want, have)
	}
	if have, want := card.Name, "Wakabayashi Tomoka"; want != have {
		t.Errorf("Error on fields name! want = %v have = %v", want, have)
	}
	if have, want := card.Rarity.Rarity, enum.RarityR; want != have {
		t.Errorf("Error on fields Rarity! want = %v have = %v", want, have)
	}
	if have, want := card.SeriesID, 300083; want != have {
		t.Errorf("Error on fields SeriesID! want = %v have = %v", want, have)
	}
	if have, want := card.BonusDance, uint16(184); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.BonusHp, uint16(2); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.BonusVisual, uint16(119); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.BonusVocal, uint16(98); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.DanceMax, uint16(3670); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.DanceMin, uint16(2166); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.HpMax, uint16(25); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.HpMin, uint16(25); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.OverallBonus, uint16(401); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.OverallMax, uint16(7979); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.OverallMin, uint16(4709); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.VisualMax, uint16(2365); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.VisualMin, uint16(1396); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.VocalMax, uint16(1944); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
	}
	if have, want := card.VocalMin, uint16(1147); want != have {
		t.Errorf("Error on fields BonusDance! want = %v have = %v", want, have)
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
