package main

import (
	"database/sql"
	"fmt"
	"go-service-auth/data"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting the authentication service")

	// TODO connect to DB

	// setup config
	app := Config{}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.route(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
