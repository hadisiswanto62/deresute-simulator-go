package csvmodels

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

type CSVDataParser struct {
}

const (
	ownedCardPath = "userdata/cards.csv"
)

func (p CSVDataParser) ParseOwnedCardRawData(path string) ([]*usermodel.OwnedCardRawData, error) {
	if path == "" {
		path = ownedCardPath
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %v", ownedCardPath, err)
	}
	defer file.Close()

	tmps := []*TmpOwnedCardRawData{}
	if err := gocsv.UnmarshalFile(file, &tmps); err != nil {
		return nil, fmt.Errorf("cannot unmarshal %s: %v", ownedCardPath, err)
	}
	var ret []*usermodel.OwnedCardRawData
	for _, tocd := range tmps {
		ocd := &usermodel.OwnedCardRawData{
			CardID:     tocd.CardID,
			SkillLevel: tocd.SkillLevel,
			StarRank:   tocd.StarRank,
			PotVisual:  tocd.PotVisual,
			PotDance:   tocd.PotDance,
			PotVocal:   tocd.PotVocal,
			PotHp:      tocd.PotHp,
			PotSkill:   tocd.PotSkill,
		}
		ret = append(ret, ocd)
	}
	return ret, nil
}
