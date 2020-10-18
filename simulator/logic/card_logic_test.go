package logic

import (
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/stretchr/testify/assert"

	"github.com/hadisiswanto62/deresute-simulator-go/cardmanager"
)

var cm *cardmanager.CardManager

func init() {
	os.Chdir("../../")
	var err error
	cm, err = cardmanager.Default()
	if err != nil {
		panic(err)
	}
}

type cardLogicTestcase struct {
	name     string
	nameID   string
	cardID   int
	expected bool
}

func testCardLogic(t *testing.T, testcases []cardLogicTestcase, logic *cardLogic) bool {
	for _, tc := range testcases {
		var card *models.Card
		// prioritize nameID
		if tc.nameID != "" {
			card = cm.Filter().SsrNameID(tc.nameID).First()
		} else if tc.cardID != 0 {
			card = cm.Filter().ID(tc.cardID).First()
		} else {
			t.Error("insufficient test data")
			return false
		}
		ocard, err := usermodel.NewOwnedCardBuilder().Card(card).Build()
		if err != nil {
			t.Errorf("cannot create ocard: %v", err)
		}
		actual := logic.isSatisfied(ocard, nil)
		result := assert.Equalf(t, tc.expected, actual, "Wrong result for: %s", tc.name)
		if !result {
			return result
		}
	}
	return true
}

func TestCardLogic_CardIsSSR(t *testing.T) {
	testcases := []cardLogicTestcase{
		{name: "SSR", cardID: 100076, expected: true},
		{name: "SR", cardID: 100044, expected: false},
		{name: "R", cardID: 100002, expected: false},
	}
	logic := cardCardIsSSR
	testCardLogic(t, testcases, logic)
}

func TestCardLogic_SkillIsNotConcentration(t *testing.T) {
	testcases := []cardLogicTestcase{
		{name: "not concentration1", nameID: "yoshino4", expected: true},
		{name: "not concentration2", nameID: "uzuki4", expected: true},
		{name: "concentration", nameID: "megumi1", expected: false},
	}
	logic := cardSkillIsNotConcentration
	testCardLogic(t, testcases, logic)
}

func TestCardLogic_SkillIsImplemented(t *testing.T) {
	testcases := []cardLogicTestcase{
		{name: "implemented1", nameID: "yoshino4", expected: true},
		{name: "implemented2", nameID: "uzuki4", expected: true},
		{name: "implemented3", nameID: "megumi1", expected: true},
		{name: "not implemented", nameID: "nagi2", expected: false},
	}
	logic := cardSkillIsImplemented
	testCardLogic(t, testcases, logic)
}
