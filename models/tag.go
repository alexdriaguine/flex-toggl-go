package models

import "time"

// Tag is a toggle tag
type Tag struct {
	ID   int       `json:"id"`
	Wid  int       `json:"wid"`
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}
