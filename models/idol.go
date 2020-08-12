package models

import (
	"fmt"

	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Idol is a idol
type Idol struct {
	ID        int            `json:"chara_id"`
	Name      string         `json:"conventional"`
	Attribute enum.Attribute `json:"type"`
}

func (i Idol) String() string {
	return fmt.Sprintf("%d %s %s", i.ID, i.Name, i.Attribute)
}
