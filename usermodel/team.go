package usermodel

import "fmt"

// Team is a team
type Team struct {
	Ocards      [5]*OwnedCard
	LeaderIndex int
}

// Leader returns OwnedCard that is the leader
func (t Team) Leader() *OwnedCard {
	return t.Ocards[t.LeaderIndex]
}

func (t Team) String() string {
	return fmt.Sprintf("Team [%d,%d,%d,%d,%d], leader = %d",
		t.Ocards[0].Card.ID,
		t.Ocards[1].Card.ID,
		t.Ocards[2].Card.ID,
		t.Ocards[3].Card.ID,
		t.Ocards[4].Card.ID,
		t.LeaderIndex,
	)
}
