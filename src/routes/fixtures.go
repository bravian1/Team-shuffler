package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/types"
)

func Fixtures(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}
	filename := path.Join(rootDir, "storage", "teams.txt")

	stats, err := os.Stat(filename)
	if err!= nil {
        http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
    }
	if stats.Size() == 0 {
        http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
    }

	content, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}
	teams := []types.Teams{}

	err = json.Unmarshal(content, &teams)
	if err != nil {
		http.Error(w, "Error creating fixtures", http.StatusInternalServerError)
	}

	fixtures := core.Fixture(teams)

	fmt.Println(fixtures)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}
