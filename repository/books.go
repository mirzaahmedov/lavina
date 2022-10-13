package repository

import (
	"github.com/mirzaahmedov/lavina/api/types"
)

func (r *Repository) CreateBook(data *types.Book) (*types.Book, error) {
	var book types.Book

	if err := r.db.QueryRow(
		"INSERT INTO books (title, author, published, pages, isbn, status, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, author, published, pages, isbn, status",
		data.Book.Title,
		data.Book.Author,
		data.Book.Published,
		data.Book.Pages,
		data.Book.Isbn,
		data.Status,
		data.UserId,
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

func (r *Repository) GetAllBooks(id int) (*[]types.Book, error) {
	var books []types.Book

	rows, err := r.db.Query(
		"SELECT id, title, author, published, pages, status, isbn FROM books WHERE id=$1;",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := types.Book{}
		rows.Scan(
			&book.Book.Id, 
			&book.Book.Title, 
			&book.Book.Author, 
			&book.Book.Published, 
			&book.Book.Pages, 
			&book.Status, 
			&book.Book.Isbn,
		)
		books = append(books, book)
	}

	return &books, nil
}

func (r *Repository) DeleteBook(id int, userId int) error {
	_, err := r.db.Exec(
		"DELETE FROM books WHERE id=$1 AND user_id=$2",
		id,
		userId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateBook(id int, userId int, status int) (*types.Book, error) {
	book := &types.Book{}

	if err := r.db.QueryRow(
		"UPDATE books SET status=$1 WHERE id=$2 AND user_id=$3 RETURNING id, title, author, published, pages, isbn, status",
		status,
		id,
		userId,
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

	return book, nil
}