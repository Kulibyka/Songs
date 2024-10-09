package getSongsWithFilter

import (
	"Kulibyka/internal/domain/models"
	"Kulibyka/internal/storage/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetSongsWithFilterHandler(db *postgresql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter models.SongFilter
		if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		songs, err := db.GetSongsWithFilter(filter)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to get songs", http.StatusInternalServerError)
			return
		}

		if len(songs) == 0 {
			http.Error(w, "No songs found", http.StatusNotFound)
			return
		}

		// Возвращаем результат
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(songs)
	}
}
