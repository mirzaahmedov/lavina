package repository

import "github.com/mirzaahmedov/lavina/api/types"

func (r *Repository) CreateBook(data *types.Book) (*types.Book, error) {
	var book types.Book

	if err := r.db.QueryRow(
		"INSERT INTO books (name, key, secret) VALUES ($1, $2, $3) RETURNING id, name, key, secret",
		data.me,
		data.Key,
		data.Secret,
	).Scan(&book.Id, &book.Name, &book.Key, &book.Secret); err != nil {
		return nil, err
	}

	return &user, nil
}