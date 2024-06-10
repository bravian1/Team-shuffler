package routes

import "net/http"

func RulesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/rules.html")
}
