package filework

import "os"

// Action describes an action taken.
type Action struct {
	SourcePath string
	DestPath   string
	Op         Op
}

// Result is the result of a filework operation.
type Result struct {
	Action
	os.FileMode
	Outcome Outcome
	Err     error
}
