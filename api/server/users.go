package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/mirzaahmedov/lavina/api/types"
	"github.com/mirzaahmedov/lavina/tools"
)

func (s *Server)signUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body types.User

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {				
				log.Println(err)	

				w.WriteHeader(500)

				json.NewEncoder(w).Encode(&types.Response[types.User]{
					IsOk: false,
					Message: err.Error(),
					Data: nil,
				})
				
				return
		}
		defer r.Body.Close()

		user, err := s.repository.SignUp(&body)
		if err != nil {
			log.Println(err)

			w.WriteHeader(500)
			
			json.NewEncoder(w).Encode(&types.Response[types.User]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}
		
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)

		json.NewEncoder(w).Encode(&types.Response[types.User]{
			IsOk: true,
			Message: "ok",
			Data: user,
		})
	}
}

func (s *Server) mySelfHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Key")
		sign := r.Header.Get("Sign")

		user, err := s.repository.GetMySelf(key)
		if err != nil {
			log.Println(err)

			w.WriteHeader(401)

			json.NewEncoder(w).Encode(&types.Response[types.User]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})

			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)

			w.WriteHeader(500)

			json.NewEncoder(w).Encode(&types.Response[types.User]{
				IsOk: false,
				Message: err.Error(),
				Data: nil,
			})
		}
		defer r.Body.Close()

		if !tools.CompareSign(sign, &tools.SignParams{
			Method: r.Method,
			IsTLS: r.TLS != nil,
			Host: r.Host,
			Path: r.RequestURI,
			Body: string(body),
			Secret: user.Secret,
		}) {
			w.WriteHeader(401)

			json.NewEncoder(w).Encode(&types.Response[types.User]{
				IsOk: false,
				Message: "You Are Not Authorized",
				Data: nil,
			})

			return
		}

		w.WriteHeader(200)

		json.NewEncoder(w).Encode(&types.Response[types.User]{
			IsOk: true,
			Message: "ok",
			Data: user,
		})
	}
}