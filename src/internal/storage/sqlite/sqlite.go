package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.new.sqlite"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	// todo make migration
	_, err = db.Exec(`
create table if not exists urls (
    id integer primary key,
    alias text not null unique,
    url text not null
);
create index if not exists idx_alias on urls(alias);
`)

	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return &Storage{db: db}, nil
}
