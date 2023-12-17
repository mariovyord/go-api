package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "3000"

type Config struct{}

func main() {
	app := Config{}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: app.routes(), 
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Server is listening on port %s", port)
}

