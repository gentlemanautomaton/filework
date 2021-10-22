package filework

import "io"

// CopyFileContent copies data from content to fs and reports the outcome.
func CopyFileContent(content io.Reader, fs FileSystemWriter, path string) (Outcome, error) {
	comp, err := CompareFileContent(content, fs, path)
	if err != nil {
		return Unchanged, err
	}
	if comp == Same {
		return Unchanged, nil
	}

	dest, err := fs.Create(path)
	if err != nil {
		return Unchanged, err
	}
	if closer, ok := dest.(io.Closer); ok {
		defer closer.Close()
	}

	if _, err := io.Copy(dest, content); err != nil {
		return Modified, err
	}

	if comp == Missing {
		return Created, nil
	}
	return Modified, nil
}
