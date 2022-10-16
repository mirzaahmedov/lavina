package server

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"unicode"

	"github.com/mirzaahmedov/lavina/api/types"
)

func (s *Server)respondWithError(w http.ResponseWriter, status int, message string) {
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(types.Response{
		IsOk: false,
		Message: message,
		Data: nil,
	})
}
func (s *Server)respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(types.Response{
		IsOk: true,
		Message: "ok",
		Data: payload,	
	})	
}

func (s *Server)checkSign(r *http.Request) (bool, int)  {
		key := r.Header.Get("Key")
		sign := r.Header.Get("Sign")
		
		user, err := s.repository.GetMySelf(key)
		if err != nil {
			return false, -1
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return false, -1
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		body = []byte(strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
					return -1
			}
			return r
		}, string(body)))

		payload := r.Method + "https" + "://" + r.Host + r.RequestURI + string(body) + user.Secret

		sum := make([]byte, 16)
		if _, err := hex.Decode(sum, []byte(sign)); err != nil {
			return false, -1
		} else {
			result := md5.Sum([]byte(payload))
			if bytes.Equal(result[:], sum) {
				return true, user.Id
			} else {
				return false, -1
			}
		}
}