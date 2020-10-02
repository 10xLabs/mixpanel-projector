package repository

import "errors"

// Errors ...
var (
	ErrEmptyID          = errors.New("empty id")
	ErrNoDocuments      = errors.New("no documents found")
	ErrDocumentsMissing = errors.New("documents missing")
)
