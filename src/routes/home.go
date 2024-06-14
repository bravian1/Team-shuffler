package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"text/template"

	"bravian1/team-shuffler/src/db"
)

func Home(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
	}
	filename := path.Join(rootDir, "templates", "home.html")
	tmpl := template.Must(template.ParseFiles(filename))

	tmpl.Execute(w, nil)
}

func Table(w http.ResponseWriter, r *http.Request) {
	df := db.GetPlayers()

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(df)
}
