package repository

import "github.com/mirzaahmedov/lavina/api/types"

func (r *Repository) CreateBook(data *types.Book) (*types.Book, error) {
	var book types.Book

	if err := r.db.QueryRow(
		"INSERT INTO books (title, author, published, pages, isbn, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, author, published, pages, isbn, status",
		data.Book.Title,
		data.Book.Author,
		data.Book.Published,
		data.Book.Pages,
		data.Book.Isbn,
		data.Status,
	).Scan(
		&book.Book.Id,
		&book.Book.Title, 
		&book.Book.Author, 
		&book.Book.Published, 
		&book.Book.Pages, 
		&book.Book.Isbn, 
		&book.Status,
	); err != nil {
		return nil, err
	}

	return &book, nil
}