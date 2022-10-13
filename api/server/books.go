package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mirzaahmedov/lavina/api/types"
)

type OpenLibraryResponse struct {
	Title string `json:"title"`
	Publishers []string `json:"publishers"`
	PublishDate string `json:"publish_date"`
	Isbn []string `json:"isbn_13"`
	Pages int `json:"number_of_pages"`
}

func (s *Server)createBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, userId := s.checkSign(r)
		if !isAuth {
			s.respondWithError(w, 401, "You Are Not Authorized")
			return
		}

		var body struct {
			Isbn string `json:"isbn"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			s.respondWithError(w, 400, "Can Not Read Request Body")
			return
		}
		defer r.Body.Close()

		response, err := http.Get("https://openlibrary.org/isbn/" + body.Isbn + ".json")
		if err != nil {
			s.respondWithError(w, 404, "Can Not Get Book From OpenLibrary")
			return
		}

		var book OpenLibraryResponse
		if err := json.NewDecoder(response.Body).Decode(&book); err != nil {
			s.respondWithError(w, 500, "Cannot Parse Results")
			return
		}
		response.Body.Close()

		newBook, err := s.repository.CreateBook(&types.Book{
			Book: types.BookInfo {
				Title: book.Title,
				Isbn: book.Isbn[0],
				Author: book.Publishers[0],
				Published: book.PublishDate,
				Pages: book.Pages,
			},
			UserId: userId,
		})
		if err != nil {
			s.respondWithError(w, 500, err.Error())
			return
		}

		s.respondWithJson(w, 201, newBook)
	}
}

func (s *Server) getAllBooksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, userId := s.checkSign(r)
		if !isAuth {
			s.respondWithError(w, 401, "You Are Not Authorized")
		}

		books, err := s.repository.GetAllBooks(userId)
		if err != nil {
			s.respondWithError(w, 500, err.Error())
		}

		s.respondWithJson(w, 200, books)
	}
}

func (s *Server) deleteBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, userId := s.checkSign(r)
		if !isAuth {
			s.respondWithError(w, 403, "Not Authorized")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.respondWithError(w, 400, "Invalid Book Id")
			return
		}

		if err := s.repository.DeleteBook(id, userId); err != nil {
			s.respondWithError(w, 500, err.Error())
			return
		}

		s.respondWithJson(w, 200, "Successfully deleted")
	}
}

func (s *Server) updateBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, userId := s.checkSign(r)
		if !isAuth {
			s.respondWithError(w, 403, "Not Authorized")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.respondWithError(w, 400, "Invalid Book Id")
			return
		}

		var updates struct {
			Status int `json:"status,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			s.respondWithError(w, 400, "Can Not Read Request Body")
		}

		modified, err := s.repository.UpdateBook(id, userId, updates.Status)
		if err != nil {
			s.respondWithError(w, 500, err.Error())
			return
		}

		s.respondWithJson(w, 200, modified)
	}
}