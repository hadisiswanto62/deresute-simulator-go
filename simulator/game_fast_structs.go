package simulator

type hpCostTimestamp struct {
	timestamp int
	hpCost    int
}
type activeSkillData struct {
	activeSkillTimestamps []*activeSkillTimestamp
	hpCostTimestamps      []*hpCostTimestamp
}

type activeSkillTimestamp struct {
	cardIndex      int
	startTimestamp int
	endTimestamp   int
}
