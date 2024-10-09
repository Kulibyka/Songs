package api

import (
	"Kulibyka/internal/http/handlers/addSong"
	"Kulibyka/internal/http/handlers/deleteSong"
	"Kulibyka/internal/http/handlers/getSongCouplets"
	"Kulibyka/internal/http/handlers/getSongsWithFilter"
	"Kulibyka/internal/http/handlers/updateSong"
	"Kulibyka/internal/storage/postgresql"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *postgresql.Storage) {
	router.HandleFunc("/songs/add", addSong.AddSongHandler(db)).Methods("POST")

	router.HandleFunc("/songs/delete/{id}", deleteSong.DeleteSongHandler(db)).Methods("DELETE")

	router.HandleFunc("/songs/update/{id}", updateSong.UpdateSongHandler(db)).Methods("PATCH")

	router.HandleFunc("/songs/songtext/{id}", getSongCouplets.GetSongCouplets(db)).Methods("GET")

	router.HandleFunc("/songs/filter", getSongsWithFilter.GetSongsWithFilterHandler(db)).Methods("GET")
}
