package getSongCouplets

import (
	"Kulibyka/internal/domain/models"
	"Kulibyka/internal/storage/postgresql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetSongCouplets(db *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid song ID", http.StatusBadRequest)
			return
		}

		var pagination models.Pagination
		if err := json.NewDecoder(r.Body).Decode(&pagination); err != nil {
			http.Error(w, "Invalid pagination parameters", http.StatusBadRequest)
			return
		}

		lyrics, err := db.GetSongLyrics(int64(id), pagination)
		if err != nil {
			http.Error(w, "Failed to get song lyrics", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lyrics)
	}
}
