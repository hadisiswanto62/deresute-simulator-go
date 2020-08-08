package jsonmodels

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/models"
)

func makeLeadSkill(tmp TmpLeadSkill) (*models.LeadSkill, error) {
	lskill, err := models.GetLeadSkill(tmp.Name)
	if err != nil {
		return lskill, fmt.Errorf("could not make lead skill: %v", err)
	}
	return lskill, nil
}
