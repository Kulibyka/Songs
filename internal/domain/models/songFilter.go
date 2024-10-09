package models

type SongFilter struct {
	GroupName   string     `json:"group_name,omitempty"`
	Song        string     `json:"song,omitempty"`
	ReleaseDate string     `json:"release_date,omitempty"`
	Pagination  Pagination `json:"pagination"`
}
