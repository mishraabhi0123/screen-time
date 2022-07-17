package main

import (
	"log"
	"net/http"

	"github.com/mishrabhi0123/screen-time/handlers"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/", handlers.RegisterUser)

	message := "Server started at http://localhost:" + PORT
	log.Println(message)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
