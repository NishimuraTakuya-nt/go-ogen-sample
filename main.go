package main

import (
	"log"
	"net/http"

	petstore "github.com/NishimuraTakuya-nt/go-ogen-sample/petstore"
)

func main() {
	// Create service instance.
	h := &handler{
		pets: map[int64]petstore.Pet{},
	}
	// Create generated server.
	srv, err := petstore.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8180", srv); err != nil {
		log.Fatal(err)
	}
}
