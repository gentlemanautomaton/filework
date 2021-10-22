package filework

import (
	"io"
)

// CopyPath copies a file or directory described by path from source to dest
// and returns a log of its actions.
//
// TODO: Allow the user to specify perms for directories?
func CopyPath(sourceFS FileSystem, sourcePath string, destFS FileSystemWriter, destPath string) (lg Log) {
	action := Action{SourcePath: sourcePath, DestPath: destPath, Op: CopyOperation}

	// Examine source
	source, err := sourceFS.Open(sourcePath)
	if err != nil {
		return Log{Result{Action: action, Outcome: Unchanged, Err: err}}
	}
	if closer, ok := source.(io.Closer); ok {
		defer closer.Close()
	}

	sourceInfo, err := source.Stat()
	if err != nil {
		return Log{Result{Action: action, Outcome: Unchanged, Err: err}}
	}
	mode := sourceInfo.Mode()

	// Copy directories
	if mode.IsDir() {
		dir, ok := source.(Dir)
		if !ok {
			return Log{Result{Action: action, Outcome: Unchanged, Err: ErrBadDirectory}}
		}
		return CopyDir(sourceFS, sourcePath, dir, destFS, destPath, mode.Perm())
	}

	// Copy files
	return Log{CopyFile(sourceFS, sourcePath, source, destFS, destPath)}
}
