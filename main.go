package main

import (
	"log"
	"os"
	"net/http"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// static
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))));

	log.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}