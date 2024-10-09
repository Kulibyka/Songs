package updateSong

import (
	"Kulibyka/internal/domain/models"
	"Kulibyka/internal/storage/postgresql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

// UpdateSongHandler @Summary Обновить данные песни
// @Description Обновляет данные песни по ID
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновленные данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Song not found"
// @Router /songs/update/{id} [patch]
func UpdateSongHandler(db *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid song ID", http.StatusBadRequest)
			return
		}

		var song models.Song

		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := db.UpdateSong(int64(id), song); err != nil {
			slog.Error("Failed to patch song", slog.String("error", err.Error()))
			http.Error(w, "Failed to update song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
