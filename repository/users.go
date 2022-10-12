package repository

import (
	"github.com/mirzaahmedov/lavina/api/types"
)

func (r *Repository)SignUp(data *types.User) (*types.User, error) {
	var user types.User 
	
	if err := r.db.QueryRow(
		"INSERT INTO users (name, key, secret) VALUES ($1, $2, $3) RETURNING id, name, key, secret",
		data.Name,
		data.Key,
		data.Secret,
	).Scan(&user.Id, &user.Name, &user.Key, &user.Secret); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetMySelf(key string) (*types.User, error) {
	var user types.User

	if err := r.db.QueryRow(
		"SELECT id, name, key, secret FROM users WHERE key=$1",
		key,
	).Scan(&user.Id, &user.Name, &user.Key, &user.Secret); err != nil {
		return nil, err
	}

	return &user, nil
}