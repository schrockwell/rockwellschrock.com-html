package server

import (
	"log"
	"net/http"
)

func Start() {
	fs := http.FileServer(http.Dir("./_out"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
