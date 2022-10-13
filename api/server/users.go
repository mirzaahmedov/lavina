package server

import (
	"encoding/json"
	"net/http"

	"github.com/mirzaahmedov/lavina/api/types"
)

func (s *Server)signUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body types.User

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {				
				s.respondWithError(w, 400, "Can Not Read Request Body")
				return
		}
		defer r.Body.Close()

		user, err := s.repository.SignUp(&body)
		if err != nil {
			s.respondWithError(w, 500, err.Error())
			return
		}
		
		s.respondWithJson(w, 201, user)
	}
}

func (s *Server) mySelfHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth,_ := s.checkSign(r)
		if !isAuth {
			s.respondWithError(w, 401, "You Are Not Auhtorized")
			return
		}

		key := r.Header.Get("Key")

		user, err := s.repository.GetMySelf(key)
		if err != nil {
			s.respondWithError(w, 500, err.Error())
			return
		}

		s.respondWithJson(w, 200, user)
	}
}