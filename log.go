package filework

// A Log is a set of file copy results.
type Log []Result

/*
// Select returns results with the given outcome.
func (lg Log) Select(outcome Outcome) Log {
	var selection Log
	for i := range lg {
		if lg[i].Outcome == outcome {
			selection = append(selection, lg[i])
		}
	}
	return selection
}
*/

// Err returns the first error in the log if one exists, or nil.
func (lg Log) Err() error {
	for i := range lg {
		if lg[i].Err != nil {
			return lg[i].Err
		}
	}
	return nil
}
