package filework

import (
	"io"
	"os"
)

// CompareFileContent compares content with a destination file system.
//
// If content implements io.Seeker, a seek to the start of content will be
// performed before and after a comparison is performed. NOTE: This behavior
// is subject to change in future revisions.
func CompareFileContent(content io.Reader, fs FileSystem, path string) (Comparison, error) {
	dest, err := fs.Open(path)
	switch {
	case os.IsNotExist(err):
		return Missing, nil
	case err != nil:
		return Different, err
	}
	if closer, ok := dest.(io.Closer); ok {
		defer closer.Close()
	}

	destInfo, err := dest.Stat()
	if err != nil {
		return Different, err
	}

	if destInfo.IsDir() {
		return Different, nil
	}

	if seeker, ok := content.(io.Seeker); ok {
		_, err = seeker.Seek(0, io.SeekStart)
		if err != nil {
			return Different, err
		}
		defer seeker.Seek(0, io.SeekStart)
	}

	same, err := CompareContent(content, dest)
	if err != nil {
		return Different, err
	}
	if !same {
		return Different, nil
	}

	return Same, nil
}
