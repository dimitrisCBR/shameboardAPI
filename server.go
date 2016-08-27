package main

import (
	"log"
	"net/http"

	"github.com/dimitrisCBR/shameboardAPI/v2/routes"
)

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8888", router))
}
