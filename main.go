package main

import (
	"log"
	"net/http"

	"github.com/drinks-protobuf/routes"
)

func main() {
	router := routes.NewRouter()

	// launch server
	log.Println(http.ListenAndServe(":10000", router))
}
