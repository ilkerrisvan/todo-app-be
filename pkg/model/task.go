package model

import "time"

type Task struct {
	CreatedAt time.Time `json:"date"`
	Text      string    `json:"text"`
}
