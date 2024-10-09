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
