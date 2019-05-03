package filework

import (
	"io"
	"os"
)

// File is the minimum interface required to read files.
type File interface {
	io.Reader
	io.Seeker
	Stat() (os.FileInfo, error)
}

// FileWriter is the minimum interface required to write files.
type FileWriter interface {
	File
	Write(b []byte) (n int, err error)
}

// A Namer is a file capable of reporting its name.
type Namer interface {
	Name() string
}

// A Locker is a file capable of locking.
type Locker interface {
	Lock() error
	Unlock() error
}

// A Syncer is a file capable of flushing its data to stable storage.
type Syncer interface {
	Sync() error
}
