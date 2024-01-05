package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"url-shortener/src/internal/storage"
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

func (s *Storage) SaveURL(ctx context.Context, urlToSave string, alias string) (int64, error) {
	const op = "storage.sqlite.saveURL"

	stmt, err := s.db.PrepareContext(ctx, "insert into urls (alias, url) values (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, alias, urlToSave)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return 0, fmt.Errorf("%s %w", op, storage.ErrUrlExists)
		}

		return 0, fmt.Errorf("%s %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id, %s %w", op, err)
	}

	return id, nil
}
