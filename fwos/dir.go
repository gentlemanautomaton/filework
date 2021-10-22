package fwos

import (
	"os"
	pathpkg "path"
	"path/filepath"

	"github.com/gentlemanautomaton/filework"
)

// Dir is a local filesystem implementing both source and destination.
type Dir string

// Open returns a file ready for reading.
func (d Dir) Open(name string) (filework.File, error) {
	return os.Open(d.join(name))
}

// Create returns a file ready for writing.
func (d Dir) Create(name string) (filework.FileWriter, error) {
	return os.Create(d.join(name))
}

// Stat returns information about a file.
func (d Dir) Stat(name string) (os.FileInfo, error) {
	return os.Stat(d.join(name)) // FIXME: Use Lstat intead?
}

// Remove removes path. If path is a directory, it must be empty.
func (d Dir) Remove(path string) error {
	return os.Remove(d.join(path))
}

// RemoveAll removes path. If path is a directory, all files and directories
// contained in it are also removed.
func (d Dir) RemoveAll(path string) error {
	return os.RemoveAll(d.join(path))
}

// MkdirAll creates the requested directory and all parent directories.
func (d Dir) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(d.join(path), perm)
}

// Join returns the directory resulting from d being joined with p.
func (d Dir) Join(path string) Dir {
	return Dir(pathpkg.Join(string(d), pathpkg.Clean("/"+path)))
}

// join returns path p joined to directory d.
func (d Dir) join(path string) string {
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, filepath.FromSlash(pathpkg.Clean("/"+path)))
}
