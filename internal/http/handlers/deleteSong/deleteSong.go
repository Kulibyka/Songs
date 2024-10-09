package deleteSong

import (
	"Kulibyka/internal/storage/postgresql"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

func DeleteSongHandler(db *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid song ID", http.StatusBadRequest)
			return
		}

		if err := db.DeleteSong(int64(id)); err != nil {
			slog.Error("Failed to delete song", slog.String("error", err.Error()))
			http.Error(w, "Failed to delete song", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
