package filework

import (
	"os"
)

// CopyDir copies a directory and all of its contents.
//
// TODO: Determine content copying by option?
// TODO: Determine perms by option?
//
// TODO: Receive a context?
func CopyDir(sourceFS FileSystem, sourcePath string, source Dir, destFS FileSystemWriter, destPath string, destPerms os.FileMode) (lg Log) {
	action := Action{SourcePath: sourcePath, DestPath: destPath, Op: CopyOperation}

	destInfo, err := destFS.Stat(destPath)
	switch {
	case err == nil:
		if !destInfo.IsDir() {
			return Log{Result{Action: action, Outcome: Unchanged, Err: ErrExistingFile}}
		}
		lg = Log{Result{Action: action, Outcome: Unchanged}}
		// TODO: Update perms?
	case os.IsNotExist(err):
		if err := destFS.MkdirAll(destPath, destPerms.Perm()); err != nil {
			return Log{Result{Action: action, Outcome: Unchanged, Err: err}}
		}
		lg = Log{Result{Action: action, Outcome: Created}}
	default:
		return Log{Result{Action: action, Outcome: Unchanged, Err: err}}
	}

	// TODO: process in chunks?

	dirents, err := source.Readdir(-1)
	if err != nil {
		return Log{Result{Action: action, Outcome: Unchanged, Err: err}}
	}
	for _, fi := range dirents {
		if name := fi.Name(); name != "" {
			sourceChildPath := fsJoin(sourceFS, sourcePath, name)
			destChildPath := fsJoin(destFS, destPath, name)
			lg = append(lg, CopyPath(sourceFS, sourceChildPath, destFS, destChildPath)...)
		}
	}

	return lg
}
