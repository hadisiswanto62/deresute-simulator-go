package simulator

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/helper"
	"github.com/hadisiswanto62/deresute-simulator-go/models"
	"github.com/hadisiswanto62/deresute-simulator-go/usermodel"
)

// FindOptimal log the results of all (good) combinations of cards+guest given
// the album and list of guest (for current song) to the specified path.
func FindOptimal(album *usermodel.Album, guests []*usermodel.OwnedCard,
	song *models.Song, times int, filename string) error {
	defer helper.MeasureTime(time.Now(), "FindOptimal")

	beneran := helper.Features.ReallySimulate()
	start := time.Now()
	resultChannel := make(chan SimulationSummary)
	i := 0
	expectedNumberOfResults := (album.MaxTeamID() + 1) * len(guests)
	actualNumberofResults := 0

	log.Printf("Running %d samples:\n", expectedNumberOfResults)
	helper.MeasureTime(start, "Initializing")
	start = time.Now()
	for album.Next() {
		team := album.GetTeam()
		if !isTeamOk(team, song) {
			continue
		}
		supports, err := album.FindSupportsFor(team, song.Attribute)
		i++
		if err != nil {
			return fmt.Errorf("could not find optimal: %v", err)
		}
		guestCount := []*usermodel.OwnedCard{}
		for _, guest := range guests {
			if !isGameConfigOk(team, song, guest) {
				continue
			}
			guestCount = append(guestCount, guest)
			gameConfig := NewGameConfig(team.Ocards[:], team.LeaderIndex, supports[:], guest, song)
			if helper.Features.LimitAppeals() {
				if gameConfig.getAppeal() < 310000 {
					if !gameConfig.isResonantActive() {
						continue
					}
				}
			}
			actualNumberofResults++
			if beneran {
				go func(gameConfig *GameConfig, channel chan SimulationSummary) {
					channel <- Simulate(gameConfig, times)
				}(gameConfig, resultChannel)
			}
		}
	}
	log.Printf("Finding supports %d times\n", i)
	helper.MeasureTime(start, "Queueing all jobs")
	log.Printf("%d jobs queued\n", actualNumberofResults)
	if !beneran {
		return nil
	}
	logPath := fmt.Sprintf("log/%s", filename)
	readableLogPath := fmt.Sprintf("log/readable/%s", filename)
	os.Remove(logPath)
	os.Remove(readableLogPath)
	start = time.Now()

	maxAvg := 0.0
	i = 0
	var summaries []SimulationSummary
	var maxAvgSummary SimulationSummary
	for summary := range resultChannel {
		summaries = append(summaries, summary)
		if summary.Average > maxAvg {
			maxAvg = summary.Average
			maxAvgSummary = summary
		}
		i++
		if i == actualNumberofResults {
			close(resultChannel)
		}
	}
	helper.MeasureTime(start, "Finished all jobs")
	if len(summaries) != actualNumberofResults {
		panic("")
	}
	log.Printf("%f", maxAvgSummary.Average)

	sort.SliceStable(summaries, func(i int, j int) bool {
		return summaries[i].Average > summaries[j].Average
	})
	buffer := []string{}
	readableBuffer := []string{}
	bufferMaxLength := 10000
	sampahCount := 0
	for _, summary := range summaries {
		if summary.SimCount == -1 {
			sampahCount++
			continue
		}
		ids := []string{}
		for _, ocard := range summary.GameConfig.getCards() {
			ids = append(ids, strconv.Itoa(ocard.Card.ID))
		}
		id := strings.Join(ids, ",")
		str := fmt.Sprintf("%s %d %d %d %.0f",
			id, summary.GameConfig.getLeaderIndex(),
			summary.GameConfig.getGuest().Card.ID, summary.GameConfig.getAppeal(), summary.Average)
		buffer = append(buffer, str)
		str = ""
		for _, ocard := range summary.GameConfig.getCards() {
			str = fmt.Sprintf("%s %s %s | ", str, ocard.Card.Idol.Name, ocard.Card.Skill.SkillType.Name)
		}
		str = fmt.Sprintf("%s %s %d %s %s %.0f", str,
			summary.GameConfig.getCards()[summary.GameConfig.getLeaderIndex()].LeadSkill.Name,
			summary.GameConfig.getLeaderIndex(),
			summary.GameConfig.getGuest().Card.Idol.Attribute,
			summary.GameConfig.getGuest().Card.LeadSkill.Name, summary.Average)
		readableBuffer = append(readableBuffer, str)
		if len(buffer) > bufferMaxLength {
			saveBuffer(&buffer, fmt.Sprintf("log/%s", filename))
			buffer = []string{}
		}
		if len(readableBuffer) > bufferMaxLength {
			saveBuffer(&readableBuffer, fmt.Sprintf("log/readable/%s", filename))
			readableBuffer = []string{}
		}
	}
	log.Printf("Found %d sampahs\n", sampahCount)
	saveBuffer(&buffer, fmt.Sprintf("log/%s", filename))
	saveBuffer(&readableBuffer, fmt.Sprintf("log/readable/%s", filename))
	return nil
}

func saveBuffer(buffer *[]string, filename string) {
	if len(*buffer) == 0 {
		return
	}
	text := strings.Join(*buffer, "\n")
	text += "\n"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		panic(err)
	}
	fmt.Println("Writing:")
}
