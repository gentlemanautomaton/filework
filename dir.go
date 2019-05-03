package filework

import "os"

// Dir is the minimum interface required to read directories.
type Dir interface {
	Readdir(count int) ([]os.FileInfo, error)
}
