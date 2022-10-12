package main

import (
	"log"

	"github.com/mirzaahmedov/lavina/api/server"
	"github.com/mirzaahmedov/lavina/repository"
)

func main() {
	repository := repository.New()

	if err := repository.Open(); err != nil {
		log.Fatal(err)
	}
	defer repository.Close()

	srv := server.New(repository)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}