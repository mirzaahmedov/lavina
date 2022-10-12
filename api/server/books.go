package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mirzaahmedov/lavina/api/types"
)

func (s *Server)createBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Isbn string `json:"isbn"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Println(err)

			w.WriteHeader(400)
			
			json.NewEncoder(w).Encode(&types.Response[types.Book]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}
		defer r.Body.Close()

		response, err := http.Get("https://openlibrary.org/isbn/" + body.Isbn + ".json")
		if err != nil {
			log.Println(err)

			w.WriteHeader(404)
			
			json.NewEncoder(w).Encode(&types.Response[types.Book]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}

		var book types.Book
		if err := json.NewDecoder(response.Body).Decode(&book); err != nil {
			log.Println(err)

			w.WriteHeader(500)
			
			json.NewEncoder(w).Encode(&types.Response[types.Book]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}

		newBook, err := s.repository.CreateBook(&book)
		if err != nil {
			log.Println(err)

			w.WriteHeader(500)

			json.NewEncoder(w).Encode(&types.Response[types.Book]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(&types.Response[types.Book]{
			IsOk: true,
			Message: "ok",
			Data: newBook,
		})
	}
}