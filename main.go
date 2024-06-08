package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/core"

	handler "bravian1/team-shuffler/src/routes"
)

const (
	PORT = 8000
	HOST = "localhost"
	APP  = "Team Management App"
)

func main() {
	result, success := core.OpenOrCreate("teams.txt")
	if !success {
		log.Fatalf(result)
	}

	port := PORT
	host := HOST
	app := APP
	config, success := core.ReadConfig()
	if !success {
		fmt.Printf("\t\t[[Unable to read configuration file. Using default port: %d]]\n\n", PORT)
	} else {
		port = config.HostPort
		host = config.HostName
		app = config.AppName
	}

	url := fmt.Sprintf("%s:%d", host, port)

	mutex := &sync.Mutex{}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler.Home)

	// serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/players", handler.Players)
	http.HandleFunc("/fixtures", handler.Fixtures)
	http.HandleFunc("/shuffle", handler.Shuffle(mutex))
	http.HandleFunc("/register", handler.Register(mutex))
	http.HandleFunc("/toroot", handler.Index)

	fmt.Printf("\n\n\t---[%s]---\n\n\tServer running at %s:%d\n\n", app, host, port)

	http.ListenAndServe(url, nil)
}
