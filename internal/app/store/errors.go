package store

import "errors"

var (
	// ErrRecordNotFound - Result not found
	ErrRecordNotFound = errors.New("Record not found")
	// ErrRecordNotFound - general select error
	ErrCanNotGetResults = errors.New("Can not get results")
)
