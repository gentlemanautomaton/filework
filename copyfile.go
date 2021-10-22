package filework

import (
	"io"
)

// CopyFile copies data from source to dest.
func CopyFile(sourceFS FileSystem, sourcePath string, source File, destFS FileSystemWriter, destPath string) Result {
	action := Action{SourcePath: sourcePath, DestPath: destPath, Op: CopyOperation}

	diff, err := CompareFileContent(source, destFS, destPath)
	if err != nil {
		return Result{Action: action, Outcome: Unchanged, Err: err}
	}
	if diff == Same {
		return Result{Action: action, Outcome: Unchanged}
	}

	dest, err := destFS.Create(destPath)
	if err != nil {
		return Result{Action: action, Outcome: Unchanged, Err: err}
	}
	if closer, ok := dest.(io.Closer); ok {
		defer closer.Close()
	}

	if _, err := io.Copy(dest, source); err != nil {
		return Result{Action: action, Outcome: Modified, Err: err}
	}

	if diff == Missing {
		return Result{Action: action, Outcome: Created}
	}
	return Result{Action: action, Outcome: Modified}
}
