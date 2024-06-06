package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	handler "bravian1/team-shuffler/routes"
	"bravian1/team-shuffler/types"
)

var mu sync.Mutex

func main() {
	file, err := os.Open("teams.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	http.HandleFunc("/players", handler.Players)
	http.HandleFunc("/shuffle", handler.Shuffle)
	http.HandleFunc("/register", registerHandlers)
	http.HandleFunc("/", handler.Index)
	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", nil)
}

func registerHandlers(w http.ResponseWriter, r *http.Request) {
	role := r.FormValue("role")
	name := r.FormValue("name")
	fmt.Println(role + " " + name)
	if role == "" || name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	line := fmt.Sprintf("\n%s: %s", role, name)

	file, err := os.OpenFile("teams.txt", os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(line)
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}
	// players := []Player{}
	mu.Lock()
	player := types.Player{Role: role, Name: name}
	mu.Unlock()
	w.Header().Set("content-type", "application.json")
	json.NewEncoder(w).Encode(player)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}
