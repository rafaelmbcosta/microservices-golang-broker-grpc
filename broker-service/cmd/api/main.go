package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3001"

type Config struct{}

func main() {
	// app := Config{}

	log.Println("Starting server on port 3001")

	//define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		// Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
