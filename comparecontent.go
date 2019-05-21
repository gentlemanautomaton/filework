package filework

import (
	"bytes"
	"io"
)

// TODO: Include multiple comparison functions optimized for different
// expected sizes? Use a chunk size option?

// TODO: Add compareSeekers function with concurrency limit, which calls
// compareReaders?

// CompareContent compares the content of two readers. It returns true if
// the content is the same.
func CompareContent(a, b io.Reader) (same bool, err error) {
	const chunkSize = 8 * 1024 // 8 kibibytes
	var scratch1, scratch2 [chunkSize]byte

	// TODO: Make reads concurrent

	// TODO: Make a fast path for small files, and an optimized path for
	//       large files.

	// TODO: If one or both readers provide Stat() functions, use that
	//       to compare file sizes.

	for {
		b1, b2 := scratch1[:], scratch2[:]
		n1, err1 := a.Read(b1)
		if err1 != nil && err1 != io.EOF {
			return false, err1
		}
		n2, err2 := b.Read(b2)
		if err2 != nil && err2 != io.EOF {
			return false, err2
		}
		if n1 != n2 {
			return false, nil
		}
		if n1 > 0 {
			if !bytes.Equal(b1[:n1], b2[:n2]) {
				return false, nil
			}
		}
		if err1 == io.EOF && err2 == io.EOF {
			return true, nil
		}
	}
}
