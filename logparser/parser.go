package logparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReportItem struct {
	CardIDs     [5]int
	LeaderIndex int
	GuestID     int
	Average     int
}

func (r ReportItem) Report() {
	fmt.Printf("%v %d %d %d\n", r.CardIDs, r.LeaderIndex, r.GuestID, r.Average)
}

func (r ReportItem) SameTeamAs(other ReportItem) bool {
	itemMap := make(map[int]bool, 5)
	for _, id := range r.CardIDs {
		itemMap[id] = true
	}
	for _, id := range other.CardIDs {
		_, ok := itemMap[id]
		if !ok {
			return false
		}
	}
	return (r.CardIDs[r.LeaderIndex] == other.CardIDs[other.LeaderIndex] &&
		r.GuestID == other.GuestID)
}

func Parse(filename string, topX int) ([]ReportItem, error) {
	if topX == 0 {
		topX = 9999999999
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot parse log: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	configs := []ReportItem{}
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		idsStr := splitted[0]
		leaderIndex, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, fmt.Errorf("cannot parse log: %v", err)
		}
		guestID, err := strconv.Atoi(splitted[2])
		if err != nil {
			return nil, fmt.Errorf("cannot parse log: %v", err)
		}
		average, err := strconv.Atoi(splitted[4])
		if err != nil {
			return nil, fmt.Errorf("cannot parse log: %v", err)
		}

		ids := [5]int{}
		for i, id := range strings.Split(idsStr, ",") {
			ids[i], err = strconv.Atoi(id)
			if err != nil {
				return nil, fmt.Errorf("cannot parse log: %v", err)
			}
		}
		configs = append(configs, ReportItem{
			CardIDs:     ids,
			LeaderIndex: leaderIndex,
			GuestID:     guestID,
			Average:     average,
		})
		if len(configs) >= topX {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("cannot parse log: %v", err)
	}
	return configs, nil
}
