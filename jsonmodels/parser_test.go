package jsonmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	os.Chdir("../")
}

func TestParseLeadSkill(t *testing.T) {
	dp := JSONDataParser{}
	tmpLeadSkill := []TmpLeadSkill{sampleLeadSkill}
	lskills, err := dp.parseLeadSkills(tmpLeadSkill)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}
	fmt.Println(len(lskills))
	for _, lskill := range lskills {
		fmt.Printf("%v", lskill)
	}
}

func TestParseAllLeadSkill(t *testing.T) {
	dp := JSONDataParser{}
	text, _ := ioutil.ReadFile(leadSkillPath)
	var tmpLeadSkills []TmpLeadSkill
	json.Unmarshal(text, &tmpLeadSkills)
	lskills, err := dp.parseLeadSkills(tmpLeadSkills)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}
	for _, lskill := range lskills {
		fmt.Printf("%p: %v\n", lskill, *lskill)
	}
}
