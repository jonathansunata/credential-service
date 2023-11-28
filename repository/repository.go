// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Host          string
	Port          string
	User          string
	Password      string
	Dsn           string
	DBName        string
	ConnectionUrl string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.ConnectionUrl)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: db,
	}
}
