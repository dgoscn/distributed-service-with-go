package main

import (
	"log"

	"github.com/dgoscn/distributed-services-with-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":1020")
	log.Fatal(srv.ListenAndServe())
}

//Built a simple JSON/HTTP commit log service that accepts and responds with JSON and stores the records
//in those requests to an in-memory log.
