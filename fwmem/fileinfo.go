package fwmem

import (
	"os"
	"time"
)

// FileInfo holds information about a file.
type FileInfo struct {
	name     string
	size     int64
	dir      bool
	modified time.Time
}

// Name returns the name of the file.
func (fi FileInfo) Name() string {
	return fi.name
}

// Size returns the size of the file in bytes.
func (fi FileInfo) Size() int64 {
	return fi.size
}

// Mode returns the mode of the file.
func (fi FileInfo) Mode() os.FileMode {
	return 0
}

// ModTime returns the last modification time of the file.
func (fi FileInfo) ModTime() time.Time {
	return fi.modified
}

// IsDir returns true if the file is a directory.
func (fi FileInfo) IsDir() bool {
	return fi.dir
}

// Sys returns the underlying data source for the file. It returns nil.
//
// TODO: Return *File or FileInfo?
func (fi FileInfo) Sys() interface{} {
	return nil
}
