package filework

import "fmt"

// Comparison describes the difference between two files.
type Comparison int

// Comparison constants.
const (
	Different Comparison = 0
	Missing   Comparison = 1
	Same      Comparison = 2
)

// String returns a string representation of the comparison.
func (c Comparison) String() string {
	switch c {
	case Missing:
		return "Missing"
	case Same:
		return "Same"
	case Different:
		return "Different"
	default:
		return fmt.Sprintf("Comparison(%d)", c)
	}
}
