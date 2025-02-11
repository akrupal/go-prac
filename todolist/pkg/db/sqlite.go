package db

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func CreateDb() (*sqlx.DB, error) {
	log.Debug().Msg("Creating Db")

	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("sqlite3", filepath.Join(filepath.Dir(ex), "todolist.db"))
	if err != nil {
		return nil, err
	}
	CreateSchema(db)
	return db, nil
}

func CreateSchema(db *sqlx.DB) {
	log.Debug().Msg("Creating Table")
	var schema = `
	DROP TABLE IF EXISTS todolist;
	CREATE TABLE todolist(
		id CHAR(40) NOT NULL,
		item VARCHAR(250) NOT NULL,
		CONSTRAINT rid_pkey PRIMARY KEY (id)
	);
	`
	db.MustExec(schema)
	log.Debug().Msg("DB init Completed")
}
