package routes

import (
	"encoding/json"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/types"
)

func Register(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		role := r.FormValue("role")

		if name == "" || role == "" {
			http.Error(w, "Name and role are required", http.StatusBadRequest)
			return
		}

		player := types.Player{
			Name: name,
			Role: role,
		}

		if err := DB.Create(&player).Error; err != nil {
			http.Error(w, "Error creating player", http.StatusInternalServerError)
			return
		}


		w.Header().Set("content-type", "application.json")
		json.NewEncoder(w).Encode(player)
	}
}
