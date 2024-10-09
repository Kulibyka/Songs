package postgresql

import (
	"Kulibyka/internal/config"
	"Kulibyka/internal/domain/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(cfg config.PostgresConfig) (*Storage, error) {
	const op = "storage.postgresql.New"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateSong(song models.Song) (int64, error) {
	const op = "storage.postgresql.CreateSong"

	query := `INSERT INTO songs (group_name, song, release_date, text, link, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int64
	err := s.db.QueryRow(query,
		song.GroupName, song.Song, song.ReleaseDate, song.Text, song.Link, time.Now()).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) DeleteSong(id int64) error {
	const op = "storage.postgresql.DeleteSong"

	query := `DELETE FROM songs WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdateSong(id int64, song models.Song) error {
	const op = "storage.postgresql.UpdateSong"

	var queryParts []string
	var args []interface{}
	if song.GroupName != "" {
		queryParts = append(queryParts, "group_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, song.GroupName)
	}
	if song.Song != "" {
		queryParts = append(queryParts, "song = $"+strconv.Itoa(len(args)+1))
		args = append(args, song.Song)
	}
	if song.ReleaseDate != "" {
		queryParts = append(queryParts, "release_date = $"+strconv.Itoa(len(args)+1))
		args = append(args, song.ReleaseDate)
	}
	if song.Text != "" {
		queryParts = append(queryParts, "text = $"+strconv.Itoa(len(args)+1))
		args = append(args, song.Text)
	}
	if song.Link != "" {
		queryParts = append(queryParts, "link = $"+strconv.Itoa(len(args)+1))
		args = append(args, song.Link)
	}

	if len(queryParts) == 0 {
		return nil
	}

	query := "UPDATE songs SET " + strings.Join(queryParts, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1)
	args = append(args, id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetSongLyrics(id int64, pagination models.Pagination) ([]string, error) {
	const op = "storage.postgresql.GetSongLyrics"

	var text string
	query := "SELECT text FROM songs WHERE id = $1"

	err := s.db.QueryRow(query, id).Scan(&text)
	if err != nil {
		return []string{}, fmt.Errorf("%s: %w", op, err)
	}

	lines := strings.Split(text, "\n")
	var couplets []string

	for i := 0; i < len(lines); i += 4 {
		end := i + 4
		if end > len(lines) {
			end = len(lines)
		}
		couplet := strings.Join(lines[i:end], "\n")
		couplets = append(couplets, couplet)
	}

	start := (pagination.PageNum - 1) * pagination.LimitNum
	if start > len(couplets) {
		return nil, nil
	}

	end := start + pagination.LimitNum
	if end > len(couplets) {
		end = len(couplets)
	}

	return couplets[start:end], nil
}

func (s *Storage) GetSongsWithFilter(filter models.SongFilter) ([]models.Song, error) {
	const op = "storage.postgresql.GetSongsWithFilter"

	query := "SELECT id, group_name, song, release_date, text, link, created_at FROM songs WHERE 1=1"
	var args []interface{}
	argID := 1

	if filter.GroupName != "" {
		query += fmt.Sprintf(" AND group_name = $%d", argID)
		args = append(args, filter.GroupName)
		argID++
	}
	if filter.Song != "" {
		query += fmt.Sprintf(" AND song = $%d", argID)
		args = append(args, filter.Song)
		argID++
	}
	if filter.ReleaseDate != "" {
		query += fmt.Sprintf(" AND release_date = $%d", argID)
		args = append(args, filter.ReleaseDate)
		argID++
	}

	offset := (filter.Pagination.PageNum - 1) * filter.Pagination.LimitNum
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, filter.Pagination.LimitNum, offset)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.GroupName, &song.Song,
			&song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
