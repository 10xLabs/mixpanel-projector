package store

import "errors"

// Errors ...
var (
	ErrNoDocuments      = errors.New("no documents in result")
	ErrDocumentsMissing = errors.New("some documents could not be found")
)
