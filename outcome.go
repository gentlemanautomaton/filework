package filework

import "fmt"

// Outcome holds the outcome of an attempted action on a file or set of files.
type Outcome int

// Outcome constants.
const (
	Unchanged Outcome = 0
	Created   Outcome = 1
	Modified  Outcome = 2
	Deleted   Outcome = 3
)

// String returns a string of the outcome type.
func (outcome Outcome) String() string {
	switch outcome {
	case Unchanged:
		return "Unchanged"
	case Created:
		return "Created"
	case Modified:
		return "Modified"
	case Deleted:
		return "Deleted"
	default:
		return fmt.Sprintf("Outcome(%d)", outcome)
	}
}
