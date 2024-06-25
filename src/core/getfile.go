package core

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile() (err error) {

	url := "https://bravian1.github.io/players/players.txt"

	os.Mkdir("./storage", os.ModePerm)
	// Create the file
	file, err := os.Create("./storage/players.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// create a client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	// Put contents on file
	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
