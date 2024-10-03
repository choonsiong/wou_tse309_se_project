package models

import "time"

type Publisher struct {
	ID            int       `json:"id"`
	PublisherName string    `json:"publisher_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
