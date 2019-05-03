package filework

import "os"

// FileSystem is the minimum interface required to read from a file system.
type FileSystem interface {
	Open(name string) (File, error)
	Stat(name string) (os.FileInfo, error)
}

// FileSystemWriter is the minimum interface required to write to a
// file system.
type FileSystemWriter interface {
	FileSystem
	Create(name string) (FileWriter, error)
	Remove(path string) error
	RemoveAll(path string) error
	MkdirAll(path string, perm os.FileMode) error
}
