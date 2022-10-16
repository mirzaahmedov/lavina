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

// func logger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
// 		log.Println(r)

// 		next.ServeHTTP(w, r)
// 	})
// }

func (s *Server)Start() error {
	var address string = ":5000"
	
	if os.Getenv("PORT") != "" {
		address = ":" + os.Getenv("PORT")
	}

	// s.router.Use(logger)
	
	s.router.HandleFunc("/signup", s.signUpHandler()).Methods("POST")	
	s.router.HandleFunc("/myself", s.mySelfHandler()).Methods("GET")

	s.router.HandleFunc("/books", s.createBookHandler()).Methods("POST")
	s.router.HandleFunc("/books", s.getAllBooksHandler()).Methods("GET")
	s.router.HandleFunc("/books/{id}", s.deleteBookHandler()).Methods("DELETE")
	s.router.HandleFunc("/books/{id}", s.updateBookHandler()).Methods("PATCH")

	log.Printf("Starting Server on %v", address)

	return http.ListenAndServe(address, s.router)
}