package routes

import (
	"net/http"
	"os"
	"path"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
	}
	filename := path.Join(rootDir, "templates", "index.html")
	tmpl := template.Must(template.ParseFiles(filename))
	tmpl.Execute(w, nil)
}
