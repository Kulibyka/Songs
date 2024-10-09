package addSong

import (
	"Kulibyka/internal/domain/models"
	"Kulibyka/internal/storage/postgresql"
	"encoding/json"
	"log/slog"
	"net/http"
)

// AddSongHandler @Summary Добавить песню
//
//	// @Description Добавляет новую песню в библиотеку
//	// @Accept  json
//	// @Produce  json
//	// @Param song body models.Song true "Данные песни"
//	// @Success 201 {object} models.Song
//	// @Failure 400 {string} string "Invalid request body"
//	// @Failure 500 {string} string "Failed to create song"
//	// @Router /songs/add [post]
func AddSongHandler(db *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var song models.Song

		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if _, err := db.CreateSong(song); err != nil {
			slog.Error("Failed to create song", slog.String("error", err.Error()))
			http.Error(w, "Failed to create song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(song)
	}
}
