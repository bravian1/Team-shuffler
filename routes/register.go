package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"bravian1/team-shuffler/core"
	"bravian1/team-shuffler/types"
)

func Register(mutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.FormValue("role")
		name := r.FormValue("name")
		fmt.Println(role + " " + name)
		if role == "" || name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		line := fmt.Sprintf("\n%s: %s", role, name)

		successWrite := false
		// use a goroutine to avoid blocking other requests while satifying an IO bound process
		go func() {
			successWrite = core.WriteLine(mutex, "teams.txt", line)
		}()

		if !successWrite {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		// players := []Player{}
		player := types.Player{Role: role, Name: name}
		w.Header().Set("content-type", "application.json")
		json.NewEncoder(w).Encode(player)
		// http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
