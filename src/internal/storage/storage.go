package storage

import "errors"

var (
	ErrUrlExists = errors.New("alias already exists")
	ErrNotFound  = errors.New("url not found")
)
