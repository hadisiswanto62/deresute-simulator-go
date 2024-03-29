package shortcut

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/csvmodels"
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/simulator"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodelmanager"
)

func Optimize(config BaseOptimizeConfig,
	customOwnParams *usermodelmanager.CustomOwnedCardParameters) error {
	defer helper.MeasureTime(time.Now(), "Optimize")
	cardsPath := config.CardsPath
	guestPath := config.GuestsPath
	song := config.getSong()
	simulateTimes := 100

	dp := csvmodels.CSVDataParser{}
	baseOcards, err := usermodelmanager.ParseOwnedCard(dp, cardsPath, customOwnParams)
	if err != nil {
		panic(err)
	}
	guests, err := usermodelmanager.ParseOwnedCard(dp, guestPath, customOwnParams)
	if err != nil {
		panic(err)
	}

	ocards := []*usermodel.OwnedCard{}
	if song.Attribute == enum.AttrAll {
		ocards = baseOcards
	} else {
		for _, ocard := range baseOcards {
			if ocard.Card.Idol.Attribute == song.Attribute {
				ocards = append(ocards, ocard)
			}
		}
	}

	album := usermodel.NewAlbum(ocards)
	filename := makeFilename(config, customOwnParams)

	err = simulator.FindOptimal(album, guests, song, simulateTimes, filename)
	if err != nil {
		return err
	}
	if helper.Features.ReallySimulate() {
		fmt.Printf("Results saved to %s\n", filename)
	}
	return nil
}

func makeFilename(config BaseOptimizeConfig,
	customOwnParams *usermodelmanager.CustomOwnedCardParameters) string {
	filenameParts := []string{}
	filenameParts = append(filenameParts, string(config.getSong().Attribute))
	filenameParts = append(filenameParts, fmt.Sprintf("do=%dx", config.SimulateTimes))
	filenameParts = append(filenameParts, fmt.Sprintf("skLv=%d", customOwnParams.SkillLevel))
	filenameParts = append(filenameParts, fmt.Sprintf("potSk=%d", customOwnParams.PotSkill))
	if helper.Features.UseConcentration() {
		filenameParts = append(filenameParts, "conc")
	}
	if helper.Features.AlwaysGoodRolls() {
		filenameParts = append(filenameParts, "skillOn")
	}
	if helper.Features.UseFastGame() {
		filenameParts = append(filenameParts, "fast")
	}
	filenameParts = append(filenameParts, strconv.Itoa(int(time.Now().Unix())))
	filename := fmt.Sprintf("%s.txt", strings.Join(filenameParts, "_"))
	return filename
}
