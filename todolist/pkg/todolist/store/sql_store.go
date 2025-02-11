package store

import "github.com/jmoiron/sqlx"

type sqlStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) Store {
	return &sqlStore{
		db: db,
	}
}
