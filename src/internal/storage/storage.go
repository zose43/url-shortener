package storage

import "errors"

var (
	ErrUrlExists = errors.New("url already exists")
	ErrNotFound  = errors.New("url not found")
)
