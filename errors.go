package filework

import "errors"

var (
	// ErrExistingDirectory is returned when a directory is found where a file
	// was expected.
	ErrExistingDirectory = errors.New("a directory exists where a file was expected")

	// ErrExistingFile is returned when a file is found where a directory
	// was expected.
	ErrExistingFile = errors.New("a file exists where a directory was expected")

	// ErrBadDirectory is returned when a file is found where a directory
	// was expected.
	ErrBadDirectory = errors.New("a file does not implement the directory interface despite being marked as a directory")
)
