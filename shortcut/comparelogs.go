package shortcut

import (
	"fmt"
	"math"

	"github.com/hadisiswanto62/deresute-simulator-go/logparser"
)

// CompareLogs is used to get max differences between two logs (with the same game configs)
// Will return error when error > threshold
func CompareLogs(logname1, logname2 string, threshold int) (int, error) {
	log1, err := logparser.Parse(logname1, 0)
	if err != nil {
		return 0, fmt.Errorf("Cannot parse: %v", err)
	}
	log2, err := logparser.Parse(logname2, 0)
	if err != nil {
		return 0, fmt.Errorf("Cannot parse: %v", err)
	}

	maxDiff := 0
	for _, report := range log1 {
		for _, report2 := range log2 {
			if report.SameTeamAs(report2) {
				difference := int(math.Abs(float64(report.Average) - float64(report2.Average)))
				if difference > threshold {
					return 0, fmt.Errorf("suspicious result on %v %d %d, (1=%d, 2=%d, %d)",
						report.CardIDs, report.LeaderIndex, report.GuestID,
						report.Average, report2.Average, difference,
					)
				}
				if difference > maxDiff {
					maxDiff = difference
				}
				break
			}
		}
	}
	return maxDiff, nil
}
