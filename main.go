package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"bravian1/team-shuffler/src/core"
	"bravian1/team-shuffler/src/data"
	"bravian1/team-shuffler/src/routes"
	"bravian1/team-shuffler/src/types"

	handler "bravian1/team-shuffler/src/routes"

	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	PORT = 8765
	HOST = "localhost"
	APP  = "Team Management App"
)

func main() {
	db, err := data.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	routes.DB = db
	// Auto Migrate the schema
	db.AutoMigrate(&(types.Player{}), &types.Team{}, &types.Table{}, &types.Fixture{})

	// err := core.DownloadFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result, success := core.OpenOrCreate("teams.txt")
	// if !success {
	// 	log.Fatalf(result)
	// }

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

	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/", handler.Home)

	// serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/admin/players", handler.AddPlayers)
	http.HandleFunc("/rules", handler.RulesHandler)
	http.HandleFunc("/table", handler.Table)
	http.HandleFunc("/playerlist", handler.PlayerlistHandler)
	http.HandleFunc("/players", handler.Players)

	http.HandleFunc("/fixtures", handler.Fixtures)
	http.HandleFunc("/shuffle", handler.Shuffle)
	http.HandleFunc("/register", handler.Register(mutex))
	http.HandleFunc("/toroot", handler.Index)

	fmt.Printf("\n\n\t---[%s]---\n\n\tServer running at http://%s:%d\n\n", app, host, port)

	http.ListenAndServe(url, nil)
}
