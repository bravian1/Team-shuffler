package routes

import (
	"fmt"
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

		// Return HTML fragment for the new player row
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", player.Role, player.Name)
	}
}
