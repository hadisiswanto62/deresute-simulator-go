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

func doJob(gameConfig *GameConfig, channel chan SimulationSummary) {
	channel <- Simulate(gameConfig, 1)
}

func FindOptimal(album *usermodel.Album, guests []*usermodel.OwnedCard,
	song *models.Song, beneran bool, times int, filename string) error {
	defer helper.MeasureTime(time.Now(), "FindOptimal")

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
			actualNumberofResults++
			guestCount = append(guestCount, guest)
			// if len(guestCount) > 1 {
			// 	fmt.Println("No guest: ")
			// 	fmt.Println(team)
			// 	da, vo, vi := 0, 0, 0
			// 	for _, ocard := range team.Ocards {
			// 		fmt.Println(ocard.Card.Idol.Attribute, ocard.Skill.SkillType.Name, ocard.LeadSkill.Name)
			// 		da += ocard.Dance
			// 		vo += ocard.Vocal
			// 		vi += ocard.Visual
			// 	}
			// 	fmt.Println(da, vo, vi)
			// 	fmt.Println(team.Leader().LeadSkill.Name)
			// 	for _, g := range guestCount {
			// 		fmt.Println("Guest = ", g.LeadSkill.Name)
			// 	}
			// 	fmt.Println()
			// }
			if beneran {
				gameConfig := NewGameConfig(team, supports, guest, song)
				go func(gameConfig *GameConfig, channel chan SimulationSummary) {
					channel <- Simulate(gameConfig, times)
					// resultChannel <- SimulationSummary{Average: 100}
				}(gameConfig, resultChannel)
			}

		}
		// if len(guestCount) == 0 {
		// 	if team.Leader().Card.ID == 300830 {
		// 		continue
		// 	}
		// 	fmt.Println("No guest: ")
		// 	fmt.Println(team)
		// 	da, vo, vi := 0, 0, 0
		// 	for _, ocard := range team.Ocards {
		// 		fmt.Println(ocard.Card.Idol.Attribute, ocard.Skill.SkillType.Name, ocard.LeadSkill.Name)
		// 		da += ocard.Dance
		// 		vo += ocard.Vocal
		// 		vi += ocard.Visual
		// 	}
		// 	fmt.Println(da, vo, vi)
		// 	fmt.Println(team.Leader().LeadSkill.Name)
		// 	for _, guest := range guests {
		// 		fmt.Println(guest.LeadSkill.Name, isGameConfigOkDebug(team, song, guest))
		// 	}
		// 	fmt.Println()
		// }
		if i%100000 == 0 {
			fmt.Println(i)
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
	for _, summary := range summaries {
		ids := []string{}
		for _, ocard := range summary.GameConfig.team.Ocards {
			ids = append(ids, strconv.Itoa(ocard.Card.ID))
		}
		id := strings.Join(ids, ",")
		str := fmt.Sprintf("%s %d %d %d %.0f",
			id, summary.GameConfig.team.LeaderIndex, summary.GameConfig.guest.Card.ID, summary.Appeal, summary.Average)
		buffer = append(buffer, str)
		str = ""
		for _, ocard := range summary.GameConfig.team.Ocards {
			str = fmt.Sprintf("%s %s %s | ", str, ocard.Card.Idol.Name, ocard.Card.Skill.SkillType.Name)
		}
		str = fmt.Sprintf("%s %s %d %s %s %.0f", str, summary.GameConfig.team.Leader().LeadSkill.Name, summary.GameConfig.team.LeaderIndex,
			summary.GameConfig.guest.Card.Idol.Attribute, summary.GameConfig.guest.Card.LeadSkill.Name, summary.Average)
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
