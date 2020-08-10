package usermodel

// Team is a team
type Team struct {
	Ocards      [5]*OwnedCard
	LeaderIndex int
}

// Leader returns OwnedCard that is the leader
func (t Team) Leader() *OwnedCard {
	return t.Ocards[t.LeaderIndex]
}
