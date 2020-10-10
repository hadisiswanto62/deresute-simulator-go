package simulator

import (
	"fmt"
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/csvmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodelmanager"

	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

func init() {
	os.Chdir("../")
}
func sampleTeam(attr enum.Attribute) usermodel.Team {
	if attr != enum.AttrAll {
		return usermodel.Team{
			Ocards: [5]*usermodel.OwnedCard{
				&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: attr}}},
				&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: attr}}},
				&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: attr}}},
				&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: attr}}},
				&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: attr}}},
			},
			LeaderIndex: 2,
		}
	}
	return usermodel.Team{
		Ocards: [5]*usermodel.OwnedCard{
			&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCute}}},
			&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
			&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrPassion}}},
			&usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}}},
		},
		LeaderIndex: 2,
	}
}

// func TestGameConfigOK(t *testing.T) {
// 	type testdata struct {
// 		team            usermodel.Team
// 		songAttr        enum.Attribute
// 		leaderLeadSkill *models.LeadSkill
// 		guestLeadSkill  *models.LeadSkill
// 		expected        bool
// 	}
// 	testcases := []testdata{
// 		testdata{sampleTeam(enum.AttrCute), enum.AttrCute, &models.LeadSkillCuteUnison, &models.LeadSkillCuteUnison, true},
// 		testdata{sampleTeam(enum.AttrCute), enum.AttrAll, &models.LeadSkillCuteUnison, &models.LeadSkillCuteUnison, false},
// 		testdata{sampleTeam(enum.AttrCute), enum.AttrAll, &models.LeadSkillTricolorMakeup, &models.LeadSkillCuteUnison, false},
// 	}
// 	for _, testcase := range testcases {
// 		testcase.team.Leader().LeadSkill = testcase.leaderLeadSkill
// 		song := models.NewDefaultSong("", 26, testcase.songAttr, 0, 0)
// 		guest := usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCute}, LeadSkill: testcase.guestLeadSkill}}
// 		want := testcase.expected
// 		have := isGameConfigOk(&testcase.team, &song, &guest)
// 		if want != have {
// 			t.Errorf("unisonInCorrectSongType not working! want=%v have=%v", want, have)
// 		}
// 	}
// }

func TestUnisonInCorrectSongType(t *testing.T) {
	type testdata struct {
		guestLeadSkill *models.LeadSkill
		songAttr       enum.Attribute
		expected       bool
	}
	testcases := []testdata{
		testdata{&models.LeadSkillCuteUnison, enum.AttrCute, false},
		testdata{&models.LeadSkillCuteUnison, enum.AttrCool, true},
		testdata{&models.LeadSkillCuteUnison, enum.AttrPassion, true},
		testdata{&models.LeadSkillCuteUnison, enum.AttrAll, true},
		testdata{&models.LeadSkillCoolUnison, enum.AttrCute, true},
		testdata{&models.LeadSkillCoolUnison, enum.AttrCool, false},
		testdata{&models.LeadSkillCoolUnison, enum.AttrPassion, true},
		testdata{&models.LeadSkillCoolUnison, enum.AttrAll, true},
		testdata{&models.LeadSkillPassionUnison, enum.AttrCute, true},
		testdata{&models.LeadSkillPassionUnison, enum.AttrCool, true},
		testdata{&models.LeadSkillPassionUnison, enum.AttrPassion, false},
		testdata{&models.LeadSkillPassionUnison, enum.AttrAll, true},
	}
	team := usermodel.Team{}
	for _, testcase := range testcases {
		song := models.NewDefaultSong("", 26, testcase.songAttr, 0, 0)
		guest := usermodel.OwnedCard{Card: &models.Card{LeadSkill: testcase.guestLeadSkill}}
		want := testcase.expected
		have := unisonInCorrectSongType.IsViolated(&team, &song, &guest)
		if want != have {
			t.Errorf("unisonInCorrectSongType not working! want=%v have=%v", want, have)
		}
	}
}

// func TestBothLeadSkillsActive(t *testing.T) {
// 	// testData assumes ALL COOL cards (including the guest)
// 	type testdata struct {
// 		leaderLeadSkill *models.LeadSkill
// 		guestLeadSkill  *models.LeadSkill
// 		expected        bool
// 	}
// 	testcases := []testdata{
// 		testdata{&models.LeadSkillCoolPrincess, &models.LeadSkillCoolUnison, false},
// 		testdata{&models.LeadSkillCutePrincess, &models.LeadSkillCoolUnison, true},
// 		testdata{&models.LeadSkillCutePrincess, &models.LeadSkillCuteUnison, true},
// 		testdata{&models.LeadSkillCutePrincess, &models.LeadSkillCuteUnison, true},
// 	}
// 	team := sampleTeam(enum.AttrCool)
// 	for _, testcase := range testcases {
// 		song := models.NewDefaultSong("", 26, enum.AttrAll, 0, 0)
// 		team.Leader().LeadSkill = testcase.leaderLeadSkill
// 		guest := usermodel.OwnedCard{Card: &models.Card{Idol: &models.Idol{Attribute: enum.AttrCool}, LeadSkill: testcase.guestLeadSkill}}
// 		want := testcase.expected
// 		have := bothLeadSkillIsActive.IsViolated(&team, &song, &guest)
// 		if want != have {
// 			t.Errorf("bothLeadSkillsActive not working! want=%v have=%v", want, have)
// 		}
// 	}
// }

func TestTriColorCorrectStat(t *testing.T) {
	dp := csvmodels.CSVDataParser{}
	ocards, err := usermodelmanager.ParseOwnedCard(dp, "userdata/cards.csv", nil)
	if err != nil {
		panic(err)
	}
	guests, err := usermodelmanager.ParseOwnedCard(dp, "userdata/guest tricolor.csv", nil)
	if err != nil {
		panic(err)
	}
	cardIDs := [5]int{
		// 300572, 300571, 300856, 100798, 300740, (card orang, leader=2)
		300830, 300572, 300236, 200314, 200726,
	}
	leaderIndex := 2

	cards := [5]*usermodel.OwnedCard{}
	for i, id := range cardIDs {
		for _, ocard := range ocards {
			if id == ocard.Card.ID {
				cards[i] = ocard
				break
			}
		}
	}
	team := usermodel.Team{Ocards: cards, LeaderIndex: leaderIndex}
	song := &models.Song{Attribute: enum.AttrAll}
	if isTeamOk(&team, song) {
		for _, guest := range guests {
			fmt.Println(guest.LeadSkill.Name, isGameConfigOkDebug(&team, song, guest))
		}
	} else {
		fmt.Println(isTeamOkDebug(&team, song))
	}
}
