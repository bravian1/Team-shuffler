package routes

import (
	"html/template"
	"net/http"
	"os"
	"path"

	"gorm.io/gorm"
)

var DB *gorm.DB

// AdminPlayersHandler renders the admin page for managing players
func AddPlayers(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
	}
	filename := path.Join(rootDir, "templates", "admin_players.html")
	tmpl := template.Must(template.ParseFiles(filename))
	tmpl.Execute(w, nil)
}
