package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mirzaahmedov/lavina/repository"
)

type Server struct {
	router *mux.Router
	repository *repository.Repository
}

func New(repository *repository.Repository) *Server {
	return &Server{
		router: mux.NewRouter(),
		repository: repository,
	}
}

func (s *Server)Start() error {
	var address string = ":5000"
	
	if os.Getenv("PORT") != "" {
		address = ":" + os.Getenv("PORT")
	}

	s.router.HandleFunc("/signup", s.SignUpHandler()).Methods("POST")
	s.router.HandleFunc("/myself", s.MySelfHandler()).Methods("GET")

	log.Printf("Starting Server on %v", address)

	return http.ListenAndServe(address, s.router)
}