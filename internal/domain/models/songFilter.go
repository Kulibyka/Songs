package models

type SongFilter struct {
	Group       string     `json:"group,omitempty"`
	Song        string     `json:"song,omitempty"`
	ReleaseDate string     `json:"release_date,omitempty"`
	Pagination  Pagination `json:"pagination"`
}
