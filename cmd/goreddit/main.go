package main

import (
	"github.com/demyanovs/go-reddit/postgres"
	"github.com/demyanovs/go-reddit/web"
	"log"
	"net/http"
)

func main()  {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
