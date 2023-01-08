package models

import "time"

type Topic struct {
	Id        uint64    `json:"id,omitempty"`
	Topic     string    `json:"topic,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
