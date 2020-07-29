package models

import (
	"example.com/deresute/testo/enum"
)

// Idol is a idol
type Idol struct {
	ID        int            `json:"chara_id"`
	Name      string         `json:"conventional"`
	Attribute enum.Attribute `json:"type"`
}
