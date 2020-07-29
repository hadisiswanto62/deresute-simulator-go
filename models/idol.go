package models

import (
	"github.com/hadisiswanto62/deresute-simulator-go/enum"
)

// Idol is a idol
type Idol struct {
	ID        int            `json:"chara_id"`
	Name      string         `json:"conventional"`
	Attribute enum.Attribute `json:"type"`
}
