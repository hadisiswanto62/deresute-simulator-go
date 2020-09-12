package shortcut

import (
	"fmt"
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/logparser"
)

func CompareLogs(logname1, logname2 string, threshold int) error {
	log1, err := logparser.Parse(logname1, 0)
	if err != nil {
		return fmt.Errorf("Cannot parse: %v", err)
	}
	log2, err := logparser.Parse(logname2, 0)
	if err != nil {
		return fmt.Errorf("Cannot parse: %v", err)
	}

	differences := 0
	for _, report := range log1 {
		for _, report2 := range log2 {
			if report.SameTeamAs(report2) {
				difference := int(math.Abs(float64(report.Average) - float64(report2.Average)))
				if difference > threshold {
					return fmt.Errorf("suspicious result on %v %d %d, (1=%d, 2=%d)",
						report.CardIDs, report.LeaderIndex, report.GuestID,
						report.Average, report2.Average,
					)
				}
				differences += difference
				break
			}
		}
	}
	return nil
}
