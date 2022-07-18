package main

import (
	"log"
	"net/http"

	"github.com/mishrabhi0123/screen-time/handlers"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/registerUser", handlers.RegisterUser)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/device/", handlers.GetDevice)

	message := "Server started at http://localhost:" + PORT
	log.Println(message)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
