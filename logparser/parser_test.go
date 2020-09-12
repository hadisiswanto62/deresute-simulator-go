package logparser

import "testing"

func TestSameTeamAs(t *testing.T) {
	testCases := []struct {
		report1  ReportItem
		report2  ReportItem
		expected bool
	}{
		{
			report1:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			report2:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			expected: true,
		},
		{
			report1:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			report2:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 2},
			expected: false,
		},
		{
			report1:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			report2:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 6}, LeaderIndex: 1},
			expected: false,
		},
		{
			report1:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			report2:  ReportItem{CardIDs: [5]int{5, 4, 3, 2, 1}, LeaderIndex: 3},
			expected: true,
		},
		{
			report1:  ReportItem{CardIDs: [5]int{1, 2, 3, 4, 5}, LeaderIndex: 1},
			report2:  ReportItem{CardIDs: [5]int{5, 4, 3, 2, 1}, LeaderIndex: 1},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		if testCase.report1.SameTeamAs(testCase.report2) != testCase.expected {
			t.Errorf("wrong")
		}
	}
}
