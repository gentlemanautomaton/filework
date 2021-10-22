package fwmem

// File is a representation of a file that is buffered in memory.
type File struct {
	Name    string
	Content []byte
}

/*
// Read reads data from f.
func (f *File) Read(p []byte) (n int, err error) {
	//return 0, f.Buffer.Rea
	return 0, nil
}

// Stat returns information about f.
func (f *File) Stat() (os.FileInfo, error) {
	return FileInfo{
		name: f.name,
		size: int64(len(f.content)),
	}, nil
}
*/
