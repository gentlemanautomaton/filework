package fwmem

import (
	"errors"
	"os"

	"github.com/gentlemanautomaton/filework"
)

// Dir is an in-memory directory.
type Dir struct {
	files map[string]*File
}

// Open opens a file within d.
func (d *Dir) Open(name string) (filework.File, error) {
	return nil, errors.New("fwmem: not implemented: Dir.Open")
}

func (d *Dir) open(name string) *File {
	f, found := d.files[name]
	if !found {
		return nil
	}
	return f
}

// Stat returns information about a file within d.
func (d *Dir) Stat(name string) (os.FileInfo, error) {
	//f, ok := d.files
	return nil, errors.New("fwmem: not implemented: Dir.Stat")
}

// Create returns a file writer for a file within d.
func (d *Dir) Create(name string) (filework.FileWriter, error) {
	return nil, errors.New("fwmem: not implemented: Dir.Create")
}

// Remove removes a file or directory within d.
func (d *Dir) Remove(path string) error {
	return errors.New("fwmem: not implemented: Dir.Remove")
}

// RemoveAll removes a set of subdirectories belonging to a path within d.
func (d *Dir) RemoveAll(path string) error {
	return errors.New("fwmem: not implemented: Dir.RemoveAll")
}

// MkdirAll creates a set of subdirectories belonging to a path within d.
func (d *Dir) MkdirAll(path string, perm os.FileMode) error {
	return errors.New("fwmem: not implemented: Dir.MkdirAll")
}
