package usermodelmanager

import (
	"fmt"
	"os"
	"testing"

	"github.com/hadisiswanto62/deresute-simulator-go/csvmodels"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("../")
}
func TestParseOcard(t *testing.T) {
	assertion := assert.New(t)
	// dp := jsonmodels.JSONDataParser{}
	dp := csvmodels.CSVDataParser{}
	ocards, err := ParseOwnedCard(dp, "")
	assertion.Nilf(err, "Test failed: %v", err)
	for _, ocard := range ocards {
		fmt.Println(ocard)
	}
}
