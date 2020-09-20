package songmanager

import (
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

func init() {
	os.Chdir("../")
}

func TestDefaultInstance(t *testing.T) {
	a, err := Default()
	b, _ := Default()
	if err != nil {
		t.Errorf("cannot create songmaanger object: %v", err)
	}
	if a != b {
		t.Errorf("Default returns different instance!")
	}
	if len(a.Songs) == 0 {
		t.Errorf("It does not create song!")
	}
}

func TestQueryAttribute(t *testing.T) {
	a, _ := Default()
	song := a.Filter().Attribute(enum.AttrCute).First()
	if song.Attribute != enum.AttrCute {
		t.Errorf("Filter by attr not working")
	}
}

func TestQueryNameLike(t *testing.T) {
	a, _ := Default()
	songs := a.Filter().NameLike("Sunshine").Get()
	if len(songs) > 20 {
		t.Errorf("Filter by name not working, len=%d", len(songs))
	}
}

func TestQueryDifficulty(t *testing.T) {
	a, _ := Default()
	songs := a.Filter().Difficulty(enum.SongDifficultyMasterPlus).Get()
	if len(songs) < 10 {
		t.Errorf("Filter by diff not working, len=%d", len(songs))
	}
}
