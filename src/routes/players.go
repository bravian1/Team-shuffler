package routes

import (
	"encoding/json"
	"net/http"

	"bravian1/team-shuffler/src/data"
	"bravian1/team-shuffler/src/types"
)

func Players(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	var players []types.Player
	result := db.Find(&players)
	if result.Error != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func PlayerlistHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/players.html")
}
