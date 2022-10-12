package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Open() error {
	log.Println("Connecting To Database")

	db, err := sql.Open("postgres", "postgresql://postgres:Od6v8S83I9f4KuNuyx46@containers-us-west-98.railway.app:6530/railway")
	if err != nil {
		return err
	}
	r.db = db

	log.Println("Successfully Connected Database")
	return nil
}

func (r *Repository) Close() {
	log.Println("Closing Database Connection")
	err := r.db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connection Closed")
}