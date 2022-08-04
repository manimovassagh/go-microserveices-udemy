package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker Micro-Service on port %s\n", webPort)
	//define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	//Start The Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Can not Listen and Serve server 😢", err)
	}

}
