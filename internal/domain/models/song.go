package models

import "time"

type Song struct {
	ID          int64     `json:"id"`
	GroupName   string    `json:"group_name"`
	Song        string    `json:"song"`
	ReleaseDate string    `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
}
