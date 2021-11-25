package main

import (
	"log"
	"net/http"

	"github.com/biFebriansyah/gosolid/routers"
)

func main() {
	mux := routers.New()
	log.Println("Service running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Unable to start server")
	}

}
