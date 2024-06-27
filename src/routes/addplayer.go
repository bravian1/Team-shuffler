package routes

import (
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

// AdminPlayersHandler renders the admin page for managing players
func AddPlayers(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/admin_players.html")
}
